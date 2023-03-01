package Graph

import (
	"fmt"
	"log"
	"strings"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("")
}

type Vertex struct {
	V     byte
	Label any
}

type Graph interface {
	Vertices() []*Vertex
	Edges() [][2]*Vertex
}

func New(vertices byte, edges [][2]byte) *MyGraph {
	g := &MyGraph{}
	for v := byte(0); v < vertices; v++ {
		g.gV = append(g.gV, &Vertex{V: v, Label: v})
	}

	M, L := make([][]byte, len(g.gV)), make([][]byte, len(g.gV))
	for i := range M {
		M[i] = make([]byte, len(g.gV))
	}
	for _, edge := range edges {
		v, u := edge[0], edge[1]
		M[v][u], M[u][v] = 1, 1
		L[v], L[u] = append(L[v], u), append(L[u], v)
	}
	g.adjMatrix, g.adjList = M, L

	return g
}

// Small Undirected Graph of [0..255] Vertices
type MyGraph struct {
	gV        []*Vertex
	gE        [][2]*Vertex
	adjMatrix [][]byte
	adjList   [][]byte
}

func (g *MyGraph) Vertices() []*Vertex { return g.gV }
func (g *MyGraph) Edges() [][2]*Vertex { return g.gE }

func (g MyGraph) String() string {
	var sb strings.Builder
	for _, v := range g.gV {
		sb.WriteString(fmt.Sprintf("(%v) -> %v\n", v.Label, g.adjList[v.V]))
	}
	return sb.String()
}

func (g *MyGraph) SetVertex(i int, v Vertex) { g.gV[i].Label = v.Label }
