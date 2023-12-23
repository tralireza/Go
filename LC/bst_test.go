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

	iInOrder := func(root *Tree, visit func(*Tree)) {
		n, S := root, []*Tree{}
		for len(S) > 0 || n != nil {
			if n != nil {
				S = append(S, n)
				n = n.Left
			} else {
				n, S = S[len(S)-1], S[:len(S)-1]

				visit(n)

				n = n.Right
			}
		}
	}

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
	r := &T{'1', &T{'2', &T{Val: '4'}, &T{Val: '5'}}, &T{'3', &T{Val: '6'}, nil}}
	rInOrder(r, visit)
	fmt.Println()
	iInOrder(r, visit)
	fmt.Println()
}

// 530 Minimum Absolute Difference in BST
func Test530(t *testing.T) {
	type TreeNode struct {
		Val         int
		Left, Right *TreeNode
	}

	minimumDifference := func(root *TreeNode) int {
		mnVal, prvVal := 100_001, -1

		S, n := []*TreeNode{}, root
		for len(S) > 0 || n != nil {
			if n != nil {
				S = append(S, n)
				n = n.Left
			} else {
				n, S = S[len(S)-1], S[:len(S)-1]

				if prvVal != -1 {
					mnVal = min(mnVal, n.Val-prvVal)
				}
				prvVal = n.Val

				n = n.Right
			}
		}

		return mnVal
	}

	type T = TreeNode
	log.Print("1 =? ", minimumDifference(&T{2, &T{Val: 1}, &T{Val: 3}}))
	log.Print("1 =? ", minimumDifference(&T{4, &T{2, nil, &T{Val: 3}}, &T{Val: 6}}))
}
