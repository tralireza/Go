package lc

import (
	"log"
)

// 5m Longest Palindromic Substring
func longestPalindrome(s string) string {
	// D[length, start] -> longest
	D := make([][]int, len(s)+1)
	for l := range D {
		D[l] = make([]int, len(s))
	}

	for start := range D[1] {
		D[1][start] = 1
	}
	longest := s[:1]

	for l := 2; l <= len(s); l++ {
		for start := 0; start+l <= len(s); start++ {
			if s[start] == s[start+l-1] && D[l-2][start+1] == l-2 {
				D[l][start] = l

				if len(longest) < l {
					longest = s[start : start+l]
				}
			}
		}
	}

	log.Print("D: ", D)

	return longest
}

// 64m Minimum Path Sum
func minPathSum(grid [][]int) int {
	for c := range grid[0][1:] {
		grid[0][c+1] += grid[0][c]
	}
	for r := range grid[1:] {
		grid[r+1][0] += grid[r][0]
	}

	for r := range grid[1:] {
		for c := range grid[r+1][1:] {
			grid[r+1][c+1] = min(grid[r+1][c], grid[r][c+1]) + grid[r+1][c+1]
		}
	}

	log.Print(grid)

	return grid[len(grid)-1][len(grid[0])-1]
}

// 152m Maximum Product Subarray
// Kadane's Maximum Sum -> augmented for Product
func maxProduct(nums []int) int {
	Max, lMax := nums[0], nums[0]
	lMin := nums[0]
	for _, n := range nums[1:] {
		if n < 0 {
			lMax, lMin = lMin, lMax
		}
		lMax = max(lMax*n, n)
		Max = max(lMax, Max)

		lMin = min(lMin*n, n)
	}

	return Max
}

// 300m Longest Increasing Subsequence
func lengthOfLIS(nums []int) int {
	liSub := []int{}

	liSub = append(liSub, nums[0])
	for _, n := range nums[1:] {
		if n > liSub[len(liSub)-1] {
			liSub = append(liSub, n)
			continue
		}

		l, r := 0, len(liSub)-1
		for l < r {
			m := l + (r-l)>>1 // l <= m < r
			if liSub[m] >= n {
				r = m
			} else {
				l = m + 1
			}
		}
		liSub[l] = n
	}

	return len(liSub)
}

// 322m Coin Change
func coinChange(coins []int, amount int) int {
	D := make([]int, amount+1)
	for m := 1; m <= amount; m++ {
		D[m] = -1
		for _, c := range coins {
			n := m - c
			if n >= 0 && D[n] >= 0 {
				if D[m] == -1 {
					D[m] = 1 + D[n]
				} else {
					D[m] = min(D[m], D[n]+1)
				}
			}
		}
	}
	return D[amount]
}

// 1143m Longest Common Subsequence
func longestCommonSubsequence(text1 string, text2 string) int {
	D := make([][]int, len(text1)+1)
	for r := range D {
		D[r] = make([]int, len(text2)+1)
	}

	for r := 1; r <= len(text1); r++ {
		for c := 1; c <= len(text2); c++ {
			if text1[r-1] == text2[c-1] {
				D[r][c] = D[r-1][c-1] + 1
			} else {
				D[r][c] = max(D[r-1][c], D[r][c-1])
			}
		}
	}
	return D[len(text1)][len(text2)]
}

// 416m Partition Equal Subset Sum
func canPartition(nums []int) bool {
	// 1 <= len(nums) <= 200 & 1 <= nums[i] <= 100
	frq, tSum := make([]int, 100+1), 0
	for _, n := range nums {
		tSum += n
		frq[n]++
	}

	if tSum&1 == 1 {
		return false
	}

	Val := []int{}
	for _, n := range nums {
		if n <= tSum/2 {
			Val = append(Val, n)
		}
	}

	N, W := len(Val), tSum/2
	Knapsack01 := make([][]int, N+1)
	for i := range Knapsack01 {
		Knapsack01[i] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if Val[i-1] > w {
				Knapsack01[i][w] = Knapsack01[i-1][w]
			} else {
				Knapsack01[i][w] = max(Knapsack01[i-1][w], Knapsack01[i-1][w-Val[i-1]]+Val[i-1])
			}
		}
	}

	return Knapsack01[N][W] == W
}

// 72m Edit Distance
func minDistance(word1 string, word2 string) int {
	D := make([][]int, len(word1)+1)
	for r := range D {
		D[r] = make([]int, len(word2)+1)
	}
	for r := range D {
		D[r][0] = r
	}
	for c := range D[0] {
		D[0][c] = c
	}

	for r := 1; r <= len(word1); r++ {
		for c := 1; c <= len(word2); c++ {
			if word1[r-1] == word2[c-1] {
				D[r][c] = D[r-1][c-1]
			} else {
				D[r][c] = 1 + min(D[r-1][c-1], D[r-1][c], D[r][c-1])
			}
		}
	}

	return D[len(word1)][len(word2)]
}
