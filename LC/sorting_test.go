package lc

import (
	"fmt"
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
	// 3Way Partitioning
	sortColors := func(flags []int) {
		l, r := 0, len(flags)-1
		i := 0
		for i <= r {
			if flags[i] < WHITE {
				flags[i], flags[l] = flags[l], RED
				l++
				i++
			} else if flags[i] > WHITE {
				flags[i], flags[r] = flags[r], BLUE
				r--
			} else {
				i++
			}
		}
	}

	draw := func(flags []int) {
		fmt.Print("[ ")
		for _, f := range flags {
			switch f {
			case RED:
				fmt.Print("R")
			case WHITE:
				fmt.Print("W")
			case BLUE:
				fmt.Print("B")
			}
		}
		fmt.Println(" ]")
	}

	vs := []int{RED, WHITE, BLUE, BLUE, RED, BLUE, WHITE, BLUE, WHITE, RED, WHITE, RED}
	sortColors(vs)
	draw(vs)

	for range 4096 {
		flags := []int{RED, WHITE, BLUE}
		for range rand.Intn(4096) {
			flags = append(flags, rand.Intn(3))
		}
		sortColors(flags)
		for i := range flags[:len(flags)-1] {
			if flags[i] > flags[i+1] {
				t.Fatalf("Bad Dutch(%d) Flag! %d: %d %d", len(flags), i, flags[i], flags[i+1])
			}
		}
	}
}
