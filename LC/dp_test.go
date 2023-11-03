package lc

import (
	"bytes"
	"container/heap"
	"encoding/csv"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func init() {
	log.Print("> DP (mXdim)")
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
		slices.SortFunc(intervals,
			func(a, b []int) int {
				if a[0] == b[0] {
					return a[1] - b[1]
				}
				return a[0] - b[0]
			})
		log.Print(intervals)

		x := 0
		return x
	}

	log.Print("1 ?= ", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	log.Print("7 ?= ", eraseOverlapIntervals([][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}))
}
