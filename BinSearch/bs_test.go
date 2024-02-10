package bsearch

import (
	"fmt"
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

var S = []int{1, 3, 4, 8, 10, 12, 17, 18, 19, 20, 21, 23, 24, 25, 26, 27, 31, 33, 38, 39, 41, 43, 44, 45, 46, 51, 55, 56, 59}

func TestBinSearch3(t *testing.T) {
	for i := range S {
		fmt.Printf("|%2d", i)
	}
	fmt.Println("|")
	for _, v := range S {
		fmt.Printf("|%2d", v)
	}
	fmt.Println("|")

	for _, v := range []int{1, 59, 33, 31, 2, 58, 0, 60} {
		log.Printf("3. %3v -> % 3d   | 2. %3v -> % 3d", v, BinSearch3(S, v), v, BinSearch2(S, v))
	}
}

func BenchmarkBinSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slices.BinarySearch(S, 41)
		slices.BinarySearch(S, 2)
	}
}

func BenchmarkBinSearch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinSearch2(S, 41)
		BinSearch2(S, 2)
	}
}

func BenchmarkBinSearch3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinSearch3(S, 41)
		BinSearch3(S, 2)
	}
}
