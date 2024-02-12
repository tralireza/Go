package bsearch

import (
	"log"
	"slices"
)

func init() {
	log.SetFlags(0)
	log.Print("> binSearch")
}

func BinSearch2R(S []int, x int) int {
	l, r := 0, len(S)-1
	for l != r {
		m := l + (r-l)/2
		if S[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	if x == S[r] {
		return r
	}
	return -1
}

func BinSearch2(S []int, x int) int {
	l, r := 0, len(S)-1
	for l != r {
		m := l + (r+1-l)/2
		if S[m] > x {
			r = m - 1
		} else {
			l = m
		}
	}
	if x == S[l] {
		return l
	}
	return -1
}

func BinSearch3(S []int, x int) int {
	l, r := 0, len(S)-1
	for l <= r {
		m := l + (r-l)/2
		if x < S[m] {
			r = m - 1
		} else if x > S[m] {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}

// 2300
func SuccessfulPairs(spells []int, potions []int, success int64) []int {
	leftBSearch := func(Sdup []int, x int) int {
		l, r := 0, len(Sdup)-1
		for l < r {
			m := (l + r) / 2
			if Sdup[m] < x {
				l = m + 1
			} else {
				r = m
			}
		}
		return l
	}

	slices.Sort(potions)

	pairs := make([]int, len(spells))
	for i, v := range spells {
		pairs[i] = len(potions)
		x := (success + int64(v) - 1) / int64(v)
		if x > 1 {
			l := leftBSearch(potions, int(x))
			pairs[i] -= l

			if l == len(potions)-1 && int64(potions[l]) < x {
				pairs[i]--
			}
		}
	}
	return pairs
}
