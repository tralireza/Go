package lc

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func init() {
	log.Print("> DJS: Disjoint Set")
}

func TestDJS(t *testing.T) {
	type entry struct{ parent, rank int }
	type DJS map[int]entry

	makeSet := func(djs DJS, n int) {
		djs[n] = entry{n, 0}
	}

	var findSet func(DJS, int) (int, bool)
	findSet = func(djs DJS, n int) (int, bool) {
		if e, ok := djs[n]; ok {
			if n != e.parent {
				e.parent, _ = findSet(djs, e.parent)
				djs[n] = e
			}
			return e.parent, true
		} else {
			return n, false
		}
	}

	union := func(djs DJS, x, y int) {
		var ok bool
		x, ok = findSet(djs, x)
		if !ok {
			return
		}
		y, ok = findSet(djs, y)
		if !ok {
			return
		}

		X, Y := djs[x], djs[y]
		if X.rank > Y.rank {
			Y.parent = x
			djs[y] = Y
		} else {
			X.parent = y
			if Y.rank == X.rank {
				Y.rank++
			}
			djs[y], djs[x] = Y, X
		}
	}

	countSets := func(djs DJS) int {
		count := 0
		for x, X := range djs {
			if x == X.parent {
				count++
			}
		}
		return count
	}

	djs := DJS{}

	m := map[int]struct{}{}
	for range 25 {
		v := rand.Intn(100)
		m[v] = struct{}{}
		makeSet(djs, v)
	}
	log.Printf("%d:%d -> %v", countSets(djs), len(m), djs)

	log.Print(findSet(djs, 2))
	log.Print(findSet(djs, 7))

	var vs []int
	for v := range m {
		vs = append(vs, v)
	}
	for range 25 {
		union(djs, vs[rand.Intn(len(vs))], vs[rand.Intn(len(vs))])
	}

	log.Printf("%d -> %v", countSets(djs), djs)

	fmt.Print("+ Leaders: ")
	for x, X := range djs {
		if x == X.parent {
			fmt.Printf("%d:%v ", x, X)
		}
	}
	fmt.Println()
}
