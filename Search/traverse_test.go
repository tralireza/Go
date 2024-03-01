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

/*
				     1
				4         2
			3   5     7   9
		 8     6       4
	    11 13         15
*/
var root *TreeNode = &TreeNode{1,
	&TreeNode{4,
		&TreeNode{3, &TreeNode{8, nil, &TreeNode{Val: 11}}, nil},
		&TreeNode{5, nil, &TreeNode{6, &TreeNode{Val: 13}, nil}}},
	&TreeNode{2,
		&TreeNode{Val: 7},
		&TreeNode{9, &TreeNode{4, nil, &TreeNode{Val: 15}}, nil},
	},
}

func Test1161(t *testing.T) {
	log.Print(MaxLevelSum(root))
}

func Test1372(t *testing.T) {
	log.Print(LongestZigZag(root))
}

func Test437(t *testing.T) {
	log.Print(PathSum3(root, 9))
}

func Test1609(t *testing.T) {
	log.Print(IsEvenOddTree(root))
}

func TestDFS1448(t *testing.T) {
	log.Print(DFS1448(root))
}
