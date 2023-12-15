package lc

import (
	"log"
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
			return 0, false
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

	djs := DJS{}

	makeSet(djs, 1)
	makeSet(djs, 2)
	makeSet(djs, 3)
	makeSet(djs, 1024)
	log.Printf("%v", djs)

	log.Print(findSet(djs, 2))
	log.Print(findSet(djs, 7))

	union(djs, 3, 1024)
	log.Print(findSet(djs, 1024))
	makeSet(djs, 7)
	makeSet(djs, 9)
	union(djs, 3, 7)
	union(djs, 7, 9)
	log.Print(djs)
}
