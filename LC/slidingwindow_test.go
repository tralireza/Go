package lc

import (
	"log"
	"testing"
)

func init() {
	log.Print("> SlidingWindow")
}

// 1248m Count Number of Subarrays
func Test1248(t *testing.T) {
	numberOfSubarrays := func(nums []int, k int) int {
		for i, n := range nums {
			nums[i] = n & 1
		}
		log.Print(nums)

		leadingZeros := 0
		x := 0
		l, csum := 0, 0
		for r := range nums {
			csum += nums[r]
			for l <= r && nums[l] == 0 || csum > k {
				if nums[l] == 0 {
					leadingZeros++
				} else {
					leadingZeros = 0
				}
				csum -= nums[l]
			}
			if csum == k {
				x += 1 + leadingZeros
			}
		}
		return x
	}

	// 3 odd numbers in subarray: [1 7 3] [1 7 3 4] [7 3 4 5]
	log.Print("3 ?= ", numberOfSubarrays([]int{1, 7, 3, 4, 5}, 3))
}
