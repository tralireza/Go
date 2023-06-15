package lpq

import (
	"log"
	"testing"
)

func init() {
	log.SetFlags(0)
}

func Test2462(t *testing.T) {
	log.Print(TotalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4))
	log.Print(TotalCost([]int{1, 2, 4, 1}, 3, 3))
}
