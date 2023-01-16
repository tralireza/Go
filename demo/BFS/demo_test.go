package dBFS

import (
	"log"
	"testing"
)

func init() {
	log.Print("> BFS Demo")
}

func TestDemo(t *testing.T) {
	d := NewDemo(9, 56)
	d.AddBlock(64)
	d.AddDoor(8)
	d.Draw()
}

func TestBFS(t *testing.T) {
	d := NewDemo(9, 48)
	d.BFS(3, 5)
}
