package bsearch

import (
	"log"
	"slices"
	"testing"
)

func Test2300(t *testing.T) {
	for _, v := range [][][]int{
		{[]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, []int{7}, []int{4, 0, 3}},
		{[]int{1, 2, 3}, []int{4, 5}, []int{9}, []int{0, 1, 2}},
		{[]int{16, 13, 7, 36, 16, 25, 22, 18, 29}, []int{38, 25}, []int{223}, []int{2, 2, 1, 2, 2, 2, 2, 2, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7}, []int{25}, []int{0, 0, 0, 1, 3, 3, 4}},
		{[]int{3, 1, 2, 16}, []int{8, 5, 8}, []int{16}, []int{2, 0, 2, 3}}} {

		spells, potions, success := v[0], v[1], int64(v[2][0])
		if pairs := SuccessfulPairs(spells, potions, success); slices.Compare(pairs, v[3]) != 0 {
			t.Fatalf("Wrong pairs! %v != %v", pairs, v[3])
		}
		log.Printf("+ %v %v  -%d->  %v", spells, potions, success, v[3])
	}
}
