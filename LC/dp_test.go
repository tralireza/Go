package lc

import (
	"log"
	"math"
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

		profit[0][0][1] = math.MinInt

		for i := 1; i < len(prices); i++ {
			log.Print("-> ", profit)
		}

		// last day [n-1]-- of using transaction allowed -[1]- with no stock left (ie sell) --[0]
		return profit[len(prices)-1][1][0]
	}

	log.Print("8 ?= ", maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}

// 121e Best Time to Buy & Sell: Kadane's algorithm
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
