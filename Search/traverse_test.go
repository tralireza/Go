package tgraph

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(0)
}

type AGraph func() map[int][]int

func (f AGraph) Nodes() []int {
	ns := []int{}
	for n := range f() {
		ns = append(ns, n)
	}
	return ns
}
func (f AGraph) Edges(n int) []int { return f()[n] }

var tGraph AGraph = (func() map[int][]int {
	return map[int][]int{
		7: {1, 3, 9, 2},
		1: {4, 1, 3},
		2: {6, 5},
		3: {7, 4},
		4: {5, 9},
		5: {2, 3},
		6: {9, 3},
		9: {7, 1},
		8: {0},
		0: {8},
	}
})

func TestBFS(t *testing.T) {
	Visited := map[int]bool{}

	BFS := func(u int) {
		log.Printf(">>> BFS: %d", u)

		Q := []int{}
		Q = append(Q, u)
		for len(Q) > 0 {
			u := Q[0]
			Q = Q[1:]

			log.Printf("%d -> [V]", u)

			Visited[u] = true
			for _, v := range tGraph.Edges(u) {
				if !Visited[v] {
					Q = append(Q, v)
				}
			}
		}
	}

	for _, v := range tGraph.Nodes() {
		if !Visited[v] {
			BFS(v)
		}
	}
}

func TestDFS(t *testing.T) {
	Visited := map[int]bool{}

	var DFS func(int)
	DFS = func(u int) {
		log.Printf("%d -> [V]", u)

		Visited[u] = true
		for _, v := range tGraph.Edges(u) {
			if !Visited[v] {
				DFS(v)
			}
		}
	}

	for _, u := range tGraph.Nodes() {
		if !Visited[u] {
			log.Printf(">>> DFS: %d", u)
			DFS(u)
		}
	}
}

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

func toOddEvenList(head *ListNode) *ListNode {
	heven := head.Next

	odd, even := head, head.Next
	for odd.Next != nil && odd.Next.Next != nil {
		odd.Next, even.Next = odd.Next.Next, odd.Next.Next.Next
		even, odd = even.Next, odd.Next
	}

	odd.Next = heven
	return head
}

func Test328(t *testing.T) {
	h := toOddEvenList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, nil}}}}}}}})
	for n := h; n != nil; n = n.Next {
		if n.Next != nil {
			fmt.Print(n, "-> ")
		} else {
			fmt.Println(n)
		}
	}
}

func isEvenOddTree(root *TreeNode) bool {
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

func aTree() *TreeNode {
	return &TreeNode{1, &TreeNode{4, &TreeNode{Val: 3}, &TreeNode{Val: 5}}, &TreeNode{2, &TreeNode{Val: 7}, nil}}
}

func Test1609(t *testing.T) {
	log.Print(isEvenOddTree(aTree()))
}

func TestDFS1448(t *testing.T) {
	log.Print(DFS1448(aTree()))
}
