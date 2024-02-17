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

type tArr []int

func New() tArr {
	return tArr([]int{1, 3, 4, 8, 10, 12, 17, 18, 19, 20, 21, 23, 24, 25, 26, 27, 31, 33, 38, 39, 41, 43, 44, 45, 46, 51, 55, 56, 59})
}
func (a tArr) Draw() {
	for i := range a {
		fmt.Printf("|%2d", i)
	}
	fmt.Println("|")
	for _, v := range a {
		fmt.Printf("|%2d", v)
	}
	fmt.Println("|")
}
func (a tArr) Dups() {
	a[1], a[2] = 1, 1
	a[18], a[19], a[20] = 33, 33, 33
}

func TestRankBinSearch(t *testing.T) {
	S := New()
	S.Dups()
	S.Draw()
	for _, v := range []int{0, 1, 8, 29, 43, 55, 59, 60} {
		l, r := LeftBinSearch(S, v), RightBinSearch(S, v)
		log.Printf("%3d -> Rank(L): %2d (found? %-5t)   | Rank(R): %2d (%2d) (found? %-5t)", v,
			l, l < len(S) && S[l] == v,
			len(S)-r-1, r, r > 0 && S[r] == v)
	}
}

func TestBinSearch(t *testing.T) {
	S := New()
	for _, v := range []int{1, 59, 33, 31, 2, 58, 0, 60} {
		log.Printf("3. %3v -> % 3d   | 2. % 3d   | 2r. % 3d", v, BinSearch3(S, v), BinSearch2(S, v), BinSearch2R(S, v))
	}
}

func BenchmarkBinSearch(b *testing.B) {
	S := New()
	for i := 0; i < b.N; i++ {
		slices.BinarySearch(S, 41)
		slices.BinarySearch(S, 2)
	}
}

func BenchmarkBinSearch2(b *testing.B) {
	S := New()
	for i := 0; i < b.N; i++ {
		BinSearch2(S, 41)
		BinSearch2(S, 2)
	}
}

func BenchmarkBinSearch3(b *testing.B) {
	S := New()
	for i := 0; i < b.N; i++ {
		BinSearch3(S, 41)
		BinSearch3(S, 2)
	}
}
