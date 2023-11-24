package lc

import (
	"log"
	"testing"
)

func init() {
	log.Print("> BinSearch")
}

func TestBSLeftMost(t *testing.T) {
	bSearch := func(nums []int, k int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if nums[m] >= k {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	log.Print("2 ?= ", bSearch([]int{1, 2, 3, 4, 5}, 3))
	log.Print("0 ?= ", bSearch([]int{1, 2, 3}, 1))
	log.Print("0 ?= ", bSearch([]int{1, 2, 3, 4}, 0))
	log.Print("2 ?= ", bSearch([]int{1, 2, 5, 7}, 3))
	log.Print("3 ?= ", bSearch([]int{1, 2, 5, 7}, 7))
	log.Print("4 ?= ", bSearch([]int{1, 2, 3, 4}, 5))
}
