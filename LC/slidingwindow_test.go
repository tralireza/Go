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
		return 0
	}

	log.Print("6 ?= ", countSubarrays([]int{1, 3, 2, 3, 3}, 2))
}
