package lc

import (
	"bytes"
	"container/heap"
	"encoding/csv"
	"io"
	"log"
	"os"
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

func TestString(t *testing.T) {
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
		profit := make([][2][2]int, len(prices))

		return profit[len(prices)-1][1][0]
	}

	log.Print("6 ?= ", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	log.Print("4 ?= ", maxProfit([]int{1, 2, 3, 4, 5}))
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

func (q Q) Less(i int, j int) bool {
	if q[i].n == q[j].n {
		return q[i].i < q[j].i
	}
	return q[i].n < q[j].n
}
func (q Q) Len() int          { return len(q) }
func (q Q) Swap(i int, j int) { q[i], q[j] = q[j], q[i] }
func (q *Q) Push(x any)       { *q = append(*q, x.(Qe)) }
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
