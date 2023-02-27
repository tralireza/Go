package lc

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"testing"
)

func init() {
	log.Print("> Binary Tree & BST")
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (o TreeNode) String() string {
	l, r := '-', '-'
	if o.Left != nil {
		l = '*'
	}
	if o.Right != nil {
		r = '*'
	}
	return fmt.Sprintf("{%c %d %c}", l, o.Val, r)
}

// 536m Construct Binary Tree from String
func Test536(t *testing.T) {
	str2Tree := func(s string) *TreeNode {
		S := []*TreeNode{}
		var n *TreeNode

		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '(':
				S = append(S, n)
			case ')':
				c := n
				n, S = S[len(S)-1], S[:len(S)-1]
				if n.Left != nil {
					n.Right = c
				} else {
					n.Left = c
				}
			default:
				start := i
				for ; i < len(s) && s[i] != '(' && s[i] != ')'; i++ {
				}
				v, _ := strconv.Atoi(s[start:i])
				n = &TreeNode{Val: v}
				i--
			}
		}

		return n
	}

	Draw := func(n *TreeNode) {
		Q, l := []*TreeNode{n}, 0
		for len(Q) > 0 {
			fmt.Printf("%d ", l)
			for range len(Q) {
				n, Q = Q[0], Q[1:]
				fmt.Print(n)
				if n.Left != nil {
					Q = append(Q, n.Left)
				}
				if n.Right != nil {
					Q = append(Q, n.Right)
				}
			}
			l++
			fmt.Print("\n")
		}
	}

	for _, s := range []string{"4(2(3)(1))(6(5))", "-4203", "7(3901)(4(-29(1)(891)))"} {
		Draw(str2Tree(s))
		log.Print("===")
	}
}

// InOrder walk with Stack, non-recursive
func TestInOrder(t *testing.T) {
	type Tree = TreeNode

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

// 404 Sum of Left Leaves
func Test404(t *testing.T) {
	type TreeNode struct {
		Val         int
		Left, Right *TreeNode
	}

	var sumOfLeftLeaves func(*TreeNode) int
	sumOfLeftLeaves = func(root *TreeNode) int {
		lsum := 0
		if root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				lsum += root.Left.Val
			} else {
				lsum += sumOfLeftLeaves(root.Left)
			}
		}
		if root.Right != nil {
			lsum += sumOfLeftLeaves(root.Right)
		}
		return lsum
	}

	type T = TreeNode

	flagLeft := func(root *T) int {
		var fsum func(*T, bool) int
		fsum = func(n *T, left bool) int {
			if n.Left == nil && n.Right == nil {
				if left {
					return n.Val
				}
				return 0
			}

			v := 0
			if n.Left != nil {
				v += fsum(n.Left, true)
			}
			if n.Right != nil {
				v += fsum(n.Right, false)
			}
			return v
		}

		return fsum(root, false)
	}

	for _, f := range []func(*T) int{sumOfLeftLeaves, flagLeft} {
		log.Print("24 ?= ", f(&T{3, &T{Val: 9}, &T{20, &T{Val: 15}, &T{Val: 7}}}))
		log.Print("0 ?= ", f(&T{Val: 3}))
	}
}

// 129m Sum Root to Leaf Numbers
func Test129(t *testing.T) {
	type TreeNode struct {
		Val         int
		Left, Right *TreeNode
	}

	sumNumbers := func(root *TreeNode) int {
		tsum := 0

		S, V := []*TreeNode{root}, []int{root.Val}
		var n *TreeNode
		var v int
		for len(S) > 0 {
			n, S = S[len(S)-1], S[:len(S)-1]
			v, V = V[len(V)-1], V[:len(V)-1]

			for _, c := range []*TreeNode{n.Left, n.Right} {
				if c != nil {
					S, V = append(S, c), append(V, 10*v+c.Val)
				}
			}

			log.Print(S, V)
			if n.Left == nil && n.Right == nil {
				tsum += v
			}
		}

		return tsum
	}

	recursive := func(root *TreeNode) int {
		tsum := 0

		var walk func(*TreeNode, int)
		walk = func(n *TreeNode, v int) {
			if n.Left == nil && n.Right == nil {
				tsum += 10*v + n.Val
			}

			if n.Left != nil {
				walk(n.Left, 10*v+n.Val)
			}
			if n.Right != nil {
				walk(n.Right, 10*v+n.Val)
			}
		}

		walk(root, 0)
		return tsum
	}

	type T = TreeNode
	for _, f := range []func(*TreeNode) int{sumNumbers, recursive} {
		log.Print("12(12) ?= ", f(&T{1, &T{Val: 2}, nil}))
		log.Print("25(12+13) ?= ", f(&T{1, &T{Val: 2}, &T{Val: 3}}))
		log.Print("1026(495+491+40) ?= ", f(&T{4, &T{9, &T{Val: 5}, &T{Val: 1}}, &T{Val: 0}}))
	}
}

// 988m Smallest Starting from Leaf
func Test988(t *testing.T) {
	type TreeNode struct {
		Val         int
		Left, Right *TreeNode
	}

	smallestFromLeaf := func(root *TreeNode) string {
		ms := ""

		var walk func(*TreeNode, string)
		walk = func(n *TreeNode, s string) {
			s = string('a'+byte(n.Val)) + s
			if n.Left == nil && n.Right == nil {
				if ms == "" || s < ms {
					ms = s
				}
			}

			if n.Left != nil {
				walk(n.Left, s)
			}
			if n.Right != nil {
				walk(n.Right, s)
			}
		}

		walk(root, "")
		return ms
	}

	type T = TreeNode
	log.Print("dba ?= ", smallestFromLeaf(&T{0, &T{1, &T{Val: 3}, &T{Val: 4}}, &T{2, &T{Val: 3}, &T{Val: 4}}}))
}

