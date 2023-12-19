package lc

import (
	"log"
	"math/rand"
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

	// Red, White, Green: Dutch national flag Partitioning Sort
	sortColors := func(flags []int) {
		l, r := 0, len(flags)-1
		i := 0
		for i <= r {
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
	}

	vs := []int{RED, WHITE, BLUE, RED, BLUE, WHITE, BLUE, WHITE, RED, WHITE, RED}
	sortColors(vs)
	log.Print(vs)

	flags := []int{}
	for range 8196 {
		flags = append(flags, rand.Intn(3))
	}
	sortColors(flags)
	for i := range flags[:len(flags)-1] {
		if flags[i] > flags[i+1] {
			t.Fatalf("Bad Dutch flag! %d: %d %d", i, flags[i], flags[i+1])
		}
	}
}
