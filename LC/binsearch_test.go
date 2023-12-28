package lc

import (
	"log"
	"testing"
)

func init() {
	log.Print("> BinSearch")
}

func TestBSSearch(t *testing.T) {
	// Leftmost Index
	bSearchLeft := func(nums []int, k int) int {
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

	// Rightmost Index
	bSearchRight := func(nums []int, k int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if nums[m] > k {
				r = m
			} else {
				l = m + 1
			}
		}
		return r
	}

	log.Print("+ Rightmost Index")
	log.Print("4 ?= ", bSearchRight([]int{1, 2, 3, 4}, 5))
	log.Print("3 ?= ", bSearchRight([]int{1, 2, 3, 4, 5}, 3))
	log.Print("2 ?= ", bSearchRight([]int{1, 1, 3}, 1))
	log.Print("0 ?= ", bSearchRight([]int{1, 2, 3, 4}, 0))
	log.Print("+ Leftmost Index")
	log.Print("2 ?= ", bSearchLeft([]int{1, 2, 3, 4, 5}, 3))
	log.Print("0 ?= ", bSearchLeft([]int{1, 1, 3}, 1))
	log.Print("0 ?= ", bSearchLeft([]int{1, 2, 3, 4}, 0))
	log.Print("3 ?= ", bSearchLeft([]int{1, 2, 2, 7}, 3))
	log.Print("3 ?= ", bSearchLeft([]int{1, 2, 5, 7}, 7))
	log.Print("4 ?= ", bSearchLeft([]int{1, 2, 3, 4}, 5))
}

// 33m Search in Rotated Sorted Array
func Test33(t *testing.T) {
	search := func(nums []int, target int) int {
		l, r := 0, len(nums)-1
		for l < r {
			m := l + (r-l)>>1

			log.Printf("%d   %d>%d | %d>%d | %d>%d ", target, l, nums[l], m, nums[m], r, nums[r])

			if nums[l] <= nums[m] {
				if nums[l] <= target && target <= nums[m] {
					r = m
				} else {
					l = m + 1
				}

			} else {
				if nums[m] <= target && target <= nums[r] {
					l = m
				} else {
					r = m - 1
				}
			}

		}

		if l < len(nums) && nums[l] == target {
			return l
		}
		return -1
	}

	log.Print("4 ?= ", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	log.Print("-1 ?= ", search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	log.Print("-1 ?= ", search([]int{5, 5, 6, 7, 1, 1, 2}, 0))
}

// 153m Find Minimum in Rotated Sorted Array
func Test153(t *testing.T) {
	findMin := func(nums []int) int {
		l, r := 0, len(nums)-1

		for l < r {
			m := l + (r-l)>>1
			log.Printf("%d>%d | %d>%d | %d>%d", l, nums[l], m, nums[m], r, nums[r])

			if nums[m] < nums[r] {
				r = m
			} else {
				l = m + 1
			}
		}

		return nums[l]
	}

	log.Print("1 ?= ", findMin([]int{5, 1, 2, 3, 4}))
	log.Print("1 ?= ", findMin([]int{3, 4, 5, 1, 2}))
	log.Print("0 ?= ", findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	log.Print("0 ?= ", findMin([]int{5, 6, 7, 0, 1, 2, 3, 4}))
	log.Print("0 ?= ", findMin([]int{0, 1, 2, 3, 4, 5, 6, 7}))
	log.Print("0 ?= ", findMin([]int{1, 2, 3, 4, 5, 6, 7, 0}))
}

// 4h Median of Two Sorted Arrays
func Test4(t *testing.T) {
	findMedianSortedArrays := func(nums1, nums2 []int) float64 {
		var kValue func(int, int, int, int, int) int
		kValue = func(k, l1, r1, l2, r2 int) int {
			if r1 < l1 {
				return nums2[k-l1]
			}
			if r2 < l2 {
				return nums1[k-l2]
			}

			m1 := l1 + (r1-l1)>>1
			m2 := l2 + (r2-l2)>>1
			v1, v2 := nums1[m1], nums2[m2]

			log.Printf("%d   %d %d>%d %d   %d %d>%d %d   %v %v", k, l1, m1, v1, r1, l2, m2, v2, r2, nums1, nums2)

			if m1+m2 < k {
				if v1 > v2 {
					return kValue(k, l1, r1, m2+1, r2)
				} else {
					return kValue(k, m1+1, r2, l2, r2)
				}
			} else {
				if v1 > v2 {
					return kValue(k, l1, m1-1, l2, r2)
				} else {
					return kValue(k, l1, r1, l2, m2-1)
				}
			}
		}

		ln1, ln2 := len(nums1), len(nums2)
		ln := ln1 + ln2
		v := kValue(ln>>1, 0, ln1-1, 0, ln2-1)
		if ln&1 == 1 {
			return float64(v)
		}

		v += kValue(ln>>1-1, 0, ln1-1, 0, ln2-1)
		return float64(v) / float64(2)
	}

	log.Print("2 ?= ", findMedianSortedArrays([]int{1, 3}, []int{2}))
	log.Print("3.5 ?= ", findMedianSortedArrays([]int{1, 2, 3, 4}, []int{2, 4, 5, 5}))
}

// 441 Arranging Coins
func Test441(t *testing.T) {
	arrangeCoins := func(n int) int {
		triangular := func(v int) int {
			if v&1 == 1 {
				return v * (v + 1) / 2
			}
			return v / 2 * (v + 1)
		}

		l, r := 1, n+1
		for l < r {
			m := l + (r-l)>>1
			if triangular(m) > n {
				r = m
			} else {
				l = m + 1
			}
		}

		return l - 1
	}

	log.Print("1 ?= ", arrangeCoins(1))
	log.Print("2 ?= ", arrangeCoins(5))
	log.Print("3 ?= ", arrangeCoins(8))
	log.Print("6 ?= ", arrangeCoins(21))
	log.Print("6 ?= ", arrangeCoins(22))
	log.Print(" ?= ", arrangeCoins(1000_000_000))
}

// 1351 Count Negative Numbers in Sorted Matrix
func Test1351(t *testing.T) {
	countNegatives := func(grid [][]int) int {
		count := 0
		i, j := 0, len(grid[0])-1
		for i < len(grid) && 0 <= j {
			if grid[i][j] < 0 {
				count += len(grid) - i
				j--
			} else {
				i++
			}
		}
		return count
	}

	log.Print("8 ?= ", countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
}

// 1539 Kth Missing Positive Number
func Test1539(t *testing.T) {
	findKthPositive := func(arr []int, k int) int {
		kV := 1
		for k > 0 {
			l, r := 0, len(arr)
			for l < r {
				m := l + (r-l)>>1
				if arr[m] >= kV {
					r = m
				} else {
					l = m + 1
				}
			}
			if l == len(arr) || arr[l] != kV {
				k--
			}
			kV++
		}
		return kV - 1
	}

	log.Print("9 ?= ", findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	log.Print("6 ?= ", findKthPositive([]int{1, 2, 3, 4}, 2))
}
