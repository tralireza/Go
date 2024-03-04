package lheap

import (
	"log"
	"slices"
)

// 2542
func MaxScore(nums1, nums2 []int, k int) int {

	return 0
}

// 948
func BagOfTokensScore(tokens []int, power int) int {
	slices.Sort(tokens)
	score := 0
	l, r := 0, len(tokens)-1
	for l <= r {
		log.Printf("|> (%d,%d) %d %d", l, r, score, power)

		if power >= tokens[l] {
			score++
			power -= tokens[l]
			l++
		} else if score > 0 && l < r {
			score--
			power += tokens[r]
			r--
		} else {
			return score
		}

		log.Printf("<| (%d,%d) %d %d", l, r, score, power)
	}
	return score
}
