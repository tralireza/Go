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

	log.Print(D)

	return longest
}
