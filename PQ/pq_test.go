package lpq

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(0)
}

func Test2462(t *testing.T) {
	for _, v := range [][][]int{{{17, 12, 10, 2, 7, 2, 11, 20, 8}, {3, 4, 11}}, {{1, 2, 4, 1}, {3, 3, 4}}} {
		costs := v[0]
		k, candidates, tcost := v[1][0], v[1][1], int64(v[1][2])

		if r := TotalCost(costs, k, candidates); r != tcost {
			t.Fatalf("Wront totalCost! must be %d != %d", r, tcost)
		}
		if r := TotalCost1(costs, k, candidates); r != tcost {
			t.Fatalf("Wront totalCost! must be %d != %d", r, tcost)
		}
	}
}

func minLength(s string) int {
	l, r := 0, len(s)-1
	for l < r && s[l] == s[r] {
		chr := s[l]
		for l <= r && chr == s[l] {
			l++
		}
		for l < r && chr == s[r] {
			r--
		}
	}
	return r - l + 1
}

func Test1750(t *testing.T) {
	for _, s := range []string{"aabccabba", "a", "aab", "aaabcccbacaabcbbba", "bb", "aaaca", "bbacabbb"} {
		log.Printf("%q -> %d", s, minLength(s))
	}
}
