package bsearch

import (
	"log"
	"testing"
)

func Test2300(t *testing.T) {
	log.Print(SuccessfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, int64(7)))
	log.Print(SuccessfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, int64(16)))
}
