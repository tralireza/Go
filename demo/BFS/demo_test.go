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

func TestGrid(t *testing.T) {
	fmt.Printf("\033[2J") // cls: clear screen

	d := NewDemo(10, 56)
	d.AddBlock(128)
	d.AddDoor(16)
	d.Draw()

	fmt.Printf("\x1b[?25l") // low(hide) cursor
	time.Sleep(time.Second)
	fmt.Printf("\x1b[?25h") // high(show) cursor
	fmt.Printf("\n")
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

func TestLargeGrid(t *testing.T) {
	d := NewDemo(40, 56)
	d.AddBlock(512)
	d.AddDoor(16)
	d.DFS(Point{5, 28})
}
