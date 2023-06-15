package lpq

import (
	"container/heap"
	"log"
	"sort"
)

type IHeap struct{ sort.IntSlice }

func (h *IHeap) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *IHeap) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func TotalCost(costs []int, k int, candidates int) int64 {
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
