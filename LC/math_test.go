package lc

import (
	"log"
	"math"
	"testing"
)

func init() {
	log.Print("> Math")
}

// 3102h Minimize Manhattan Distances
func Test3102(t *testing.T) {
	minimumDistance := func(points [][]int) int {
		xdst := func(iSkip int) (d, i, j int) {
			x, n, xi, ni := math.MinInt, math.MaxInt, 0, 0
			xd, nd, xdi, ndi := math.MinInt, math.MaxInt, 0, 0

			for i, p := range points {
				if i == iSkip {
					continue
				}

				// y=x diagonal
				if x < p[0]+p[1] {
					x, xi = p[0]+p[1], i
				}
				if n > p[0]+p[1] {
					n, xi = p[0]+p[1], i
				}

				// y=-x diagonal
				if xd < p[0]-p[1] {
					xd, xdi = p[0]-p[1], i
				}
				if nd > p[0]-p[1] {
					nd, xdi = p[0]-p[1], i
				}
			}

			if x-n >= xd-nd {
				log.Print("y=x Diagonal")
				return x - n, xi, ni
			}
			log.Print("y=-x Diagonal")
			return xd - nd, xdi, ndi
		}

		x, i, j := xdst(-1)
		log.Print(x, points[i], points[j])

		d1, _, _ := xdst(i)
		d2, _, _ := xdst(j)
		return min(d1, d2)
	}

	log.Print("0 ?= ", minimumDistance([][]int{{1, 1}, {1, 1}, {1, 1}}))
	log.Print("12 ?= ", minimumDistance([][]int{{3, 10}, {5, 15}, {10, 2}, {4, 4}}))
	log.Print("10 ?= ", minimumDistance([][]int{{3, 2}, {3, 9}, {7, 10}, {4, 4}, {8, 10}, {2, 7}}))
}

// 205 Isomorphic Strings
func Test205(t *testing.T) {
	isIsomorphic := func(s, t string) bool {
		ms, mt := map[byte]byte{}, map[byte]byte{}

		for i := 0; i < len(s) && i < len(t); i++ {
			log.Printf("%q %q", s[i], t[i])

			if ms[s[i]] == 0 {
				ms[s[i]] = t[i]
			}
			if mt[t[i]] == 0 {
				mt[t[i]] = s[i]
			}

			if ms[s[i]] != t[i] || mt[t[i]] != s[i] {
				return false
			}
		}

		log.Printf("%q\n%q", ms, mt)
		return true
	}

	log.Print("true ?= ", isIsomorphic("egg", "add"))
	log.Print("false ?= ", isIsomorphic("foo", "bar"))
	log.Print("false ?= ", isIsomorphic("aba", "xxy"))
	log.Print("true ?= ", isIsomorphic("abcdefghij", "0123456789"))
}

// 189m Rotate Array
func Test189(t *testing.T) {
	rotate := func(nums []int, k int) {
		k = k % len(nums)
		if k == 0 {
			return
		}

		gcd := func(a, b int) int {
			for b > 0 {
				a, b = b, a%b
			}
			return a
		}
		log.Printf(".GCD(%d,%d) %d", len(nums), k, gcd(len(nums), k))

		reverse := func(l, r int) {
			for l < r {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r--
			}
		}

		reverse(0, len(nums)-1)
		reverse(0, k-1)
		reverse(k, len(nums)-1)
	}

	for _, ts := range [][][]int{
		{{1, 2, 3, 4, 5, 6}, {4}},
		{{1, 2, 3, 4, 5, 6}, {5}},
		{{1, 2, 3, 4, 5, 6}, {1}},
		{{1, 2, 3, 4, 5, 6, 7}, {3}},
	} {
		nums, k := ts[0], ts[1][0]
		rotate(nums, k)
		log.Printf("+ k: %d -> %v", k, nums)
	}
}
