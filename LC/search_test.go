package lc

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strings"
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
			grid[i][j] = '🏰'
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

	grid := [][]rune{{'🏠', '🏠', '🏠', '🏠', '🌊'}, {'🏠', '🏠', '🌊', '🏠', '🌊'}, {'🏠', '🏠', '🌊', '🌊', '🌊'}, {'🌊', '🌊', '🌊', '🌊', '🌊'}}
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
	draw()
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

// 463 Island Perimeter
func Test463(t *testing.T) {
	islandPerimeter := func(grid [][]int) int {
		Rows, Cols := len(grid), len(grid[0])
		Dirs := []int{0, 1, 0, -1, 0}

		perimeter := 0
		for r := 0; r < Rows; r++ {
			for c := 0; c < Cols; c++ {
				if grid[r][c] == 0 {
					continue
				}

				for k := range Dirs[:4] {
					p, q := r+Dirs[k], c+Dirs[k+1]
					if p < 0 || p > Rows-1 || q < 0 || q > Cols-1 || grid[p][q] == 0 {
						perimeter++
					}
				}
			}
		}
		return perimeter
	}

	log.Print("16 ?= ", islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	log.Print("4 ?= ", islandPerimeter([][]int{{0, 1, 0}}))
	log.Print("4 ?= ", islandPerimeter([][]int{{1}}))
}

// 1992m Find All Groups of Farmland
func Test1992(t *testing.T) {
	findFarmland := func(land [][]int) [][]int {
		Cords := [][]int{}
		Rows, Cols := len(land), len(land[0])
		Dx, Dy := []int{0, 0, 1, 1}, []int{1, -1, 0, 0}

		var m, n int
		var dfs func(i, j int)
		dfs = func(i, j int) {
			land[i][j] = -1
			m, n = max(m, i), max(n, j)

			for k := 0; k < 4; k++ {
				x, y := i+Dx[k], j+Dy[k]
				if x >= 0 && x < Rows && y >= 0 && y < Cols && land[x][y] == 1 {
					dfs(x, y)
				}
			}
		}

		for i := 0; i < Rows; i++ {
			for j := 0; j < Cols; j++ {
				if land[i][j] == 1 {
					m, n = i, j
					dfs(i, j)
					Cords = append(Cords, []int{i, j, m, n})
				}
			}
		}

		return Cords
	}

	log.Print(" ?= ", findFarmland([][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}))
	log.Print(" ?= ", findFarmland([][]int{{1}}))
	log.Print(" ?= ", findFarmland([][]int{{0}}))
}

// 419m Battleships in a Board
func Test419(t *testing.T) {
	countBattleships := func(board [][]byte) int {
		x := 0
		for r := 0; r < len(board); r++ {
			for c := 0; c < len(board[r]); c++ {
				if board[r][c] == 'X' &&
					(r == 0 || board[r-1][c] == '.') &&
					(c == 0 || board[r][c-1] == '.') {
					x++
				}
			}
		}
		return x
	}

	log.Print("2 ?= ", countBattleships([][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}))
}

// 140h Word Break II
func Test140(t *testing.T) {
	wordBreak := func(s string, wordDict []string) []string {
		m := map[string]bool{}
		for _, w := range wordDict {
			m[w] = true
		}

		K := make([][]byte, len(s))
		for i := range K {
			K[i] = make([]byte, len(s))
		}

		D := make([]bool, len(s)+1)
		D[0] = true

		for l := 1; l <= len(s); l++ {
			for start := 0; start < l; start++ {
				if D[start] && m[s[start:l]] {
					D[l] = true
					K[start][l-1] = 1
				}
			}
		}

		for i, v := range K {
			log.Printf("%2d %v", i, v)
		}

		W := []string{}

		var draw func(int, []string)
		draw = func(start int, ws []string) {
			if start == len(s) {
				W = append(W, strings.Join(ws, " "))
				return
			}

			for end, v := range K[start] {
				if v == 1 {
					draw(end+1, append(ws, s[start:end+1]))
				}
			}
		}
		draw(0, []string{})

		return W
	}

	log.Printf(" ?= %q", wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	log.Printf(" ?= %q", wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	log.Printf(" ?= %q", wordBreak("catsandog", []string{"cat", "cats", "and", "sand", "dog"}))
	log.Printf(" ?= %q", wordBreak("a", []string{"a"}))
}

// 514h Freedom Trail
func Test514(t *testing.T) {
	findRotateSteps := func(ring string, key string) int {
		R, K := len(ring), len(key)
		Mem := map[[2]int]int{}

		var search func(r, k int) int
		search = func(r, k int) int {
			if k == K {
				return 0
			}

			if v, ok := Mem[[2]int{r, k}]; ok {
				return v
			}

			minSteps := math.MaxInt
			for x := range R {
				if ring[x] == key[k] {
					steps := x - r
					if steps < 0 {
						steps *= -1
					}
					minSteps = min(minSteps, search(x, k+1)+min(steps, R-steps))
				}
			}

			Mem[[2]int{r, k}] = minSteps
			return minSteps
		}

		return K + search(0, 0)
	}

	bottomUp := func(ring, key string) int {
		R, K := len(ring), len(key)

		D := make([][]int, K+1)
		for k := range D {
			D[k] = make([]int, R)
		}

		for k := K - 1; k >= 0; k-- {
			for r := range R {
				D[k][r] = math.MaxInt

				for i := range R {
					if ring[i] == key[k] {

						steps := i - r
						if steps < 0 {
							steps *= -1
						}
						Si := min(steps, R-steps)

						D[k][r] = min(D[k][r], D[k+1][i]+Si)
					}
				}
			}

		}

		return K + D[0][0]
	}

	spaceOptimized := func(ring string, key string) int {
		R, K := len(ring), len(key)
		D, P := make([]int, R), make([]int, R)

		for k := K - 1; k >= 0; k-- {
			copy(P, D)

			for r := range R {
				D[r] = math.MaxInt

				for x := range R {
					if ring[x] == key[k] {
						cw := x - r
						if cw < 0 {
							cw *= -1
						}
						acw := R - cw

						Si := min(cw, acw)
						D[r] = min(D[r], Si+P[x])
					}
				}

			}
		}

		return K + D[0]
	}

	for _, f := range []func(string, string) int{findRotateSteps, bottomUp, spaceOptimized} {
		log.Print("4 ?= ", f("godding", "gd"))
		log.Print("13 ?= ", f("godding", "godding"))
		log.Print("14 ?= ", f("godding", "dogdog"))
		log.Print("6 ?= ", f("abcde", "ade"))
		log.Print("===")
	}
}
