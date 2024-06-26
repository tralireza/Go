package lc

import (
	"log"
	"slices"
	"testing"
	"time"
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

	numberOfSubarrays2 := func(nums []int, k int) int {
		atMost := func(k int) int {
			x := 0
			l, r := 0, 0
			for r < len(nums) {
				k -= nums[r] & 1
				for k < 0 && l <= r {
					k += nums[l] & 1
					l++
				}
				x += r - l + 1
				r++
			}
			return x
		}

		return atMost(k) - atMost(k-1)
	}

	// 3 odd numbers in subarray: [1 7 3] [1 7 3 4] [7 3 4 5]
	log.Print("3 ?= ", numberOfSubarrays([]int{1, 7, 3, 4, 5}, 3))

	log.Print("3 ?= ", numberOfSubarrays2([]int{1, 7, 3, 4, 5}, 3))
	log.Print("6 ?= ", numberOfSubarrays2([]int{1, 7, 3, 4, 5}, 1))
	log.Print("1 ?= ", numberOfSubarrays2([]int{1, 7, 3, 4, 5}, 0))
}

// 17 Letter Combinations
func Test17(t *testing.T) {
	letterCombinations := func(digits string) []string {
		if len(digits) == 0 {
			return nil
		}

		dMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
		type Qe struct {
			letter      byte
			idx         int
			combination []byte
		}

		Q := []Qe{}
		letters := dMap[digits[0]-'2']
		for i := 0; i < len(letters); i++ {
			Q = append(Q, Qe{letters[i], 0, []byte{}})
		}

		cs := []string{}
		for len(Q) > 0 {
			log.Print(Q)

			e := Q[len(Q)-1]
			Q = Q[:len(Q)-1]

			e.idx++
			if e.idx < len(digits) {
				letters := dMap[digits[e.idx]-'2']
				for i := 0; i < len(letters); i++ {
					Q = append(Q, Qe{letters[i], e.idx, append(e.combination, e.letter)})
				}
			} else {
				cs = append(cs, string(append(e.combination, e.letter)))
			}
		}
		log.Print(cs)
		return cs
	}

	log.Print("9 ?= ", len(letterCombinations("23")))
	log.Print("0 ?= ", len(letterCombinations("")))
	log.Print("4 ?= ", len(letterCombinations("9")))
	log.Print("36 ?= ", len(letterCombinations("273")))
}

// 216m
func Test216(t *testing.T) {
	combinationSum3 := func(k int, n int) [][]int {
		cs := [][]int{}

		var dfs func(int, int, []int, int)
		dfs = func(k int, start int, pcs []int, n int) {
			if k == 0 {
				if n == 0 {
					cs = append(cs, pcs)
				}
				return
			}
			for i := start; i <= 9; i++ {
				dfs(k-1, i+1, append([]int{i}, pcs...), n-i)
			}
		}

		dfs(k, 1, []int{}, n)
		return cs
	}

	log.Printf("%v", combinationSum3(3, 7))
	log.Printf("%v", combinationSum3(3, 9))
	log.Printf("%v", combinationSum3(4, 24))
}

// 2962m Count Subarrays Where Max Element Appears at Least K Times
func Test2962(t *testing.T) {
	countSubarrays := func(nums []int, k int) int64 {
		mxVal := slices.Max(nums)
		count := int64(0)

		l, r := 0, 0
		frq := 0
		for ; r < len(nums); r++ {
			if nums[r] == mxVal {
				frq++
			}
			for ; frq >= k; l++ {
				count += int64(len(nums) - r)
				if nums[l] == mxVal {
					frq--
				}
			}
		}

		return count
	}

	fOn := func(nums []int, k int) int64 {
		mxVal := slices.Max(nums)
		count := int64(0)

		mxIdx := []int{}
		for i, n := range nums {
			if n == mxVal {
				mxIdx = append(mxIdx, i)
			}
			if len(mxIdx) >= k {
				count += int64(mxIdx[len(mxIdx)-k] + 1)
			}
		}

		return count
	}

	fOn2 := func(nums []int, k int) int64 {
		mxVal := slices.Max(nums)
		count := int64(0)

		frq := 0
		for r, n := range nums {
			if n == mxVal {
				frq++
			}
			for l, frq := 0, frq; l <= r && frq >= k; l++ {
				count++
				if nums[l] == mxVal {
					frq--
				}
			}
		}

		return count
	}

	for _, f := range []func([]int, int) int64{countSubarrays, fOn, fOn2} {
		ts := time.Now()
		log.Print("6 ?= ", f([]int{1, 3, 2, 3, 3}, 2), " ", time.Since(ts))
	}
}

