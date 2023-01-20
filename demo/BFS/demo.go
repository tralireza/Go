package dBFS

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	log.SetFlags(0)
}

type Point struct{ row, col int }
type Demo struct {
	M, N int
	Grid map[Point]rune  // visual
	P    map[Point]Point // parent/predecessor
	D    map[Point]int   // distance to source/start
}

const (
	Space = '🪡'
	Wall  = '🧱'

	Start   = '👻' // White
	Looking = '👀' // Gray
	Done    = '🐾' // Black
	Success = '👍' // Black

	Up    = '👆'
	Down  = '👇'
	Left  = '👈'
	Right = '👉'
)

func NewDemo(m, n int) *Demo {
	d := &Demo{M: m, N: n, P: map[Point]Point{}, D: map[Point]int{}}

	g := map[Point]rune{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := Space
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				v = Wall
			}
			g[Point{i, j}] = v
		}
	}
	d.Grid = g

	return d
}

func (o *Demo) AddBlock(k int) {
	for k > 0 {
		i, j := rand.Intn(o.M-1)+1, rand.Intn(o.N-1)+1
		if o.Grid[Point{i, j}] == Space {
			o.Grid[Point{i, j}] = Wall
			k--
		}
	}
}

func (o *Demo) AddDoor(k int) {
	for k > 0 {
		var i, j int
		switch rand.Intn(2) {
		case 0:
			i = rand.Intn(2) * (o.M - 1)
			j = rand.Intn(o.N-1) + 1
		default:
			i = rand.Intn(o.M-1) + 1
			j = rand.Intn(2) * (o.N - 1)
		}
		if o.Grid[Point{i, j}] == Wall {
			o.Grid[Point{i, j}] = Space
			k--
		}
	}
}

func (o *Demo) Draw() {
	for i := range o.M {
		fmt.Printf("\033[%d;%dH", i+1, 1)
		for j := range o.N {
			fmt.Printf("%c", o.Grid[Point{i, j}])
		}
	}
}

func (o *Demo) adjacents(p Point) []Point {
	P := []Point{}
	dirs := []int{0, 1, 0, -1, 0}
	for i := range dirs[:4] {
		q := Point{p.row + dirs[i], p.col + dirs[i+1]}
		if q.row >= 0 && o.M > q.row && q.col >= 0 && o.N > q.col && o.Grid[q] != Wall {
			P = append(P, q)
		}
	}
	return P
}

func (o *Demo) success(p Point) bool {
	if p.row == 0 || p.row == o.M-1 || p.col == 0 || p.col == o.N-1 {
		for o.Grid[p] != Start {
			prv := o.P[p]
			if o.Grid[p] != Success {
				switch {
				case prv.row < p.row:
					o.Grid[p] = Up
				case prv.row > p.row:
					o.Grid[p] = Down
				case prv.col < p.col:
					o.Grid[p] = Left
				case prv.col > p.col:
					o.Grid[p] = Right
				}
			}
			p = prv
		}
		return true
	}
	return false
}

func (o *Demo) DFS(s Point) {
	o.Search(s, func(Q *[]Point) Point {
		u := (*Q)[len(*Q)-1]
		*Q = (*Q)[:len(*Q)-1]
		return u
	})
}

func (o *Demo) BFS(s Point) {
	o.Search(s, func(Q *[]Point) Point {
		u := (*Q)[0]
		*Q = (*Q)[1:]
		return u
	})
}

func (o *Demo) Search(s Point, dQueue func(Q *[]Point) Point) {
	fmt.Print("\033[2J")   // clear screen
	fmt.Print("\x1b[?25l") // low(hide) cursor

	o.Grid[s] = Start
	o.D[s] = 0
	o.Draw()

	Q := []Point{s}

	for len(Q) > 0 {
		u := dQueue(&Q)

		for _, v := range o.adjacents(u) {
			if o.Grid[v] == Space {
				o.Grid[v] = Looking
				o.D[v], o.P[v] = 1+o.D[u], u

				Q = append(Q, v)
			}
		}

		o.Draw()
		time.Sleep(75 * time.Millisecond)

		if o.success(u) {
			o.Grid[u] = Success
		} else if o.Grid[u] != Start {
			o.Grid[u] = Done
		}
	}

	o.Draw()

	fmt.Print("\x1b[?25h") // high(show) cursor
	fmt.Print("\n")
}
