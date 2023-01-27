package search

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func init() {
	log.Print("> BFS Demo")
}

func TestBreadcrumb(t *testing.T) {
	d := NewDemo(8, 16)
	d.SetStart(Point{3, 8})

	b, e := Point{3, 15}, Point{7, 13}
	d.Grid[b], d.Grid[e] = Success, Success
	for j := 15; j > 8; j-- {
		d.P[Point{3, j}] = Point{3, j - 1}
	}
	for i := 7; i > 3; i-- {
		d.P[Point{i, 13}] = Point{i - 1, 13}
	}

	d.Breadcrumb(b, 0)
	d.Draw()
	time.Sleep(time.Second)
	d.Breadcrumb(e, 2) // must not clear Beeline
	d.Draw()
	time.Sleep(time.Second)
	d.Breadcrumb(b, 1)
	d.Draw()

	fmt.Print("\n")
}

func TestValidGrid(t *testing.T) {
	d := NewDemo(3, -3)
	d.Draw()

	d.AddBlock(12)
	d.Draw()

	d.AddDoor(16)
	d.Draw()

	p := Point{7, 8}
	d.SetStart(p)
	d.Draw()

	d.BFS(p)
}

func TestGrid(t *testing.T) {
	fmt.Printf("\033[2J") // cls: clear screen

	d := NewDemo(10, 56)
	d.AddBlock(128)
	d.AddDoor(16)

	fmt.Printf("\x1b[?25l") // low(hide) cursor
	d.Draw()
	d.Stat(3)
	time.Sleep(425 * time.Millisecond)
	d.Stat(3)
	time.Sleep(425 * time.Millisecond)

	d.Stat(0)
	fmt.Printf("\n")
	fmt.Printf("\x1b[?25h") // high(show) cursor
}

func TestBFS(t *testing.T) {
	d := NewDemo(10, 56)
	d.AddBlock(128)
	d.AddDoor(16)

	d.BFS(Point{5, 28})
}

func TestDFS(t *testing.T) {
	d := NewDemo(10, 56)
	d.AddBlock(128)
	d.AddDoor(16)

	d.DFS(Point{5, 28})
}

func TestShort(t *testing.T) {
	d := NewDemo(32, 56)
	d.AddBlock(512)
	d.AddDoor(16)
	d.BFS(Point{16, 28})
}

func TestQuick(t *testing.T) {
	d := NewDemo(32, 56)
	d.AddBlock(512)
	d.AddDoor(16)
	d.DFS(Point{16, 28})
}
