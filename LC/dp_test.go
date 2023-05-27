package lc

import (
	"bytes"
	"container/heap"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"reflect"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"
)

func init() {
	log.Print("> DP (mXdim)")
}

// DP: Rod-Cutting
func TestRodCutting(t *testing.T) {
	fCalls, cache := 0, map[int]int{}

	var revenue func(int, []int, bool) int
	revenue = func(length int, prices []int, useCache bool) int {
		fCalls++
		if length == 0 {
			return 0
		}
		if r, ok := cache[length]; ok {
			return r
		}

		r := 0
		for i := 1; i <= length; i++ {
			r = max(r, prices[i-1]+revenue(length-i, prices, useCache))
		}
		if useCache {
			cache[length] = r
		}
		return r
	}

	bottomUp := func(length int, prices []int) int {
		D := make([]int, length+1)     // DP: max revenue for length
		Cut := make([][]int, length+1) // Recording choices

		for l := 1; l <= length; l++ {
			for cutAt := 1; cutAt <= l; cutAt++ {
				if D[l] < prices[cutAt-1]+D[l-cutAt] {
					D[l] = prices[cutAt-1] + D[l-cutAt]
					Cut[l] = []int{cutAt}
				}
			}

			// record other equals choices
			for cutAt := Cut[l][0] + 1; cutAt <= l; cutAt++ {
				if D[l] == prices[cutAt-1]+D[l-cutAt] {
					Cut[l] = append(Cut[l], cutAt)
				}
			}
		}

		log.Print(D)
		log.Print(Cut)

		return D[length]
	}

	prices := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30, 31, 33, 35, 37, 40}
	for _, length := range []int{3, 5, 7, 9, 10, 15} {
		log.Printf(" (%2d) ?= %2d", length, bottomUp(length, prices))
		for _, useCache := range []bool{false, true} {
			ts := time.Now()
			revenue := revenue(length, prices, useCache)
			delta := time.Since(ts)
			log.Printf(" (%2d) ?= %2d | recursive: %5d   caching? %5t  %v", length, revenue, fCalls, useCache, delta)
			fCalls = 0
			clear(cache)
		}
		log.Print("===")
	}
}

// DP: Matrix-Chain Multiplication
func TestMatrixChainMultiplication(t *testing.T) {
	bottomUp := func(dimensions [][]int) int {
		// D: M(i)...M(j) multiplication minimum cost
		D, K := make([][]int, len(dimensions)), make([][]int, len(dimensions))
		for i := range dimensions {
			D[i] = make([]int, len(dimensions))
			K[i] = make([]int, len(dimensions))
		}

		for length := 2; length <= len(D); length++ {
			for start := 0; start <= len(D)-length; start++ {
				end := start + length - 1
				D[start][end] = math.MaxInt
				for k := start; k < end; k++ {
					kCost := D[start][k] + D[k+1][end] + dimensions[start][0]*dimensions[k][1]*dimensions[end][1]
					if kCost < D[start][end] {
						D[start][end] = kCost
						K[start][end] = k
					}
				}
			}
		}

		log.Print(D)
		log.Print(K)

		var draw func(s, e int)
		draw = func(s, e int) {
			if s == e {
				fmt.Printf("A%d", s+1)
			} else {
				fmt.Print("(")
				draw(s, K[s][e])
				draw(K[s][e]+1, e)
				fmt.Print(")")
			}
		}
		draw(0, len(dimensions)-1)
		fmt.Print("\n")

		return D[0][len(dimensions)-1]
	}

	log.Print(" ?= ", bottomUp([][]int{{30, 35}, {35, 15}, {15, 5}, {5, 10}, {10, 20}, {20, 25}}))
	log.Print(" ?= ", bottomUp([][]int{{5, 10}, {10, 3}, {3, 12}, {12, 5}, {5, 50}, {50, 6}}))
	log.Print(" ?= ", bottomUp([][]int{{10, 100}, {100, 5}, {5, 50}}))
}

