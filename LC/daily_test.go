package lc

import (
	"log"
	"slices"
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

	// SlidingWindow -> space: O(1)
	countSubarraysWithSum3 := func(nums []int, goal int) int {
		// all subarrays with sum of at least: v
		atLeast := func(v int) int {
			x := 0
			l, csum := 0, 0
			for r := range nums {
				csum += nums[r]
				for l <= r && csum > v {
					csum -= nums[l]
					l++
				}
				x += r - l + 1
			}
			return x
		}
		return atLeast(goal) - atLeast(goal-1)
	}

	log.Print("4 ?= ", countSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	log.Print("8 ?= ", countSubarraysWithSum([]int{1, 0, 1, 0, 1}, 1))
	log.Print("15 ?= ", countSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))

	log.Print("4 ?= ", countSubarraysWithSum2([]int{1, 0, 1, 0, 1}, 2))

	log.Print("8 ?= ", countSubarraysWithSum3([]int{1, 0, 1, 0, 1}, 1))
	log.Print("2 ?= ", countSubarraysWithSum3([]int{1, 0, 1, 0, 1}, 0))
}

// 525m Contiguous Array
func Test525(t *testing.T) {
	findMaxLength := func(nums []int) int {
		idx := map[int]int{}
		idx[0] = -1

		x, v := 0, 0
		for i, n := range nums {
			if n == 0 {
				v--
			} else {
				v++
			}

			if start, ok := idx[v]; ok {
				x = max(x, i-start)
			} else {
				idx[v] = i
			}
		}

		return x
	}

	log.Print("2 ?= ", findMaxLength([]int{0, 1}))
	log.Print("2 ?= ", findMaxLength([]int{0, 1, 1, 0, 1, 0, 0}))
}

// 452m Min Arrows to Burst Ballons
func Test452(t *testing.T) {
	findMinArrowShots := func(points [][]int) int {
		slices.SortFunc(points, func(a, b []int) int { return a[1] - b[1] })
		log.Print(points)

		x, h := 1, points[0][1]
		for i := 1; i < len(points); i++ {
			lower, top := points[i][0], points[i][1]
			if lower > h {
				x++
				h = top
			}
		}
		return x
	}

	log.Print("2 ?= ", findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	log.Print("2 ?= ", findMinArrowShots([][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}))
}
