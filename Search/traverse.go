package tgraph

import (
	"container/list"
	"fmt"
	"log"
	"math"
)

func init() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (o ListNode) String() string {
	b := '+'
	if o.Next == nil {
		b = '-'
	}
	return fmt.Sprintf("{%d %c}", o.Val, b)
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (o TreeNode) String() string {
	l, r := '+', '+'
	if o.Left == nil {
		l = '-'
	}
	if o.Right == nil {
		r = '-'
	}
	return fmt.Sprintf("{%d %c %c}", o.Val, l, r)
}

// 236
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if p == root || q == root {
		return root
	}

	l := LowestCommonAncestor(root.Left, p, q)
	r := LowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}

// 1161
func MaxLevelSum(root *TreeNode) int {
	Q := []*TreeNode{}
	Q = append(Q, root)

	l := 0
	xs, xl := math.MinInt, 0
	for len(Q) > 0 {
		l++
		ls := 0

		log.Print(l, Q)

		for lsize := len(Q); lsize > 0; lsize-- {
			n := Q[0]
			Q = Q[1:]

			ls += n.Val

			if n.Left != nil {
				Q = append(Q, n.Left)
			}
			if n.Right != nil {
				Q = append(Q, n.Right)
			}
		}

		if ls > xs {
			xs, xl = ls, l
		}
	}

	return xl
}

// 1372
func LongestZigZag(root *TreeNode) int {
	type E struct {
		Node   *TreeNode
		lZ, rZ int
	}
	Q := list.New()

	v := 0
	Q.PushBack(E{root, 0, 0})
	for Q.Len() > 0 {
		e := Q.Remove(Q.Back()).(E)
		n := e.Node

		if n.Left != nil {
			e.rZ++
			if e.rZ > v {
				v = e.rZ
			}
			Q.PushBack(E{n.Left, e.rZ, 0})
		}
		if n.Right != nil {
			e.lZ++
			if e.lZ > v {
				v = e.lZ
			}
			Q.PushBack(E{n.Right, 0, e.lZ})
		}
	}
	return v
}

// 437
func PathSum3(root *TreeNode, targetSum int) int {
	Q, S := []*TreeNode{}, [][]int{}
	Q, S = append(Q, root), append(S, []int{})

	v := 0
	for len(Q) > 0 {
		n, vs := Q[len(Q)-1], S[len(S)-1]
		Q, S = Q[:len(Q)-1], S[:len(S)-1]

		vs = append(vs, n.Val)
		for s, i := 0, len(vs)-1; i >= 0; i-- {
			s += vs[i]
			if s == targetSum {
				v++
			}
		}

		if n.Left != nil {
			Q, S = append(Q, n.Left), append(S, vs)
		}
		if n.Right != nil {
			Q, S = append(Q, n.Right), append(S, vs)
		}
	}
	return v
}

func DFS1448(root *TreeNode) int {
	type E struct {
		Node *TreeNode
		xVal int
	}
	Q := list.New()

	v := 0
	Q.PushBack(E{root, root.Val})
	for Q.Len() > 0 {
		e := Q.Remove(Q.Back()).(E)
		log.Printf("%+v", e)
		n, x := e.Node, e.xVal
		if n.Val >= x {
			v++
			x = n.Val
		}

		if n.Left != nil {
			Q.PushBack(E{n.Left, x})
		}
		if n.Right != nil {
			Q.PushBack(E{n.Right, x})
		}
	}
	return v
}

// 1609
func IsEvenOddTree(root *TreeNode) bool {
	Q, L := []*TreeNode{}, []int{}
	Q, L = append(Q, root), append(L, 0)

	var prv int
	h := -1
	for len(Q) > 0 {
		n, l := Q[0], L[0]
		Q, L = Q[1:], L[1:]

		if l > h {
			h = l
		} else {
			if l&1 == 0 && prv >= n.Val {
				return false
			}
			if l&1 == 1 && prv <= n.Val {
				return false
			}
		}
		if l&1 == n.Val&1 {
			return false
		}
		prv = n.Val

		if n.Left != nil {
			Q, L = append(Q, n.Left), append(L, l+1)
		}
		if n.Right != nil {
			Q, L = append(Q, n.Right), append(L, l+1)
		}
	}
	return true
}
