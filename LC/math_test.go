package lc

import (
	"log"
	"slices"
	"testing"
)

func init() {
	log.Print("> Math")
}

// 3102h Minimize Manhattan Distances
func Test3102(t *testing.T) {
	minimumDistance := func(points [][]int) int {
		for i, p := range points {
			points[i] = []int{p[0] - p[1], p[0] + p[1]}
		}
		log.Print(points)

		f := func(i int) func(a, b []int) int {
			return func(a, b []int) int {
				return a[1] - b[1]
			}
		}

		mdist := 2*(100_000_000-1) + 1
		for i := range points {
			ps := make([][]int, 0, len(points)-1)
			ps = append(ps, points[:i]...)
			ps = append(ps, points[i+1:]...)

			xX, mX := slices.MaxFunc(ps, f(0)), slices.MinFunc(ps, f(0))
			xY, mY := slices.MaxFunc(ps, f(1)), slices.MinFunc(ps, f(1))

			dist := max(xX[0]-mX[0], xY[1]-mY[1])
			mdist = min(mdist, dist)

			log.Print(i, dist, ps)
		}
		return mdist
	}

	log.Print("12 ?= ", minimumDistance([][]int{{3, 10}, {5, 15}, {10, 2}, {4, 4}}))
	log.Print("10 ?= ", minimumDistance([][]int{{3, 2}, {3, 9}, {7, 10}, {4, 4}, {8, 10}, {2, 7}}))
}
