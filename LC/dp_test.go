package lc

import (
	"log"
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

		log.Print("-> ", profit)

		// last day [n-1]-- of using transaction allowed -[1]- with no stock left (ie sell) --[0]
		return profit[len(prices)-1][1][0]
	}

	log.Print("8 ?= ", maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}
