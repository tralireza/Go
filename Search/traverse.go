package tgraph

import "fmt"

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
