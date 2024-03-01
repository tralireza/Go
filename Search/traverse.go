package tgraph

import (
	"container/list"
	"fmt"
	"log"
)

func init() {}

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
		log.Print(e)
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
