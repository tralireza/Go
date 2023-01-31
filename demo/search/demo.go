package search

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func init() {
	log.SetFlags(0)
}

type Point struct{ Row, Col int }
type Demo struct {
	M, N             int
	Grid             map[Point]rune  // visual
	P                map[Point]Point // parent/predecessor
	D                map[Point]int   // distance to source/start
	start            Point           // source/start
	fdoors, shortest int             // doors found, shortest path
	exit             Point           // shortest distance exit
	steps            uint            // steps
	Color            [][]byte        // 'W'hite, 'G'ray, 'B'lack -> Not visited, Visiting, Visited
}

const (
	Space = rune(0x3000) // CJK width space
	Wall  = '🧱'

	Start   = '👻' // node color: White
	Looking = '👀' // node color: Gray
	Done    = '🥽' // node color: Black
	Success = '👍' // node color: Black
	Bee     = '🐝' // node color: Black (shortest distance)

	Up    = '👆'
	Down  = '👇'
	Left  = '👈'
	Right = '👉'
)

func NewDemo(m, n int) *Demo {
	if m < 3 {
		m = 3
	}
	if n < 3 {
		n = 3
	}

	d := &Demo{M: m, N: n, P: map[Point]Point{}, D: map[Point]int{}, shortest: math.MaxInt}

	g, color := map[Point]rune{}, make([][]byte, m)
	for i := 0; i < m; i++ {
		color[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			v := Space
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				v = Wall
			}
			g[Point{i, j}] = v
			if v == Space {
				color[i][j] = 'W'
			}
		}
	}
	d.Grid = g
	d.Color = color

	return d
}

func (o *Demo) SetStart(p Point) Point {
	if p.Row <= 0 || p.Row >= o.M-1 {
		p.Row = rand.Intn(o.M-2) + 1
	}
	if p.Col <= 0 || p.Col >= o.N-1 {
		p.Col = rand.Intn(o.N-2) + 1
	}
	o.Grid[p] = Start
	return p
}

func (o *Demo) AddBlock(k int) {
	if k > (o.M-2)*(o.N-2) {
		k = (o.M - 2) * (o.N - 2)
	}

	for k > 0 {
		i, j := rand.Intn(o.M-1)+1, rand.Intn(o.N-1)+1
		if o.Grid[Point{i, j}] == Space {
			o.Grid[Point{i, j}] = Wall
			k--
		}
	}
}

func (o *Demo) AddDoor(k int) {
	if k > 2*(o.M+o.N-2) {
		k = 2 * (o.M + o.N - 2)
	}

	for k > 0 {
		var i, j int
		switch rand.Intn(2) {
		case 0:
			i = rand.Intn(2) * (o.M - 1)
			j = rand.Intn(o.N)
		default:
			i = rand.Intn(o.M)
			j = rand.Intn(2) * (o.N - 1)
		}
		if o.Grid[Point{i, j}] == Wall {
			o.Grid[Point{i, j}] = Space
			o.Color[i][j] = 'W'
			k--
		}
	}
}

func (o *Demo) Draw() {
	for i := range o.M {
		fmt.Printf("\x1b[%d;%dH", i+1, 1)
		for j := range o.N {
			fmt.Printf("%c", o.Grid[Point{i, j}])
		}
	}
}

func (o *Demo) Stat(t int) {
	fmt.Printf("\x1b[%d;%dH", o.M+1, 1) // move cursor/position
	if t == 0 {
		fmt.Printf("[ 💅 ]")
	} else {
		fmt.Printf("[ %c ]", []rune{'💿', '📀'}[o.steps%2])
		o.steps++
	}

	fmt.Printf("     %4d %c   %4d %c   ", t, Looking, o.fdoors, Success)
	if o.shortest < math.MaxInt {
		fmt.Printf("%4d %c", o.shortest, Bee)
	} else {
		fmt.Printf("   ∞ %c", Bee)
	}
}

func (o *Demo) adjacents(p Point) []Point {
	P := []Point{}
	dirs := []int{0, 1, 0, -1, 0}
	for i := range dirs[:4] {
		q := Point{p.Row + dirs[i], p.Col + dirs[i+1]}
		if q.Row >= 0 && o.M > q.Row && q.Col >= 0 && o.N > q.Col && o.Grid[q] != Wall {
			P = append(P, q)
		}
	}
	return P
}

func (o *Demo) Breadcrumb(exit Point, m int) {
	p := exit
	for o.Grid[p] != Start {
		prv := o.P[p]
		if o.Grid[p] != Success {
			switch m {
			case 0:
				o.Grid[p] = Bee
			case 1, 2: // 2: keep Beeline
				if m == 1 || o.Grid[p] != Bee {
					switch {
					case prv.Row < p.Row:
						o.Grid[p] = Up
					case prv.Row > p.Row:
						o.Grid[p] = Down
					case prv.Col < p.Col:
						o.Grid[p] = Left
					case prv.Col > p.Col:
						o.Grid[p] = Right
					}
				}
			}
		}
		p = prv
	}
}

func (o *Demo) isDoor(p Point) bool {
	if p.Row == 0 || p.Row == o.M-1 || p.Col == 0 || p.Col == o.N-1 {
		o.fdoors++

		if o.D[p] < o.shortest {
			if o.shortest < math.MaxInt {
				o.Breadcrumb(o.exit, 1)
			}
			o.shortest = o.D[p]
			o.exit = p
			o.Breadcrumb(p, 0)
		} else {
			o.Breadcrumb(p, 2)
		}
		o.Grid[p] = Success

		return true
	}
	return false
}

func (o *Demo) DFS(s Point) {
	o.search(s, func(Q *[]Point) Point {
		u := (*Q)[len(*Q)-1]
		*Q = (*Q)[:len(*Q)-1]
		return u
	})
}

func (o *Demo) BFS(s Point) {
	o.search(s, func(Q *[]Point) Point {
		u := (*Q)[0]
		*Q = (*Q)[1:]
		return u
	})
}

func (o *Demo) search(s Point, dQueue func(Q *[]Point) Point) {
	fmt.Print("\x1b[2J")   // clear screen
	fmt.Print("\x1b[?25l") // low(hide) cursor

	o.start = o.SetStart(s)
	o.Grid[o.start] = Start
	o.D[s] = 0

	o.Draw()

	Q := []Point{o.start}
	o.Color[o.start.Row][o.start.Col] = 'G' // Gray: Visiting
	for len(Q) > 0 {
		u := dQueue(&Q)

		for _, v := range o.adjacents(u) {
			if o.Color[v.Row][v.Col] == 'W' { // White: Not visited
				o.Color[v.Row][v.Col] = 'G' // Gray: Visiting
				o.Grid[v] = Looking
				o.D[v], o.P[v] = 1+o.D[u], u

				Q = append(Q, v)
			}
		}

		o.Color[u.Row][u.Col] = 'B' // Black: Visited
		if !o.isDoor(u) && o.Grid[u] != Start {
			o.Grid[u] = Done
		}

		o.Draw()
		o.Stat(len(Q))
		time.Sleep(75 * time.Millisecond)
	}

	fmt.Print("\x1b[2J")
	o.Draw()
	o.Stat(0)
	fmt.Print("\x1b[?25h") // high(show) cursor
	fmt.Print("\n")
}
