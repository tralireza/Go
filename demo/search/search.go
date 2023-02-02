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
	Grid             [][]rune        // visual
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

	g, c := make([][]rune, m), make([][]byte, m)
	for i := 0; i < m; i++ {
		g[i], c[i] = make([]rune, n), make([]byte, n)
		for j := 0; j < n; j++ {
			g[i][j] = Space
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				g[i][j] = Wall
			} else {
				c[i][j] = 'W'
			}
		}
	}
	d.Grid = g
	d.Color = c

	return d
}

func (o *Demo) SetStart(p Point) Point {
	if p.Row <= 0 || p.Row >= o.M-1 {
		p.Row = rand.Intn(o.M-2) + 1
	}
	if p.Col <= 0 || p.Col >= o.N-1 {
		p.Col = rand.Intn(o.N-2) + 1
	}
	o.Grid[p.Row][p.Col] = Start
	return p
}

func (o *Demo) AddBlock(k int) {
	if k > (o.M-2)*(o.N-2) {
		k = (o.M - 2) * (o.N - 2)
	}

	for k > 0 {
		i, j := rand.Intn(o.M-1)+1, rand.Intn(o.N-1)+1
		if o.Grid[i][j] == Space {
			o.Grid[i][j] = Wall
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
			i, j = rand.Intn(2)*(o.M-1), rand.Intn(o.N)
		default:
			i, j = rand.Intn(o.M), rand.Intn(2)*(o.N-1)
		}
		if o.Grid[i][j] == Wall {
			o.Grid[i][j] = Space
			o.Color[i][j] = 'W'
			k--
		}
	}
}

func (o *Demo) Draw() {
	for i := range o.M {
		fmt.Printf("\x1b[%d;%dH", i+1, 1)
		for j := range o.N {
			fmt.Printf("%c", o.Grid[i][j])
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
		if q.Row >= 0 && o.M > q.Row && q.Col >= 0 && o.N > q.Col && o.Grid[q.Row][q.Col] != Wall {
			P = append(P, q)
		}
	}
	return P
}

func (o *Demo) Breadcrumb(exit Point, taste int) {
	i, j := exit.Row, exit.Col
	for o.Grid[i][j] != Start {
		prv := o.P[Point{i, j}]
		if o.Grid[i][j] != Success {
			switch taste { // Breadcrumb Taste
			case 0:
				o.Grid[i][j] = Bee
			case 1, 2: // 2: keep Beeline
				if taste == 1 || o.Grid[i][j] != Bee {
					var r rune
					switch {
					case prv.Row < i:
						r = Up
					case prv.Row > i:
						r = Down
					case prv.Col < j:
						r = Left
					case prv.Col > j:
						r = Right
					}
					o.Grid[i][j] = r
				}
			}
		}
		i, j = prv.Row, prv.Col
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
		o.Grid[p.Row][p.Col] = Success

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
	o.Grid[o.start.Row][o.start.Col] = Start
	o.D[s] = 0

	o.Draw()

	Q := []Point{o.start}
	o.Color[o.start.Row][o.start.Col] = 'G' // Gray: Visiting
	for len(Q) > 0 {
		u := dQueue(&Q)

		for _, v := range o.adjacents(u) {
			if o.Color[v.Row][v.Col] == 'W' { // White: Not visited
				o.Color[v.Row][v.Col] = 'G' // Gray: Visiting
				o.Grid[v.Row][v.Col] = Looking
				o.D[v], o.P[v] = 1+o.D[u], u

				Q = append(Q, v)
			}
		}

		o.Color[u.Row][u.Col] = 'B' // Black: Visited
		if !o.isDoor(u) && o.Grid[u.Row][u.Col] != Start {
			o.Grid[u.Row][u.Col] = Done
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
