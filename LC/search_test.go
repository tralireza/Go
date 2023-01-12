package lc

import (
	"fmt"
	"log"
	"slices"
	"testing"
)

func init() {
	log.Print("> BFS/DFS")
}

// 1091m Shortest Path in Binary Matrix
func Test1091(t *testing.T) {
	shortestPathBinaryMatrix := func(grid [][]int) int {
		type P struct{ i, j int }

		Q := []P{}
		if grid[0][0] == 0 {
			Q = append(Q, P{0, 0})
			grid[0][0] = 1
		}

		m, n := len(grid), len(grid[0])

		var v P
		for len(Q) > 0 {
			v, Q = Q[0], Q[1:]
			if v.i == m-1 && v.j == n-1 {
				return grid[m-1][n-1]
			}

			for _, d := range [8]P{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, -1}, {1, 1}, {-1, 1}, {-1, -1}} {
				u := P{v.i + d.i, v.j + d.j}

				if u.i >= 0 && u.j >= 0 && m > u.i && n > u.j && grid[u.i][u.j] == 0 {
					grid[u.i][u.j] = 1 + grid[v.i][v.j]

					Q = append(Q, u)
				}
			}
		}

		return -1
	}

	gridSet := func() [][]int { return [][]int{{0, 0, 0, 0, 0}, {1, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {0, 0, 1, 1, 0}} }
	grid := gridSet()

	draw := func() {
		for i := range grid {
			for _, v := range grid[i] {
				fmt.Printf("| %d ", v)
			}
			fmt.Println("|")
		}
	}
	draw()
	log.Print("5 ?= ", shortestPathBinaryMatrix(grid))
	draw()
}

// 200m Number of Islands
func Test200(t *testing.T) {
	numIslands := func(grid [][]rune) int {
		m, n := len(grid), len(grid[0])
		islands := 0

		dirs := []int{0, 1, 0, -1, 0}

		var dfs func(i, j int)
		dfs = func(i, j int) {
			grid[i][j] = 'X'
			for k := range dirs[:4] {
				p, q := i+dirs[k], j+dirs[k+1]
				if p >= 0 && m > p && q >= 0 && n > q && grid[p][q] == '🏠' {
					dfs(p, q)
				}
			}
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == '🏠' {
					islands++
					dfs(i, j)
				}
			}
		}

		return islands
	}

	grid := [][]rune{{'🏠', '🏠', '🏠', '🏠', '💧'}, {'🏠', '🏠', '💧', '🏠', '💧'}, {'🏠', '🏠', '💧', '💧', '💧'}, {'💧', '💧', '💧', '💧', '💧'}}
	draw := func() {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				fmt.Printf("| %c ", grid[i][j])
			}
			fmt.Println("|")
		}
	}
	draw()
	log.Print("1 ?= ", numIslands(grid))
}

// 130m Surrounded Regions
func Test130(t *testing.T) {
	solve := func(board [][]byte) {
		m, n := len(board), len(board[0])
		dirs := []int{0, 1, 0, -1, 0}

		var dfs func(i, j int)
		dfs = func(i, j int) {
			board[i][j] = '*'

			for k := range dirs[:4] {
				p, q := i+dirs[k], j+dirs[k+1]
				if p >= 0 && m > p && q >= 0 && n > q && board[p][q] == 'O' {
					dfs(p, q)
				}
			}
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {
					dfs(i, j)
				}
			}
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] == 'O' {
					board[i][j] = 'X'
				}
				if board[i][j] == '*' {
					board[i][j] = 'O'
				}
			}
		}
	}

	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	draw := func() {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				fmt.Printf("| %c ", board[i][j])
			}
			fmt.Printf("|\n")
		}
	}
	draw()
	solve(board)
	log.Print("===")
	draw()
}

// 133m Clone Graph
func Test133(t *testing.T) {
	type Node struct {
		Val       int
		Neighbors []*Node
	}

	cloneGraph := func(node *Node) *Node {
		m := map[*Node]*Node{}

		var rclone func(*Node) *Node
		rclone = func(n *Node) *Node {
			if c, ok := m[n]; ok {
				return c
			}

			c := &Node{Val: n.Val}
			m[n] = c
			for _, v := range n.Neighbors {
				c.Neighbors = append(c.Neighbors, rclone(v))
			}
			return c
		}

		return rclone(node)
	}

	type N = Node
	g := &N{1, []*N{{2, []*N{{Val: 4}}}}}
	n := &N{3, []*N{g}}
	g.Neighbors = append(g.Neighbors, n)

	c := cloneGraph(g)
	log.Print(g, g.Neighbors[0], g.Neighbors[1])
	log.Print(c, c.Neighbors[0], c.Neighbors[1])
}

// 2300m Successful Pairs of Spells and Potions
func Test2300(t *testing.T) {
	successfulPairs := func(spells []int, potions []int, success int64) []int {
		slices.Sort(potions)

		pairs := make([]int, 0, len(spells))
		for _, spell := range spells {
			l, r := 0, len(potions)
			for l < r {
				m := l + (r-l)>>1
				if int64(spell)*int64(potions[m]) >= success {
					r = m
				} else {
					l = m + 1
				}
			}
			pairs = append(pairs, len(potions)-l)
		}

		return pairs
	}

	log.Print("[4 0 3] ?= ", successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
	log.Print("[2 0 2] ?= ", successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
	log.Print("[1 4 3] ?= ", successfulPairs([]int{1, 4, 2}, []int{2, 5, 1, 3}, 4))
}