// 992h Subarrays with K Different Integers
func Test992(t *testing.T) {
	subarraysWithKDistinct_1Pass := func(nums []int, k int) int {
		frq := map[int]int{}

		count := 0
		countCur := 0
		l, r := 0, 0
		for r < len(nums) {
			frq[nums[r]]++
			if frq[nums[r]] == 1 {
				k--
			}

			if k < 0 {
				frq[nums[l]]--
				if frq[nums[l]] == 0 {
					k++
				}
				countCur = 0
				l++
			}

			if k == 0 {
				for frq[nums[l]] > 1 {
					frq[nums[l]]--
					countCur++
					l++
				}
				count += countCur + 1
			}

			r++
		}

		return count
	}

	subarraysWithKDistinct := func(nums []int, k int) int {
		atMost := func(k int) int {
			frq := map[int]int{}
			curK, count := 0, 0

			l := 0
			for r, n := range nums {
				if frq[n] == 0 {
					curK++
				}
				frq[n]++

				for curK > k {
					if frq[nums[l]] == 1 {
						curK--
					}
					frq[nums[l]]--
					l++
				}

				count += r - l + 1
			}

			return count
		}

		return atMost(k) - atMost(k-1)
	}

	log.Print("7 ?= ", subarraysWithKDistinct_1Pass([]int{1, 2, 1, 2, 3}, 2))
	log.Print("7 ?= ", subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
	log.Print("3 ?= ", subarraysWithKDistinct([]int{1, 2, 1, 4, 3}, 3))
}

// 56m Merge Intervals
func Test56(t *testing.T) {
	merge := func(intervals [][]int) [][]int {
		slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })

		ms := [][]int{intervals[0]}
		for _, i := range intervals[1:] {
			if ms[len(ms)-1][1] >= i[0] {
				ms[len(ms)-1][1] = max(ms[len(ms)-1][1], i[1])
			} else {
				ms = append(ms, i)
			}
		}
		return ms
	}

	log.Print(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	log.Print(merge([][]int{{1, 4}, {4, 5}}))
}

func TestPerms(t *testing.T) {
	counter := 0
	var Perms func([]int, int)
	Perms = func(arr []int, k int) {
		if k == len(arr)-1 {
			counter++
			log.Printf("%3d -> %v", counter, arr)
			return
		}

		for i := k; i < len(arr); i++ {
			arr[i], arr[k] = arr[k], arr[i]
			Perms(arr, k+1)
			arr[i], arr[k] = arr[k], arr[i]
		}
	}

	Perms([]int{0, 1, 2, 3, 4, 5}, 0)
	log.Print(counter)

	X := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, x := range X {
		log.Print(i, X[:i], x, X[i+1:])
	}
}

// 2444h Count Subarrays with Fixed Bounds
func Test2444(t *testing.T) {
	countSubarrays := func(nums []int, minK int, maxK int) int64 {
		count := int64(0)
		mn, mx := -1, -1

		l, r := 0, 0
		for r < len(nums) {
			n := nums[r]

			if n == minK {
				mn = r
			}
			if n == maxK {
				mx = r
			}
			if n < minK || maxK < n {
				mx = -1
				mn = -1
				l = r + 1
			}

			if mx >= 0 && mn >= 0 {
				count += int64(min(mx, mn) - l + 1)
			}

			r++
		}

		return count
	}

	log.Print("2 ?= ", countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
	log.Print("10 ?= ", countSubarrays([]int{1, 1, 1, 1}, 1, 1))
}
