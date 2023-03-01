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

func GDirected(vertices byte, edges [][2]byte) *MyGraph   { return New(vertices, edges, true) }
func GUndirected(vertices byte, edges [][2]byte) *MyGraph { return New(vertices, edges, false) }
func New(vertices byte, edges [][2]byte, isDirected bool) *MyGraph {
	g := &MyGraph{isDirected: isDirected}

	for v := byte(0); v < vertices; v++ {
		g.gV = append(g.gV, &Vertex{V: v, Label: v})
	}

	M, L := make([][]byte, len(g.gV)), make([][]byte, len(g.gV))
	for i := range M {
		M[i] = make([]byte, len(g.gV))
	}
	for _, edge := range edges {
		v, u := edge[0], edge[1]
		M[v][u] = 1
		if !isDirected {
			M[u][v] = 1
		}
		L[v] = append(L[v], u)
		if !isDirected {
			L[u] = append(L[u], v)
		}
	}
	g.adjMatrix, g.adjList = M, L

	return g
}

// Small Un/Directed Graph of [0..255] Vertices
type MyGraph struct {
	gV         []*Vertex
	gE         [][2]*Vertex
	adjMatrix  [][]byte
	adjList    [][]byte
	isDirected bool
}

func (g *MyGraph) Vertices() []*Vertex { return g.gV }
func (g *MyGraph) Edges() [][2]*Vertex { return g.gE }

func (g MyGraph) String() string {
	var sb strings.Builder
	sb.WriteString("+ G(V,E) adjacency-list:\n")
	for _, v := range g.gV {
		sb.WriteString(fmt.Sprintf(" (%v) -> [", v.Label))
		lbls := []string{}
		for _, u := range g.adjList[v.V] {
			lbls = append(lbls, fmt.Sprintf("(%v)", g.gV[u].Label))
		}
		sb.WriteString(strings.Join(lbls, " "))
		sb.WriteString("]\n")
	}
	return sb.String()
}

func (g *MyGraph) SetVertex(i int, v Vertex) { g.gV[i].Label = v.Label }
