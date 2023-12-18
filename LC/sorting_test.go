package lc

import (
	"log"
	"testing"
)

func init() {
	log.Print("> Sorting")
}

// 75m Sort Colors
func Test75(t *testing.T) {
	const (
		RED int = iota
		WHITE
		BLUE
	)

	sortColors_1Pass := func(flags []int) {
		l, r := 0, len(flags)-1
		i := 0
		for i != r {
			if flags[i] == RED {
				if l == i {
					i++
				} else {
					flags[i], flags[l] = flags[l], RED
				}
				l++
			}

			if flags[i] == BLUE {
				flags[i], flags[r] = flags[r], BLUE
				r--
			}

			if flags[i] == WHITE {
				i++
			}
		}

		log.Print(flags)
	}

	sortColors := func(flags []int) {
		frq := [3]int{}
		for _, f := range flags {
			frq[f]++
		}

		i := 0
		for f, count := range frq {
			for range count {
				flags[i] = f
				i++
			}
		}

		log.Print(flags)
	}

	sortColors([]int{0, 1, 2, 0, 2, 1, 2, 1, 0, 1, 0})
	sortColors_1Pass([]int{RED, WHITE, BLUE, RED, BLUE, WHITE, BLUE, WHITE, RED, WHITE, RED})
}