// 714m Best Time to Buy & Sell with Fee
func Test714(t *testing.T) {
	maxProfit := func(prices []int, fee int) int {
		// [day][transactions][action]
		profit := make([][2][2]int, len(prices))

		profit[0][1][1] = -prices[0] - fee
		for i := 1; i < len(prices); i++ {
			log.Print(i, profit)
			profit[i][1][0] = max(profit[i-1][1][0], profit[i-1][1][1]+prices[i])     // Sell or Rset
			profit[i][1][1] = max(profit[i-1][1][1], profit[i-1][1][0]-prices[i]-fee) // Buy or Rest
		}
		log.Print("> ", profit)

		// last day [n-1]-- of using transaction allowed -[1]- with no stock left (ie sell) --[0]
		return profit[len(prices)-1][1][0]
	}

	maxProfitO1 := func(prices []int, fee int) int {
		var wStock, woStock int // with & without Stock
		for i := range prices {
			if i == 0 {
				wStock = -prices[0] - fee
			} else {
				wStock, woStock = max(wStock, woStock-prices[i]-fee), max(woStock, wStock+prices[i])
			}
		}
		return woStock
	}

	log.Print("8 ?= ", maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	log.Print("6 ?= ", maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))

	log.Print("8 ?= ", maxProfitO1([]int{1, 3, 2, 8, 4, 9}, 2))
	log.Print("6 ?= ", maxProfitO1([]int{1, 3, 7, 5, 10, 3}, 3))
}

// 121 Best Time to Buy & Sell: Kadane's algorithm
func Test121(t *testing.T) {
	maxProfix := func(prices []int) int {
		// Kadane's
		best, cur := 0, 0
		for i := 0; i < len(prices); i++ {
			diff := 0
			if i > 0 {
				diff = prices[i] - prices[i-1]
			}

			cur = max(diff, cur+diff)
			best = max(best, cur)
		}
		return best
	}

	log.Print("5 ?= ", maxProfix([]int{7, 1, 5, 3, 6, 4}))
	log.Print("0 ?= ", maxProfix([]int{7, 6, 4, 3, 1}))
}

func TestCSV(t *testing.T) {
	log.Print(io.Copy(os.Stdout, strings.NewReader("Stdin->Stdout io.Copy n,err: ")))

	type Movie struct {
		Title        string `json:"title"`
		Director     string `json:"director"`
		YearReleased int    `json:"year_released,omitempty"`
	}

	movies := []Movie{
		{"Star Wars: Episode VIII", "Rian Johnson", 2107},
		{"-", "-", 0},
	}

	bfr := bytes.Buffer{}
	func() {
		cw := csv.NewWriter(&bfr)
		cw.Write([]string{"Move Title", "Director", "Year Released"})
		for _, m := range movies {
			if err := cw.Write([]string{m.Title, m.Director, strconv.Itoa(m.YearReleased)}); err != nil {
				t.Fatal(err)
			}
		}
		cw.Flush()
		log.Print(bfr.String())
	}()

	cr := csv.NewReader(&bfr)
	cr.Comma = ','
	cr.Comment = '#'

	for {
		record, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		log.Print(record)
	}
}

// 309m Best Time to Buy&Sell with Cooldown
func Test309(t *testing.T) {
	maxProfit := func(prices []int) int {
		profit := make([][2]int, len(prices))
		for i, p := range prices {
			switch i {
			case 0:
				profit[i][1] = -p
			case 1:
				profit[i][0] = max(profit[i-1][0], profit[i-1][1]+p)
				profit[i][1] = max(profit[i-1][1], profit[i-1][0]-p) // no Cooldown (no Sell yet)
			default:
				profit[i][0] = max(profit[i-1][0], profit[i-1][1]+p)
				profit[i][1] = max(profit[i-1][1], profit[i-2][0]-p) // 1 day Cooldown after Sell
			}
		}
		log.Print(profit)
		return profit[len(prices)-1][0]
	}

	log.Print("3 ?= ", maxProfit([]int{1, 2, 3, 0, 2}))
	log.Print("3 ?= ", maxProfit([]int{1, 2, 4}))
}

