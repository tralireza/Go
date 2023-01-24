package dBFS

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func init() {
	log.Print("> BFS Demo")
}

func TestValidGrid(t *testing.T) {
	d := NewDemo(3, -3)
	d.Draw()

	d.AddBlock(12)
	d.Draw()

	d.AddDoor(16)
	d.Draw()

	d.SetStart(Point{7, 8})
	d.Draw()
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
