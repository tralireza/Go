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
}
