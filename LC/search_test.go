package lc

import (
	"fmt"
	"log"
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
	numIslands := func(grid [][]byte) int {
		m, n := len(grid), len(grid[0])
		islands := 0

		dirs := []int{0, 1, 0, -1, 0}

		var dfs func(i, j int)
		dfs = func(i, j int) {
			grid[i][j] = 'X'
			for k := range dirs[:4] {
				p, q := i+dirs[k], j+dirs[k+1]
				if p >= 0 && m > p && q >= 0 && n > q && grid[p][q] == '1' {
					dfs(p, q)
				}
			}
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == '1' {
					islands++
					dfs(i, j)
				}
			}
		}

		return islands
	}

	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
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
