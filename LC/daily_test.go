package lc

import (
	"container/heap"
	"log"
	"reflect"
	"runtime"
	"slices"
	"sort"
	"testing"
)

func init() {
	log.Print("> Daily")
}

// 930m
func Test930(t *testing.T) {
	// PrefixSum & HashMap
	countSubarraysWithSum := func(nums []int, goal int) int {
		frq := map[int]int{}

		x, pfxSum := 0, 0
		for _, n := range nums {
			pfxSum += n
			if pfxSum == goal {
				x++
			}
			if f, ok := frq[pfxSum-goal]; ok {
				x += f
			}
			frq[pfxSum]++
		}

		return x
	}

	// SlidingWindow -> space: O(1)
	countSubarraysWithSum2 := func(nums []int, goal int) int {
		x := 0

		l, csum := 0, 0
		leadingZeros := 0
		for r, n := range nums {
			csum += n

			for ; l < r && nums[l] == 0 || csum > goal; l++ {
				if nums[l] == 0 {
					leadingZeros++
				} else {
					leadingZeros = 0
				}
				csum -= nums[l]
			}

			if csum == goal {
				x += 1 + leadingZeros
			}
		}

		return x
	}

	// SlidingWindow -> space: O(1)
	countSubarraysWithSum3 := func(nums []int, goal int) int {
		// all subarrays with sum of at least: v
		atLeast := func(v int) int {
			x := 0
			l, csum := 0, 0
			for r := range nums {
				csum += nums[r]
				for l <= r && csum > v {
					csum -= nums[l]
					l++
				}
				x += r - l + 1
			}
			return x
		}
		return atLeast(goal) - atLeast(goal-1)
	}

	log.Print("4 ?= ", countSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	log.Print("8 ?= ", countSubarraysWithSum([]int{1, 0, 1, 0, 1}, 1))
	log.Print("15 ?= ", countSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))

	log.Print("4 ?= ", countSubarraysWithSum2([]int{1, 0, 1, 0, 1}, 2))

	log.Print("8 ?= ", countSubarraysWithSum3([]int{1, 0, 1, 0, 1}, 1))
	log.Print("2 ?= ", countSubarraysWithSum3([]int{1, 0, 1, 0, 1}, 0))
}

// 525m Contiguous Array
func Test525(t *testing.T) {
	findMaxLength := func(nums []int) int {
		idx := map[int]int{}
		idx[0] = -1

		x, v := 0, 0
		for i, n := range nums {
			if n == 0 {
				v--
			} else {
				v++
			}

			if start, ok := idx[v]; ok {
				x = max(x, i-start)
			} else {
				idx[v] = i
			}
		}

		return x
	}

	log.Print("2 ?= ", findMaxLength([]int{0, 1}))
	log.Print("2 ?= ", findMaxLength([]int{0, 1, 1, 0, 1, 0, 0}))
}

// 452m Min Arrows to Burst Ballons
func Test452(t *testing.T) {
	findMinArrowShots := func(points [][]int) int {
		slices.SortFunc(points, func(a, b []int) int { return a[1] - b[1] })
		log.Print(points)

		x, h := 1, points[0][1]
		for i := 1; i < len(points); i++ {
			lower, top := points[i][0], points[i][1]
			if lower > h {
				x++
				h = top
			}
		}
		return x
	}

	log.Print("2 ?= ", findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	log.Print("2 ?= ", findMinArrowShots([][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}))
}

type iHeap struct {
	sort.IntSlice
}

func (o *iHeap) Push(v any) { o.IntSlice = append(o.IntSlice, v.(int)) }
func (o *iHeap) Pop() any {
	v := o.IntSlice[len(o.IntSlice)-1]
	o.IntSlice = o.IntSlice[:len(o.IntSlice)-1]
	return v
}

type tHeapItem struct {
	tsym byte
	frq  int
}
type tHeap []*tHeapItem

func (t tHeap) Len() int { return len(t) }
func (t tHeap) Less(i, j int) bool {
	return t[i].frq > t[j].frq || t[i].frq == t[j].frq && t[i].tsym < t[j].tsym
}
func (t tHeap) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t *tHeap) Push(v any)   { *t = append(*t, v.(*tHeapItem)) }
func (t *tHeap) Pop() any {
	v := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return v
}

// 621m Task Scheduler
func Test621(t *testing.T) {
	leastInterval := func(tasks []byte, n int) int {
		frq := make([]int, 26)
		for _, b := range tasks {
			frq[b-'A']++
		}
		slices.Sort(frq)

		frqX := frq[25] - 1
		free := frqX * n
		for i := 24; i >= 0 && frq[i] > 0; i-- {
			free -= min(frqX, frq[i])
		}

		if free > 0 {
			return free + len(tasks)
		}
		return len(tasks)
	}

	// Heap
	leastInterval2 := func(tasks []byte, n int) int {
		frq := make([]int, 26)
		for _, b := range tasks {
			frq[b-'A']++
		}

		q := &tHeap{}
		for i, f := range frq {
			if f > 0 {
				heap.Push(q, &tHeapItem{'A' + byte(i), f})
			}
		}

		log.Printf("+ %d %s", n, tasks)

		schedule := []byte{}
		for q.Len() > 0 {
			tmps := []*tHeapItem{}

			for range n + 1 {
				if q.Len() > 0 {
					e := heap.Pop(q).(*tHeapItem)
					e.frq--
					if e.frq > 0 {
						tmps = append(tmps, e)
					}
					schedule = append(schedule, e.tsym)
				} else {
					if len(tmps) > 0 {
						schedule = append(schedule, '-')
					}
				}
			}

			for _, e := range tmps {
				heap.Push(q, e)
			}
		}

		log.Printf("%s", schedule)
		return len(schedule)
	}

	_ = iHeap{}
	for _, f := range []func([]byte, int) int{leastInterval, leastInterval2} {
		log.Print("--- ", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		log.Print("2 ?= ", f([]byte{'A', 'B'}, 2))
		log.Print("8 ?= ", f([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
		log.Print("6 ?= ", f([]byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1))
		log.Print("10 ?= ", f([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3))
		log.Print("14 ?= ", f([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C'}, 5))
		log.Print("10 ?= ", f([]byte{'A', 'B', 'C', 'D', 'E', 'A', 'B', 'C', 'D', 'E'}, 4))
	}
}
