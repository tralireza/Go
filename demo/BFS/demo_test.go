package dBFS

import (
	"log"
	"testing"
)

func init() {
	log.Print("> BFS Demo")
}

func TestDemo(t *testing.T) {
	d := NewDemo(10, 56)
	d.AddBlock(64)
	d.AddDoor(8)
	d.Draw()
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
