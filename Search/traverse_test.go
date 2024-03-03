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

func aBST() *TreeNode {
	var r *TreeNode
	for _, v := range []int{5, 3, 8, 1, 4, 6, 9, 0, 7, 2} {
		r = InsertBST(r, v)
	}
	return r
}

func DrawBFS(t *TreeNode) {
	Q := []*TreeNode{}
	Q = append(Q, t)

	for l := 0; len(Q) > 0; l++ {
		log.Printf("%d -> %v", l, Q)
		for lsize := len(Q); lsize > 0; lsize-- {
			n := Q[0]
			Q = Q[1:]
			if n.Left != nil {
				Q = append(Q, n.Left)
			}
			if n.Right != nil {
				Q = append(Q, n.Right)
			}
		}
	}
}

func TestDeleteBST(t *testing.T) {
	r := aBST()
	DrawBFS(r)
	for _, d := range []int{8, 1, 7, 0} {
		log.Printf("> %d X", d)
		DeleteBST(r, d)
		DrawBFS(r)
	}
}

func TestInsertBST(t *testing.T) {
	r := aBST()

	Q := []*TreeNode{}
	Q = append(Q, r)
	l := 0
	for len(Q) > 0 {
		log.Printf("%d -> %v", l, Q)
		for i := len(Q); i > 0; i-- {
			n := Q[0]
			Q = Q[1:]
			if n.Left != nil {
				Q = append(Q, n.Left)
			}
			if n.Right != nil {
				Q = append(Q, n.Right)
			}
		}
		l++
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

func Test399(t *testing.T) {
	log.Print(CalEquation([][]string{{"a", "b"}, {"b", "c"}, {"e", "f"}, {"g", "f"}, {"c", "d"}},
		[]float64{2, 3, 4, 8, .25},
		[][]string{{"a", "c"}, {"f", "e"}, {"e", "g"}, {"x", "x"}, {"a", "f"}, {"b", "g"}, {"a", "b"}}))
}

func Test1926(t *testing.T) {
	maze := [][]byte{
		{'+', '.', '+', '+', '+', '+', '+'},
		{'+', '.', '+', '.', '.', '.', '+'},
		{'+', '.', '+', '.', '+', '.', '+'},
		{'+', '.', '.', '.', '.', '.', '+'},
		{'+', '+', '+', '+', '.', '+', '.'}}

	if v := NearestExit(maze, []int{0, 1}); v != 7 {
		t.Fatalf("Wrong number of steps: 7 != %d", v)
	}

	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("|%c", maze[i][j])
		}
		fmt.Println("|")
	}
}

func Test994(t *testing.T) {
	grid := [][]int{
		{2, 1, 1, 0, 1, 0},
		{1, 1, 0, 1, 0, 1},
		{1, 2, 0, 1, 1, 1},
		{0, 1, 1, 0, 2, 0}}

	if v := OrangesRotting(grid); v != -1 {
		t.Fatalf("Wrong time: -1 != %d", v)
	}

	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("|% d", grid[i][j])
		}
		fmt.Println("|")
	}
}