// 124h Binary Tree Maximum Path Sum
func Test124(t *testing.T) {
	type TreeNode struct {
		Val         int
		Left, Right *TreeNode
	}

	maxPathSum := func(root *TreeNode) int {
		x := math.MinInt

		var mxSum func(*TreeNode) int
		mxSum = func(n *TreeNode) int {
			if n == nil {
				return 0
			}

			ls, rs := max(0, mxSum(n.Left)), max(0, mxSum(n.Right))
			x = max(x, ls+n.Val+rs)
			return n.Val + max(ls, rs)
		}
		mxSum(root)

		return x
	}

	type T = TreeNode
	log.Print("42 ?= ", maxPathSum(&T{-10, &T{Val: 9}, &T{20, &T{Val: 15}, &T{Val: 7}}}))
	log.Print("5 ?= ", maxPathSum(&T{1, &T{Val: -5}, &T{Val: 4}}))
	log.Print("-3 ?= ", maxPathSum(&T{Val: -3}))
}

// 623m Add One Row to Tree
func Test623(t *testing.T) {
	addOneRow := func(root *TreeNode, val int, depth int) *TreeNode {
		if depth == 1 {
			return &TreeNode{val, root, nil}
		}

		Q := []*TreeNode{root}
		var n *TreeNode

		for depth-1 > 1 {
			for k := len(Q); k > 0; k-- {
				n, Q = Q[0], Q[1:]
				for _, v := range []*TreeNode{n.Left, n.Right} {
					if v != nil {
						Q = append(Q, v)
					}
				}
			}
			depth--
		}

		for len(Q) > 0 {
			n, Q = Q[0], Q[1:]
			n.Left = &TreeNode{val, n.Left, nil}
			n.Right = &TreeNode{val, nil, n.Right}
		}

		return root
	}

	recursive := func(root *TreeNode, val int, depth int) *TreeNode {
		if depth == 1 {
			return &TreeNode{Val: val, Left: root}
		}

		var preOrder func(*TreeNode, int)
		preOrder = func(n *TreeNode, depth int) {
			if n == nil {
				return
			}

			if depth == 1 {
				n.Left = &TreeNode{Val: val, Left: n.Left}
				n.Right = &TreeNode{Val: val, Right: n.Right}
				return
			}

			preOrder(n.Left, depth-1)
			preOrder(n.Right, depth-1)
		}

		preOrder(root, depth-1)
		return root
	}

	draw := func(n *TreeNode) {
		Q := []*TreeNode{}
		for len(Q) > 0 || n != nil {
			if n != nil {
				Q = append(Q, n)
				n = n.Left
			} else {
				n, Q = Q[len(Q)-1], Q[:len(Q)-1]

				l, r := '-', '-'
				if n.Left != nil {
					l = '*'
				}
				if n.Right != nil {
					r = '*'
				}
				fmt.Printf("{%c %d %c}", l, n.Val, r)

				n = n.Right
			}
		}
		fmt.Print("\n")
	}

	type T = TreeNode
	var r *T

	for _, f := range []func(*TreeNode, int, int) *TreeNode{recursive, addOneRow} {
		r = &T{Val: 1}
		draw(r)
		draw(f(r, 0, 1))
		log.Print("===")
		r = &T{1, &T{Val: 2}, &T{Val: 2}}
		draw(r)
		draw(f(r, 0, 2))
		log.Print("===")
	}
}

// 114m Flatten Binary Tree to Linked List
func Test114(t *testing.T) {
	flatten := func(root *TreeNode) {
		n := root

		for n != nil {
			if n.Left != nil {
				r := n.Left
				for r.Right != nil { // finding rightmost node of left child of n
					r = r.Right
				}
				r.Right = n.Right // right child of n is at right most of left child of n
				n.Right = n.Left  // move all to right child of n
				n.Left = nil
			}

			n = n.Right
		}
	}

	type T = TreeNode

	n := &T{1, &T{2, &T{Val: 3}, &T{Val: 4}}, &T{5, nil, &T{Val: 6}}}
	flatten(n)
	for ; n != nil; n = n.Right {
		fmt.Print(n)
	}
	fmt.Print("\n")
}

// 101 Symmetric Tree
func Test101(t *testing.T) {
	type TreeNode struct {
		Val         int
		Left, Right *TreeNode
	}

	isSymmetric := func(root *TreeNode) bool {
		var check func(l, r *TreeNode) bool
		check = func(l, r *TreeNode) bool {
			if l == nil && r == nil {
				return true
			}
			if l == nil || r == nil {
				return false
			}
			return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
		}

		return check(root.Left, root.Right)
	}

	type T = TreeNode

	log.Print("true ?= ", isSymmetric(&T{1, &T{2, &T{Val: 3}, &T{Val: 4}}, &T{2, &T{Val: 4}, &T{Val: 3}}}))
	log.Print("false ?= ", isSymmetric(&T{1, &T{2, nil, &T{Val: 3}}, &T{2, nil, &T{Val: 3}}}))
}

// 108 Convert Sorted Array to Binary Search Tree
func Test108(t *testing.T) {
	type TreeNode struct {
		Val         int
		Left, Right *TreeNode
	}

	var sortedArrayToBST func([]int) *TreeNode
	sortedArrayToBST = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}

		m := len(nums) / 2
		return &TreeNode{
			Val:   nums[m],
			Left:  sortedArrayToBST(nums[:m]),
			Right: sortedArrayToBST(nums[m+1:]),
		}
	}

	log.Print(" ?= ", sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}
