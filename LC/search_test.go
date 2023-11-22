package lc

import (
	"log"
	"testing"
)

func init() {
	log.Print("> BFS/DFS")
}

// 1091m Shortest Path in Binary Matrix
func Test1091(t *testing.T) {
	shortestPathBinaryMatrix := func(grid [][]int) int {
		return 0
	}

	gridSet := func() [][]int { return [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}} }

	log.Print("4 ?= ", shortestPathBinaryMatrix(gridSet()))
}
