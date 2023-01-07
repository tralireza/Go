package lc

import (
	"log"
	"testing"
)

func init() {
	log.Print("> Basic DS")
}

// 1614 Maximum Nesting Depth of Parentheses
func Test1614(t *testing.T) {
	maxDepth := func(s string) int {
		return 0
	}

	log.Print("3 ?= ", maxDepth("(1+(2*3)+((8)/4))+1"))
	log.Print("3 ?= ", maxDepth("(1)+((2))+(((3)))"))
}
