package lc

import (
	"bytes"
	"fmt"
	"log"
	"slices"
	"testing"
)

func init() {
	log.Print("> Basic DS")
}

// 20 Valid Parentheses
func Test20(t *testing.T) {
	isValid := func(s string) bool {
		S := []byte{}
		pairs := map[byte]byte{'(': ')', '[': ']', '{': '}'}

		for i := 0; i < len(s); i++ {
			switch s[i] {
			case ')', '}', ']':
				if len(S) == 0 || pairs[S[len(S)-1]] != s[i] {
					return false
				}
				S = S[:len(S)-1]
			default:
				S = append(S, s[i])
			}
		}

		return len(S) == 0
	}

	log.Print("true ?= ", isValid("()[]{}"))
	log.Print("true ?= ", isValid("({}[])"))
	log.Print("false ?= ", isValid("(}"))
}

// 1614 Maximum Nesting Depth of Parentheses
func Test1614(t *testing.T) {
	maxDepth := func(s string) int {
		S := []byte{}
		x := 0
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '(':
				S = append(S, '(')
				x = max(x, len(S))
			case ')':
				S = S[:len(S)-1]
			}
		}
		return x
	}

	log.Print("3 ?= ", maxDepth("(1+(2*3)+((8)/4))+1"))
	log.Print("3 ?= ", maxDepth("(1)+((2))+(((3)))"))
}

// 1544 Make String Great
func Test1544(t *testing.T) {
	makeGood := func(s string) string {
		i := 0
		for i < len(s)-1 {
			log.Printf("%q %3[1]d %q %3[2]d | %3d", s[i], s[i+1], s[i]-s[i+1])

			if s[i]-s[i+1] == 'a'-'A' || s[i]-s[i+1] == 256-('a'-'A') {
				s = s[:i] + s[i+2:]
				i = 0
			} else {
				i++
			}
		}
		return s
	}

	fmt.Print("byte 'Overflow' -> ")
	b := byte(253)
	for range 7 {
		fmt.Printf("%d,", b)
		b++
	}
	fmt.Print("\n")

	log.Print(" ?= ", makeGood("abBACc"))
	log.Print(" ?= ", makeGood("Pp"))
	log.Print("a ?= ", makeGood("aPp"))
}

// 1249m Minimum Remove to Make Valid Parentheses
func Test1249(t *testing.T) {
	minRemoveToMakeValid := func(s string) string {
		S := []int{}
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '(':
				S = append(S, i)
			case ')':
				if len(S) > 0 && s[S[len(S)-1]] == '(' {
					S = S[:len(S)-1]
				} else {
					S = append(S, i)
				}
			}
		}
		log.Printf("%v", S)

		bs := []byte{}
		for i := 0; i < len(s); i++ {
			if len(S) > 0 {
				if i < S[0] {
					bs = append(bs, s[i])
				} else {
					S = S[1:]
				}
			} else {
				bs = append(bs, s[i])
			}
		}
		return string(bs)
	}

	log.Print("lee(t(c)o)de ?= ", minRemoveToMakeValid("lee(t(c)o)de)"))
	log.Print(" ?= ", minRemoveToMakeValid("))(("))
}

// 678m Valid Parenthesis String
func Test678(t *testing.T) {
	checkValidString := func(s string) bool {
		type Y = struct{}

		m := map[int]Y{}
		m[0] = Y{}

		for i := 0; i < len(s); i++ {
			mNew := map[int]Y{}
			switch s[i] {
			case '(':
				for k := range m {
					mNew[k+1] = Y{}
				}
			case ')':
				for k := range m {
					if k > 0 {
						mNew[k-1] = Y{}
					}
				}
			case '*':
				for k := range m {
					mNew[k] = Y{}
					mNew[k+1] = Y{}
					if k > 0 {
						mNew[k-1] = Y{}
					}
				}
			}
			log.Printf("%q %v -> %v", s[i], m, mNew)
			m = mNew
		}
		log.Print(m)

		if _, ok := m[0]; ok {
			return true
		}
		return false
	}

	dynamic := func(s string) bool {
		// dp[endIndex][brackets] -> valid/notValid
		dp := make([][]byte, len(s)+1)
		for i := 0; i < len(dp); i++ {
			dp[i] = make([]byte, len(s)+1)
		}
		dp[0][0] = 1

		// substring s[0:i] ie, last index: i-1 s[0...(i-1)]
		for i := 1; i <= len(s); i++ {
			for j := 0; j <= i; j++ {
				switch s[i-1] {
				case '*':
					if j < len(s) {
						dp[i][j] = max(dp[i-1][j], dp[i-1][j+1])
					}
					if j > 0 {
						dp[i][j] = max(dp[i][j], dp[i-1][j-1])
					}
					dp[i][j] = max(dp[i][j], dp[i-1][j])
				case '(':
					if j > 0 {
						dp[i][j] = dp[i-1][j-1]
					}
				case ')':
					if j < len(s) {
						dp[i][j] = dp[i-1][j+1]
					}
				}
			}
		}

		return dp[len(s)][0] == 1
	}

	doubleStack := func(s string) bool {
		S, W := []int{}, []int{}

		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '(':
				S = append(S, i)
			case ')':
				if len(S) > 0 {
					S = S[:len(S)-1]
				} else if len(W) > 0 {
					W = W[:len(W)-1]
				} else {
					return false
				}
			case '*':
				W = append(W, i)
			}
		}
		log.Print(S, W)

		for len(S) > 0 {
			if len(W) > 0 && S[len(S)-1] < W[len(W)-1] {
				S, W = S[:len(S)-1], W[:len(W)-1]
			} else {
				return false
			}
		}
		log.Print(S, W)
		return len(S) == 0
	}

	for _, s := range []string{
		"((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()",
		"(*))",
		"(*()",
		"(*)",
		"((**(*)))",
		"(*)**)()",
	} {
		for _, f := range []func(string) bool{checkValidString, dynamic, doubleStack} {
			log.Print("true ?= ", f(s))
		}
	}
}

