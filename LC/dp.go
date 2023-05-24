package lc

import "log"

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
	Rows, Cols := len(grid), len(grid[0])
	D := make([][]int, Rows)
	for r := range D {
		D[r] = make([]int, Cols)
	}

	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			if r == 0 {
				if c > 0 {
					D[r][c] = D[r][c-1]
				}
				D[r][c] += grid[r][c]
			} else if c == 0 {
				D[r][c] = D[r-1][c] + grid[r][c]
			} else {
				D[r][c] = min(D[r][c-1], D[r-1][c]) + grid[r][c]
			}
		}
	}
	log.Print(D)

	return D[Rows-1][Cols-1]
}
