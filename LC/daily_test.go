package lc

import (
	"log"
	"testing"
)

func init() {}

// 930m
func Test930(t *testing.T) {
	countSubarraysWithSum := func(nums []int, goal int) int {
		frq := map[int]int{}

		x, pfxSum := 0, 0
		for _, n := range nums {
			pfxSum += n
			if pfxSum == goal {
				x++
			}

			if f, ok := frq[pfxSum-goal]; ok {
				x += f
			}

			frq[pfxSum]++

			log.Print(frq)
		}

		return x
	}

	log.Print("4 ?= ", countSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	log.Print("8 ?= ", countSubarraysWithSum([]int{1, 0, 1, 0, 1}, 1))
	log.Print("15 ?= ", countSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}
