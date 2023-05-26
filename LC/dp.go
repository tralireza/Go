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
