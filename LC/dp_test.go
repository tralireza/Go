package lc

import (
	"bytes"
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
		Title        string
		Director     string
		YearReleased int
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
