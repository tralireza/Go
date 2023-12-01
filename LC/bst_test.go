package lc

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.Print("> Binary Search Tree")
	_ = eBT{}
}

type eBT struct {
	Val         any
	Left, Right *eBT
}

// InOrder walk with Stack, non-recursive
func TestInOrder(t *testing.T) {
	type Tree = eBT

	rInOrder := func(root *Tree, visit func(*Tree)) {
		var walk func(*Tree, func(*Tree))
		walk = func(n *Tree, visit func(*Tree)) {
			if n == nil {
				return
			}

			walk(n.Left, visit)
			visit(n)
			walk(n.Right, visit)
		}

		walk(root, visit)
	}

	iInOrder := func(root *Tree, visit func(*Tree)) {}

	visit := func(n *Tree) {
		l, r := '-', '-'
		if n.Left != nil {
			l = '*'
		}
		if n.Right != nil {
			r = '*'
		}
		fmt.Printf("{%c %q %c} ", l, n.Val, r)
	}

	type T = Tree
	r := &T{'a', &T{'b', &T{Val: 'c'}, nil}, &T{'d', nil, &T{Val: 'e'}}}
	rInOrder(r, visit)
	fmt.Println()
	iInOrder(r, visit)
	fmt.Println()
}