// 123h Best Time to Buy & Sell with Only k=2 Transactions
func Test123(t *testing.T) {
	maxProfit := func(prices []int) int {
		profit := make([][2 + 1][2]int, len(prices))

		profit[0][1][1] = -prices[0]
		profit[0][2][1] = -prices[0]
		for i := 1; i < len(prices); i++ {
			for k := 1; k <= 2; k++ {
				profit[i][k][0] = max(profit[i-1][k][0], profit[i-1][k][1]+prices[i])
				profit[i][k][1] = max(profit[i-1][k][1], profit[i-1][k-1][0]-prices[i])
			}
		}
		log.Print(profit)

		return profit[len(prices)-1][2][0]
	}

	log.Print("6 ?= ", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	log.Print("4 ?= ", maxProfit([]int{1, 2, 3, 4, 5}))
	log.Print("0 ?= ", maxProfit([]int{7, 6, 4, 3, 1}))
	log.Print("13 ?= ", maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
}

// 188h Best Time to Buy & Sell with at most K Transactions
func Test188(t *testing.T) {
	maxProfit := func(k int, prices []int) int {
		profit := make([][][2]int, len(prices))
		for i := range profit {
			profit[i] = make([][2]int, k+1)
		}

		for k := k; k > 0; k-- {
			profit[0][k][1] = -prices[0]
		}

		for i := 1; i < len(prices); i++ {
			for k := k; k > 0; k-- {
				profit[i][k][0] = max(profit[i-1][k][0], profit[i-1][k][1]+prices[i])
				profit[i][k][1] = max(profit[i-1][k][1], profit[i-1][k-1][0]-prices[i])
			}
		}
		log.Print(profit)

		return profit[len(prices)-1][k][0]
	}

	log.Print("15 ?= ", maxProfit(3, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
	log.Print("7 ?= ", maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

// Sum of Encrypted Integers
func Test3079(t *testing.T) {
	sumOfEncryptedInts := func(nums []int) int {
		x := 0
		for _, n := range nums {
			d, i := 0, 0
			for ; n > 0; i++ {
				d = max(d, n%10)
				n /= 10
			}
			for ; i > 0; i-- {
				n = 10*n + d
			}

			x += n
		}
		return x
	}

	log.Print("6 ?= ", sumOfEncryptedInts([]int{1, 2, 3}))
	log.Print("66 ?= ", sumOfEncryptedInts([]int{10, 21, 31}))
}

// 3080
type Qe struct{ n, i int }
type Q []Qe

func (q Q) Len() int           { return len(q) }
func (q Q) Less(i, j int) bool { return q[i].n < q[j].n || q[i].n == q[j].n && q[i].i < q[j].i }
func (q Q) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *Q) Push(x any)        { *q = append(*q, x.(Qe)) }
func (q *Q) Pop() any {
	v := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return v
}

func Test3080(t *testing.T) {
	unmarkedSumArray := func(nums []int, queries [][]int) []int64 {
		q, lsum := &Q{}, int64(0)
		for i, n := range nums {
			heap.Push(q, Qe{n, i})
			lsum += int64(n)
		}
		log.Print(q)

		mkd := make([]bool, len(nums))
		xs := []int64{}

		for _, qry := range queries {
			i, k := qry[0], qry[1]
			if !mkd[i] {
				mkd[i] = true
				lsum -= int64(nums[i])
			}
			for k > 0 && q.Len() > 0 {
				e := heap.Pop(q).(Qe)
				if !mkd[e.i] {
					mkd[e.i] = true
					lsum -= int64(e.n)
					k--
				}
			}
			xs = append(xs, lsum)
		}

		return xs
	}

	log.Print("[8 3 0] ?= ", unmarkedSumArray([]int{1, 2, 2, 1, 2, 3, 1}, [][]int{{1, 2}, {3, 3}, {4, 2}}))
}

// 57m Insert Interval
func Test57(t *testing.T) {
	insert := func(intervals [][]int, newInterval []int) [][]int {
		l, r := 0, len(intervals)
		for l < r {
			m := l + (r-l)>>1
			if intervals[m][0] >= newInterval[0] {
				r = m
			} else {
				l = m + 1
			}
		}

		intervals = append(intervals[:l], append([][]int{newInterval}, intervals[l:]...)...)
		log.Print(l, newInterval, " -> ", intervals)

		rs := [][]int{}
		for _, v := range intervals {
			if len(rs) == 0 || rs[len(rs)-1][1] < v[0] {
				rs = append(rs, v)
			} else {
				rs[len(rs)-1][1] = max(rs[len(rs)-1][1], v[1])
			}
		}
		return rs
	}

	insert2 := func(intervals [][]int, newInterval []int) [][]int {
		rs, n := [][]int{}, newInterval
		for i, v := range intervals {
			if n[1] < v[0] {
				rs = append(rs, n)
				return append(rs, intervals[i:]...)
			} else if n[0] > v[1] {
				rs = append(rs, v)
			} else {
				n[0], n[1] = min(n[0], v[0]), max(n[1], v[1])
			}
		}
		return append(rs, n)
	}

	log.Print("?= ", insert([][]int{{1, 3}, {6, 9}}, []int{1, 6}))
	log.Print("?= ", insert([][]int{{1, 3}, {4, 5}, {6, 9}}, []int{4, 7}))

	log.Print("?= ", insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	log.Print("?= ", insert2([][]int{{1, 3}, {6, 9}}, []int{2, 5}))

	log.Print("?= ", insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	log.Print("?= ", insert2([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}

// 72m Edit Distance
func Test72(t *testing.T) {
	minDistance := func(word1 string, word2 string) int {
		// dist[m][n] distinace of word1[0..m] and word2[0..n]
		// dist[0][n] = n, dist[m][0] = m
		dist := make([][]int, len(word1)+1)
		for i := range dist {
			dist[i] = make([]int, len(word2)+1)
			dist[i][0] = i
		}
		for j := 0; j <= len(word2); j++ {
			dist[0][j] = j
		}

		for i := 1; i <= len(word1); i++ {
			for j := 1; j <= len(word2); j++ {
				d := 0
				if word1[i-1] != word2[j-1] {
					d++
				}
				dist[i][j] = min(dist[i][j-1]+1, dist[i-1][j]+1, dist[i-1][j-1]+d)
			}
		}

		return dist[len(word1)][len(word2)]
	}

	// O(n) space
	minDistance2 := func(word1 string, word2 string) int {
		dist := make([][]int, 2)
		for i := range dist {
			dist[i] = make([]int, len(word2)+1)
		}
		for j := range dist[0] {
			dist[0][j] = j
		}

		for i := 0; i < len(word1); i++ {
			copy(dist[1], dist[0])
			left := i + 1
			for j := 1; j <= len(word2); j++ {
				d := 0
				if word1[i] != word2[j-1] {
					d++
				}
				dist[0][j] = min(left+1, dist[1][j]+1, dist[1][j-1]+d)
				left = dist[0][j]
			}
		}
		log.Print(dist)

		return dist[0][len(word2)]
	}

	log.Print("3 ?= ", minDistance("horse", "ros"))
	log.Print("1 ?= ", minDistance("AGTCTTAGTCCAG", "AGTCTAGTCCAG"))
	log.Print("3 ?= ", minDistance2("horse", "ros"))
	log.Print("1 ?= ", minDistance2("AGTCTTAGTCCAG", "AGTCTAGTCCAG"))
}

// 583m Delete Operation for Two Strings
func Test583(t *testing.T) {
	minDistance := func(word1 string, word2 string) int {
		dist := make([][]int, len(word1)+1)
		for i := range dist {
			dist[i] = make([]int, len(word2)+1)
			dist[i][0] = i
		}
		for j := range dist[0] {
			dist[0][j] = j
		}

		for i := 1; i <= len(word1); i++ {
			for j := 1; j <= len(word2); j++ {
				if word1[i-1] == word2[j-1] {
					// only no substitude
					dist[i][j] = dist[i-1][j-1]
				} else {
					// only delete from word1 or word2
					dist[i][j] = 1 + min(dist[i-1][j], dist[i][j-1])
				}
			}
		}

		return dist[len(word1)][len(word2)]
	}

	log.Print("2 ?= ", minDistance("a", "b"))
	log.Print("2 ?= ", minDistance("sea", "eat"))
}

// 1318m Minimum Flips to Make a OR b Equal to c
func Test1318(t *testing.T) {
	minFlips := func(a, b, c int) int {
		x := 0

		for a > 0 || b > 0 || c > 0 {
			if c&1 == 1 {
				if a&1 == 0 && b&1 == 0 {
					x++
				}
			} else {
				if a&1 == 1 {
					x++
				}
				if b&1 == 1 {
					x++
				}
			}
			a, b, c = a>>1, b>>1, c>>1
		}

		return x
	}

	log.Print("3 ?= ", minFlips(2, 6, 5))
	log.Print("1 ?= ", minFlips(4, 2, 7))
	log.Print("15 ?= ", minFlips(1000_000_000, 2, 7))
}

// 435m Non-overlappinng Intervals
func Test435(t *testing.T) {
	eraseOverlapIntervals := func(intervals [][]int) int {
		slices.SortFunc(intervals, func(a, b []int) int { return a[1] - b[1] })

		x := 0
		h := 50_000 + 1 // -5*10^4 <= v(i) <= 5*10^4
		for i := len(intervals) - 1; i >= 0; i-- {
			cur := intervals[i]
			if h < cur[1] {
				x++
				h = max(cur[0], h)
			} else {
				h = cur[0]
			}
		}
		return x
	}

	f2 := func(intervals [][]int) int {
		slices.SortFunc(intervals, func(a, b []int) int { return a[1] - b[1] })

		x := 0
		h := -50_000 - 1
		for i := 0; i < len(intervals); i++ {
			cur := intervals[i]
			if h > cur[0] {
				x++
				h = min(h, cur[1])
			} else {
				h = cur[1]
			}
		}
		return x
	}

	intervalScheduling := func(intervals [][]int) int {
		slices.SortFunc(intervals, func(a, b []int) int { return a[1] - b[1] })
		mx := 0 // Maximum number of non-overlapping intervals
		h := -50_001
		for i := 0; i < len(intervals); i++ {
			cur := intervals[i]
			if h <= cur[0] {
				mx++
				h = cur[1]
			}
		}
		return len(intervals) - mx
	}

	for _, f := range []func([][]int) int{eraseOverlapIntervals, f2, intervalScheduling} {
		log.Print("+ ", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		log.Print("1 ?= ", f([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
		log.Print("2 ?= ", f([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}))
		log.Print("7 ?= ", f([][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}))
	}
}

// 496 Next Greater Element 1 Stack/Monotonic Stack
func Test496(t *testing.T) {
	nextGreaterElement := func(nums1, nums2 []int) []int {
		m := map[int]int{}

		S := []int{}
		for _, n := range nums2 {
			for len(S) > 0 && n > S[len(S)-1] {
				m[S[len(S)-1]] = n
				S = S[:len(S)-1]
			}
			S = append(S, n)
		}

		for i, n := range nums1 {
			nums1[i] = -1
			if v, ok := m[n]; ok {
				nums1[i] = v
			}
		}
		return nums1
	}

	log.Print("[7 9 9 7] ?= ", nextGreaterElement([]int{1, 3, 5, 2}, []int{6, 5, 4, 3, 9, 2, 1, 7}))
	log.Print("[7 7 7 -1] ?= ", nextGreaterElement([]int{1, 3, 5, 7}, []int{6, 5, 4, 3, 1, 7}))
}

// 1289 Minimum Falling Path Sum II
func Test1289(t *testing.T) {
	minFallingPathSum := func(grid [][]int) int {
		if len(grid) == 1 {
			return grid[0][0]
		}

		Rows, Cols := len(grid), len(grid[0])
		D := make([][]int, 2)
		for i := range D {
			D[i] = make([]int, Cols)
		}

		for r := range Rows {
			copy(D[0], D[1])

			for c := range Cols {
				D[1][c] = math.MaxInt

				for x := range Cols {
					if x != c {
						D[1][c] = min(D[1][c], D[0][x]+grid[r][c])
					}
				}
			}
		}

		return slices.Min(D[1])
	}

	topDown := func(grid [][]int) int {
		Rows, Cols := len(grid), len(grid[0])
		rCalls, mHits := 0, 0

		Mem := map[[2]int]int{}
		var dfs func(r, c int) int
		dfs = func(r, c int) int {
			rCalls++
			if r == Rows-1 {
				return grid[r][c]
			}

			if v, ok := Mem[[2]int{r, c}]; ok {
				mHits++
				return v
			}

			v := math.MaxInt
			for x := range Cols {
				if x != c {
					v = min(dfs(r+1, x), v)
				}
			}
			Mem[[2]int{r, c}] = v + grid[r][c]
			return v + grid[r][c]
		}

		n := math.MaxInt
		for c := range Cols {
			n = min(dfs(0, c), n)
		}
		log.Printf("rCalls: %d, Hits: %d", rCalls, mHits)
		return n
	}

	optimized := func(grid [][]int) int {
		prv, jPrv, prv2 := math.MaxInt, -1, math.MaxInt
		for j, v := range grid[0] {
			if v < prv {
				prv, jPrv, prv2 = v, j, prv
			} else if v < prv2 {
				prv2 = v
			}
		}

		for i := 1; i < len(grid); i++ {
			cur, jCur, cur2 := math.MaxInt, -1, math.MaxInt

			for j := 0; j < len(grid[i]); j++ {
				v := grid[i][j]

				if j != jPrv {
					v += prv
				} else {
					v += prv2
				}

				if v < cur {
					cur, jCur, cur2 = v, j, cur
				} else if v < cur2 {
					cur2 = v
				}
			}

			prv, jPrv, prv2 = cur, jCur, cur2
		}

		return prv
	}

	for _, f := range []func([][]int) int{minFallingPathSum, topDown, optimized} {
		log.Print("13 ?= ", f([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
		log.Print("7 ?= ", f([][]int{{7}}))
		log.Print("10 ?= ", f([][]int{{8, 4, 3, 7, 8}, {4, 5, 6, 5, 2}, {7, 8, 9, 3, 9}, {1, 5, 7, 1, 9}, {5, 7, 1, 3, 7}}))
		log.Print("7 ?= ", f([][]int{{2, 2, 1, 2, 2}, {2, 2, 1, 2, 2}, {2, 2, 1, 2, 2}, {2, 2, 1, 2, 2}, {2, 2, 1, 2, 2}}))
	}
}

// 5m Longest Palindromic Substring
func Test5(t *testing.T) {
	// O(n2) time & O(n) space
	centerXpand := func(s string) string {
		expand := func(l, r int) int {
			for l >= 0 && r < len(s) && s[l] == s[r] {
				l--
				r++
			}
			return r - l - 1
		}

		l, r := 0, 0
		for i := 0; i < len(s); i++ {
			lenO := expand(i, i)
			if r-l+1 < lenO {
				l, r = i-lenO/2, i+lenO/2
			}
			lenV := expand(i, i+1)
			if r-l+1 < lenV {
				l, r = i-lenV/2+1, i+lenV/2
			}
		}

		return s[l : r+1]
	}

	// O(n) time!
	Manacher := func(s string) string {
		S := "|" + strings.Join(strings.Split(s, ""), "|") + "|"
		Radius := make([]int, len(S))

		center, radius := 0, 0
		for i := range len(S) {
			mirror := 2*center - i

			if i < radius {
				Radius[i] = min(radius-i, Radius[mirror])
			}

			for i+1+Radius[i] < len(S) &&
				i-1-Radius[i] >= 0 &&
				S[i+1+Radius[i]] == S[i-1-Radius[i]] {
				Radius[i]++
			}

			if i+Radius[i] > radius {
				center = i
				radius = i + Radius[i]
			}
		}

		radius = slices.Max(Radius)
		center = slices.Index(Radius, radius)

		start := (center - radius) / 2
		return s[start : start+radius]
	}

	for _, f := range []func(string) string{longestPalindrome, centerXpand, Manacher} {
		if slices.Index([]string{"bab", "aba"}, f("babad")) < 0 {
			t.Fail()
		}
		if f("cbbd") != "bb" {
			t.Fail()
		}
		log.Print(" ?= ", f(strings.ReplaceAll("was it a car or a cat i saw?", " ", "")))
	}
}

// 64m Minimum Path Sum
func Test64(t *testing.T) {
	if minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}) != 7 {
		t.Fail()
	}
	if minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}) != 12 {
		t.Fail()
	}
}

// 118 Pascal's Triangle
func Test118(t *testing.T) {
	generate := func(numRows int) [][]int {
		Tri := [][]int{}

		for r := 0; r < numRows; r++ {
			row := make([]int, r+1)
			row[0], row[r] = 1, 1

			for c := 1; c < r; c++ {
				row[c] = Tri[r-1][c-1] + Tri[r-1][c]
			}

			Tri = append(Tri, row)
		}

		return Tri
	}

	log.Print(" ?= ", generate(7))
	log.Print(" ?= ", generate(5))
	log.Print(" ?= ", generate(1))
}

// 152m Maximum Product Subarray
func Test152(t *testing.T) {
	// Kadane's: Maximum Sum:
	// best, curr := -INF, 0
	// for n <- range
	//   curr = max(curr, n+curr)
	//   best = max(best, curr)
	Kadane := func(nums []int) int {
		n := nums[0]
		best, curr := n, n
		currSign := n // *
		for _, n := range nums[1:] {
			if n < 0 {
				curr, currSign = currSign, curr // *
			}
			curr = max(n, curr*n)
			best = max(best, curr)

			currSign = min(n, currSign*n) // *
		}

		return best
	}

	for _, f := range []func([]int) int{maxProduct, Kadane} {
		log.Print("6 ?= ", f([]int{2, 3, -2, 4}))
		log.Print("24 ?= ", f([]int{-2, 3, -4}))
		log.Print("0 ?= ", f([]int{-2, 0, -1}))
		log.Print("1 ?= ", f([]int{-2, 1}))
		log.Print("===")
	}
}

// 300m Longest Increasing Subsequence
func Test300(t *testing.T) {
	// LCS: Longest Common Subsequence
	LCS := func(rStr, cStr string) string {
		D := make([][]string, len(rStr)+1)
		for r := range D {
			D[r] = make([]string, len(cStr)+1)
		}

		for r := 1; r <= len(rStr); r++ {
			for c := 1; c <= len(cStr); c++ {
				if rStr[r-1] == cStr[c-1] {
					D[r][c] = D[r-1][c-1] + string(rStr[r-1])
				} else {
					D[r][c] = D[r][c-1]
					if len(D[r-1][c]) > len(D[r][c]) {
						D[r][c] = D[r-1][c]
					}
				}
			}
		}

		log.Printf("%q", D)
		return D[len(rStr)][len(cStr)]
	}

	dp := func(nums []int) int {
		sorted := make([]int, len(nums))
		copy(sorted, nums)

		slices.Sort(sorted)
		l := 0
		for r := range sorted {
			if sorted[l] < sorted[r] {
				l++
				sorted[l] = sorted[r]
			}
		}
		sorted = sorted[:l+1]

		// LCS sorted & nums
		D := make([][]int, len(sorted)+1)
		for r := range D {
			D[r] = make([]int, len(nums)+1)
		}

		for s := 1; s <= len(sorted); s++ {
			for n := 1; n <= len(nums); n++ {
				if sorted[s-1] == nums[n-1] {
					D[s][n] = D[s-1][n-1] + 1
				} else {
					D[s][n] = max(D[s-1][n], D[s][n-1])
				}
			}
		}

		return D[len(sorted)][len(nums)]
	}

	for _, f := range []func([]int) int{dp, lengthOfLIS} {
		log.Print("4 ?= ", f([]int{10, 9, 2, 5, 3, 7, 101, 18}))
		log.Print("4 ?= ", f([]int{0, 1, 0, 3, 2, 3}))
		log.Print("3 ?= ", f([]int{10, 9, 2, 5, 3, 4}))
		log.Print("6 ?= ", f([]int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}))
		log.Print("===")
	}

	log.Printf("LCS (1 outof 3) [%q,%q] -> %q", "GAC", "AGCAT", LCS("GAC", "AGCAT"))
}
