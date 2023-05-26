package lc

import (
	"log"
	"math"
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
	x := math.MinInt
	for r := 1; r <= len(nums); r++ {
		for l := 0; l < r; l++ {
			log.Print(l, r)
			v := nums[l]
			for x := l + 1; x < r; x++ {
				v *= nums[x]
			}
			if v > x {
				x = v
			}
		}
	}
	return x
}
