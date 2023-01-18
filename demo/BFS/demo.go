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

type Point struct{ r, c int } // r: row, c: column
type Demo struct {
	M, N int
	Grid map[Point]rune
	P    map[Point]Point
	D    map[Point]int
}

const (
	Space = '🟰'
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
		for j := range o.N {
			fmt.Printf("%c", o.Grid[Point{i, j}])
		}
		fmt.Printf("\n")
	}
}

func (o *Demo) adjacents(p Point) []Point {
	ps := []Point{}
	dirs := []int{0, 1, 0, -1, 0}
	for i := range dirs[:4] {
		q := Point{p.r + dirs[i], p.c + dirs[i+1]}
		if q.r >= 0 && o.M > q.r && q.c >= 0 && o.N > q.c && o.Grid[q] != Wall {
			ps = append(ps, q)
		}
	}
	return ps
}

func (o *Demo) success(p Point) bool {
	if p.r == 0 || p.r == o.M-1 || p.c == 0 || p.c == o.N-1 {
		for o.Grid[p] != Start {
			prv := o.P[p]
			if o.Grid[p] != Success {
				switch {
				case prv.r < p.r:
					o.Grid[p] = Up
				case prv.r > p.r:
					o.Grid[p] = Down
				case prv.c < p.c:
					o.Grid[p] = Left
				default:
					o.Grid[p] = Right
				}
			}
			p = prv
		}
		return true
	}
	return false
}

func (o *Demo) BFS(i, j int) {
	s := Point{i, j}
	o.Grid[s] = Start
	o.D[s] = 0
	o.Draw()

	Q := []Point{s}

	for len(Q) > 0 {
		u := Q[0]
		Q = Q[1:]
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
}