// 2073 Time Needed to Buy Tickets
func Test2073(t *testing.T) {
	timeRequiredToBuy := func(tickets []int, k int) int {
		t := 0
		for i, n := range tickets {
			if i <= k {
				t += min(n, tickets[k])
			} else {
				t += min(n, tickets[k]-1)
			}
		}
		return t
	}

	simulateQueue := func(tickets []int, k int) int {
		t := 0

		Q := []int{}
		for i := 0; i < len(tickets); i++ {
			Q = append(Q, i)
		}

		var front int
		for tickets[k] > 0 {
			front, Q = Q[0], Q[1:]
			if tickets[front] > 0 {
				t++
				tickets[front]--
			}
			Q = append(Q, front)
		}

		return t
	}

	for _, f := range []func([]int, int) int{simulateQueue, timeRequiredToBuy} {
		log.Print("6 ?= ", f([]int{2, 3, 2}, 2))
		log.Print("8 ?= ", f([]int{5, 1, 1, 1}, 0))
		log.Print(" ?= ", f([]int{5, 1, 2, 1}, 2))
	}
}

// 950m Reveal Cards in Increasing Order
func Test950(t *testing.T) {
	deckRevealedIncreasing := func(deck []int) []int {
		slices.Sort(deck)

		Q := []int{deck[len(deck)-1]}

		r := len(deck) - 1
		for r > 0 {
			back := Q[len(Q)-1]
			r--
			Q = append([]int{deck[r], back}, Q[:len(Q)-1]...)
		}

		return Q
	}

	log.Print("[2 13 3 11 5 17 7] ?= ", deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
}

// 402m Remove K Digits
func Test402(t *testing.T) {
	removeKdigits := func(num string, k int) string {
		S := []byte{}
		for i := 0; i < len(num); i++ {
			for len(S) > 0 && k > 0 && S[len(S)-1] > num[i] {
				S = S[:len(S)-1]
				k--
			}
			S = append(S, num[i])
		}
		log.Printf("%c", S)

		if k > 0 {
			S = S[:len(S)-k]
		}
		log.Printf("%c", S)

		v := string(bytes.TrimLeft(S, "0"))
		if v == "" {
			return "0"
		}
		return v
	}

	log.Print("1219 ?= ", removeKdigits("1432219", 3))
	log.Print("200 ?= ", removeKdigits("10200", 1))
	log.Print("0 ?= ", removeKdigits("10200", 2))
	log.Print("122 ?= ", removeKdigits("12234", 2))
}

// 42h Trapping Rain Water
func Test42(t *testing.T) {
	trap := func(height []int) int {
		l, r := make([]int, len(height)), make([]int, len(height))
		l[0], r[len(height)-1] = height[0], height[len(height)-1]

		for i := 1; i < len(height); i++ {
			l[i] = max(l[i-1], height[i])
		}
		for i := len(height) - 2; i >= 0; i-- {
			r[i] = max(r[i+1], height[i])
		}

		log.Print(l, r)

		w := 0
		for i := 0; i < len(height); i++ {
			w += min(l[i], r[i]) - height[i]
		}
		return w
	}

	pointers := func(height []int) int {
		l, r := 0, len(height)-1
		lx, rx := height[l], height[r]

		w := 0
		for l < r {
			if lx < rx {
				w += lx - height[l]
				l++
				lx = max(lx, height[l])
			} else {
				w += rx - height[r]
				r--
				rx = max(rx, height[r])
			}
		}
		return w
	}

	for _, f := range []func([]int) int{trap, pointers} {
		log.Print("6 ?= ", f([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
		log.Print("9 ?= ", f([]int{4, 2, 0, 3, 2, 5}))
	}
}

// 85 Maximal Rectangle
func Test85(t *testing.T) {
	maximalRectangle := func(matrix [][]byte) int {
		xarea := 0

		csum := make([][]int, len(matrix)+1)
		for i := range len(csum) {
			csum[i] = make([]int, len(matrix[0])+1)
		}
		for j := 0; j < len(matrix[0]); j++ {
			for i := 0; i < len(matrix); i++ {
				csum[i+1][j+1] = csum[i][j+1] + int(matrix[i][j]-'0')
			}
		}

		log.Printf("%q", matrix)
		log.Print(csum)

		for t := 0; t < len(matrix); t++ {
			for b := len(matrix); b > t; b-- {

				area := 0
				for c := 0; c < len(matrix[0]); c++ {
					if csum[b][c+1]-csum[t][c+1] == b-t {
						area += b - t
					} else {
						area = 0
					}
					xarea = max(xarea, area)
				}

			}
		}

		return xarea
	}

	log.Print("6 ?= ", maximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	log.Print("0 ?= ", maximalRectangle([][]byte{{'0'}}))
	log.Print("1 ?= ", maximalRectangle([][]byte{{'1'}}))
}

// 2371h Minimize Maximum Value in a Grid
func Test237(t *testing.T) {
	minScore := func(grid [][]int, fSelect func(grid, places [][]int)) [][]int {
		places := [][]int{} // grid[x][y], x, y
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				places = append(places, []int{grid[i][j], i, j})
			}
		}

		slices.SortFunc(places, func(a, b []int) int { return a[0] - b[0] })
		log.Print(places)

		log.Print(grid)
		fSelect(grid, places)

		return grid
	}

	// O(n3) n~m -> may TLE with large input
	simpleSelect := func(grid, places [][]int) {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				grid[i][j] = 0
			}
		}

		for _, v := range places {
			i, j := v[1], v[2]
			xcol := 0
			for k := 0; k < len(grid); k++ {
				xcol = max(xcol, grid[k][j])
			}
			x := max(slices.Max(grid[i]), xcol) + 1
			grid[i][j] = x
		}
	}

	// O(n2) n~m
	fastSelect := func(grid, places [][]int) {
		xrow, xcol := make([]int, len(grid)), make([]int, len(grid[0]))
		for _, v := range places {
			i, j := v[1], v[2]
			x := max(xrow[i], xcol[j]) + 1
			grid[i][j] = x
			xrow[i], xcol[j] = x, x
		}
	}

	for _, f := range []func([][]int, [][]int){simpleSelect, fastSelect} {
		log.Print(" ?= ", minScore([][]int{{3, 1}, {2, 5}}, f))
		log.Print(" ?= ", minScore([][]int{{2, 4, 5}, {7, 3, 9}}, f))
		log.Print(" ?= ", minScore([][]int{{2, 4, 5, 13}, {18, 3, 9, 1}, {12, 8, 17, 6}}, f))
	}
}

// 568h Maximum Vacation Days
func Test(t *testing.T) {
	maxVacationDays := func(flights, days [][]int) int {
		// W: weeks, M: cache
		W, M := len(days[0]), map[[2]int]int{}

		var calc func(c, w int) int
		calc = func(c, w int) int {
			if w >= W {
				return 0
			}

			if v, ok := M[[2]int{c, w}]; ok {
				return v
			}

			best := days[c][w] + calc(c, w+1)
			for i, f := range flights[c] {
				if f == 1 {
					best = max(days[i][w]+calc(i, w+1), best)
				}
			}

			M[[2]int{c, w}] = best
			return best
		}

		x := calc(0, 0)
		log.Print(M)

		return x
	}

	// DP: Bottom-Up
	bottomUp := func(flights, days [][]int) int {
		// dp Week X City -> max: vdays
		dp := make([][]int, len(days))
		for c := range days {
			dp[c] = make([]int, len(days[0]))
		}

		dp[0][0] = days[0][0]
		for w := 0; w < len(days[0]); w++ {
			for c := 0; c < len(days); c++ {
				if w == 0 && c == 0 || w > 0 && dp[w-1][c] > 0 {
					dp[w][c] = days[c][w]
					for d, f := range flights[c] {
						if f == 1 {
							dp[w][d] = days[d][w]
						}
					}
				}
			}
		}

		log.Print(dp)

		x := 0
		for w := 0; w < len(days[0]); w++ {
			x += slices.Max(dp[w])
		}
		return x
	}

	for _, f := range []func(flights, days [][]int) int{maxVacationDays, bottomUp} {
		log.Print("12 ?= ", f([][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}, [][]int{{1, 3, 1}, {6, 0, 3}, {3, 3, 3}}))
		log.Print("5 ?= ", f([][]int{{0, 0, 0}, {1, 0, 1}, {1, 1, 0}}, [][]int{{1, 3, 1}, {6, 0, 3}, {3, 3, 3}}))
	}
}
