package bsearch

import (
	"log"
	"slices"
)

func init() {
	log.SetFlags(0)
	log.Print("> binSearch")
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
	log.Print(potions)

	pairs := make([]int, len(spells))
	for i, v := range spells {
		x := (success + int64(v) - 1) / int64(v)
		if x > 1 {
			l := leftBSearch(potions, int(x))
			pairs[i] = len(potions) - l
			if l == len(potions)-1 && int64(potions[l]) < x {
				log.Printf("%d: %d %d", l, potions[l], x)

				pairs[i]--
			}
		} else {
			pairs[i] = len(potions)
		}
	}
	return pairs
}
