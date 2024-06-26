package Graph

import (
	"log"
	"testing"
)

func init() {
	log.Print("> Graph>Basic")
}

func TestSetup(t *testing.T) {
	dg := GDirected(7, [][2]byte{{0, 1}, {2, 3}, {1, 3}, {4, 5}, {5, 6}, {3, 6}})
	log.Print(dg)
	ug := GUndirected(7, [][2]byte{{0, 1}, {2, 3}, {1, 3}, {4, 5}, {5, 6}, {3, 6}})
	log.Print(ug)
}
