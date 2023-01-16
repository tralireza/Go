package dBFS

import (
	"fmt"
	"log"
	"math/rand"
)

func init() {
	log.SetFlags(0)
}

type Point struct{ row, col int }
type Demo struct {
	M, N int
	Grid map[Point]rune
	P    map[Point]Point
	D    map[Point]int
}

const (
	SPC  = '🟰'
	WALL = '🧱'
)

func NewDemo(m, n int) *Demo {
	d := &Demo{M: m, N: n, P: map[Point]Point{}, D: map[Point]int{}}

	g := map[Point]rune{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := SPC
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				v = WALL
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
		if o.Grid[Point{i, j}] == SPC {
			o.Grid[Point{i, j}] = WALL
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
		if o.Grid[Point{i, j}] == WALL {
			o.Grid[Point{i, j}] = SPC
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

func (o *Demo) BFS(i, j int) {
	//
}
