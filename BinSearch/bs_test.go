package bsearch

import (
	"log"
	"slices"
	"testing"
)

func Test2300(t *testing.T) {
	for _, v := range [][][]int{
		{[]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, []int{7}, []int{4, 0, 3}},
		{[]int{3, 1, 2, 16}, []int{8, 5, 8}, []int{16}, []int{2, 0, 2, 3}}} {

		spells, potions, success := v[0], v[1], int64(v[2][0])
		if pairs := SuccessfulPairs(spells, potions, success); slices.Compare(pairs, v[3]) != 0 {
			t.Fatalf("Wrong pairs! %v != %v", pairs, v[3])
		}
		log.Printf("+ %v %v  -%d->  %v", spells, potions, success, v[3])
	}
}
