package lc

import (
	"log"
	"testing"
)

func init() {}

// 930m
func Test930(t *testing.T) {
	// PrefixSum & HashMap
	countSubarraysWithSum := func(nums []int, goal int) int {
		frq := map[int]int{}

		x, pfxSum := 0, 0
		for _, n := range nums {
			pfxSum += n
			if pfxSum == goal {
				x++
			}
			if f, ok := frq[pfxSum-goal]; ok {
				x += f
			}
			frq[pfxSum]++
		}

		return x
	}

	// SlidingWindow -> space: O(1)
	countSubarraysWithSum2 := func(nums []int, goal int) int {
		x := 0

		l, csum := 0, 0
		leadingZeros := 0
		for r, n := range nums {
			csum += n

			for ; l < r && nums[l] == 0 || csum > goal; l++ {
				if nums[l] == 0 {
					leadingZeros++
				} else {
					leadingZeros = 0
				}
				csum -= nums[l]
			}

			if csum == goal {
				x += 1 + leadingZeros
			}
		}

		return x
	}

	log.Print("4 ?= ", countSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	log.Print("8 ?= ", countSubarraysWithSum([]int{1, 0, 1, 0, 1}, 1))
	log.Print("15 ?= ", countSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))

	log.Print("4 ?= ", countSubarraysWithSum2([]int{1, 0, 1, 0, 1}, 2))
}
