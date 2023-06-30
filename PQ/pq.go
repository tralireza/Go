package lpq

import (
	"container/heap"
	"log"
	"sort"
)

type SHe struct{ Cost, Side int }
type SHeap []SHe

func (h SHeap) Len() int { return len(h) }
func (h SHeap) Less(i, j int) bool {
	if h[i].Cost == h[j].Cost {
		return h[i].Side < h[j].Side
	}
	return h[i].Cost < h[j].Cost
}
func (h SHeap) Swap(i int, j int) { h[i], h[j] = h[j], h[i] }
func (h *SHeap) Push(x any)       { *h = append(*h, x.(SHe)) }
func (h *SHeap) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func TotalCost1(costs []int, k int, candidates int) int64 {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	H := &SHeap{}
	l, r := 0, len(costs)-1
	for ; l < candidates; l++ {
		heap.Push(H, SHe{costs[l], 0})
	}
	for ; r >= max(len(costs)-candidates, candidates); r-- {
		heap.Push(H, SHe{costs[r], 1})
	}
	log.Print(H)

	tcost := int64(0)
	for k > 0 {
		e := heap.Pop(H).(SHe)
		tcost += int64(e.Cost)
		if l <= r {
			if e.Side == 0 {
				heap.Push(H, SHe{costs[l], e.Side})
				l++
			} else {
				heap.Push(H, SHe{costs[r], e.Side})
				r--
			}
		}
		k--
	}
	return tcost
}

type IHeap struct{ sort.IntSlice }

func (h *IHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *IHeap) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func TotalCost(costs []int, k int, candidates int) int64 {
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	L, R := IHeap{}, IHeap{}
	l, r := 0, len(costs)-1
	for ; l < candidates; l++ {
		heap.Push(&L, costs[l])
	}
	for ; r >= max(len(costs)-candidates, candidates); r-- {
		heap.Push(&R, costs[r])
	}
	log.Printf("(%d,%d) %v %v", l, r, L, R)

	tcost := int64(0)
	for k > 0 {
		if R.Len() == 0 || L.Len() > 0 && L.IntSlice[0] <= R.IntSlice[0] {
			tcost += int64(heap.Pop(&L).(int))
			if l <= r {
				heap.Push(&L, costs[l])
				l++
			}
		} else {
			tcost += int64(heap.Pop(&R).(int))
			if l <= r {
				heap.Push(&R, costs[r])
				r--
			}
		}
		k--
	}
	return tcost
}
