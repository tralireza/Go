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
		phoneMap := func(number byte) string {
			switch number {
			case '2':
				return "abc"
			case '3':
				return "def"
			case '4':
				return "ghi"
			case '5':
				return "jkl"
			case '6':
				return "mno"
			case '7':
				return "pqrs"
			case '8':
				return "tuv"
			case '9':
				return "wxyz"
			default:
				return ""
			}
		}

		if len(digits) == 0 {
			return nil
		}

		Q, I, B := []byte{}, []int{}, [][]byte{}
		letters := phoneMap(digits[0])
		for i := 0; i < len(letters); i++ {
			Q, I, B = append(Q, letters[i]), append(I, 0), append(B, []byte{})
		}

		cs := []string{}

		for len(Q) > 0 {
			log.Print(Q, I, B)
			b, idx, bs := Q[len(Q)-1], I[len(I)-1], B[len(B)-1]
			Q, I, B = Q[:len(Q)-1], I[:len(I)-1], B[:len(B)-1]

			idx++
			if idx < len(digits) {
				letters := phoneMap(digits[idx])
				for i := 0; i < len(letters); i++ {
					Q, I, B = append(Q, letters[i]), append(I, idx), append(B, append(bs, b))
				}
			} else {
				cs = append(cs, string(append(bs, b)))
			}
		}

		log.Print(cs)
		return cs
	}

	log.Print("9 ?= ", len(letterCombinations("23")))
	log.Print("0 ?= ", len(letterCombinations("")))
	log.Print("0 ?= ", len(letterCombinations("1")))
	log.Print("4 ?= ", len(letterCombinations("9")))
}
