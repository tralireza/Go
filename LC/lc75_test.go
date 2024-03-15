package lc

import (
	"log"
	"testing"
)

func init() {
	log.Print("> 75")
}

func TestDecodeString(t *testing.T) {
	d := "2[a2[h]]1[cd]ij10[p]"
	s := decodeString(d)
	log.Printf("%s -> %s", d, s)
	if s != "ahhahhcdijpppppppppp" {
		t.Fatal()
	}

	for _, s := range []string{"0", "2[3[4[1[a]b]c]d]", "1[ab2[cd]3[xy]z", "2[3[x]co1[h]0]"} {
		log.Printf("%s -> %s", s, decodeString(s))
	}
}

// 62m Unique Paths
func Test62(t *testing.T) {
	uniquePathsOmn := func(m, n int) int {
		P := make([][]int, m)
		for r := range P {
			P[r] = make([]int, n)
		}

		for c := range P[0] {
			P[0][c] = 1
		}

		for r := 1; r < m; r++ {
			P[r][0] = 1
			for c := 1; c < n; c++ {
				P[r][c] = P[r-1][c] + P[r][c-1]
			}
		}
		return P[m-1][n-1]
	}

	uniquePathsOn := func(m, n int) int {
		row := make([]int, n)
		for c := range row {
			row[c] = 1
		}

		for r := 1; r < m; r++ {
			row[0] = 1
			for c := 1; c < n; c++ {
				row[c] = row[c-1] + row[c]
			}
		}
		return row[n-1]
	}

	log.Print("28 ?= ", uniquePathsOmn(3, 7))
	log.Print("28 ?= ", uniquePathsOn(3, 7))
}
