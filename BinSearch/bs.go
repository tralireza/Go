package bsearch

import (
	"log"
	"slices"
)

func init() {
	log.SetFlags(0)
	log.Print("> binSearch")
}

func RightBinSearch(Sdup []int, x int) int {
	l, r := 0, len(Sdup)
	for l < r {
		m := l + (r-l)/2
		if Sdup[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return r - 1
}

func LeftBinSearch(Sdup []int, x int) int {
	l, r := 0, len(Sdup)
	for l < r {
		m := l + (r-l)/2
		if Sdup[m] < x {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
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
	slices.Sort(potions)

	pairs := make([]int, len(spells))
	for i, spell := range spells {
		l, r := 0, len(potions)-1
		for l <= r {
			m := l + (r-l)>>1
			if int64(spell*potions[m]) < success {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		pairs[i] = len(potions) - l
	}
	return pairs
}
