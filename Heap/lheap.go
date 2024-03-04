package lheap

import (
	"container/heap"
	"log"
	"slices"
	"sort"
)

type IHeap struct{ sort.IntSlice }

func (o *IHeap) Push(v any) { o.IntSlice = append(o.IntSlice, v.(int)) }
func (o *IHeap) Pop() any {
	v := (o.IntSlice)[len(o.IntSlice)-1]
	o.IntSlice = o.IntSlice[:len(o.IntSlice)-1]
	return v
}

// 2542
func MaxScore(nums1, nums2 []int, k int) int64 {
	type E struct{ s, m int }
	L := []E{}
	for i := 0; i < len(nums1); i++ {
		L = append(L, E{s: nums1[i], m: nums2[i]})
	}
	slices.SortFunc(L, func(a, b E) int { return b.m - a.m })
	log.Print(L)

	H := IHeap{}
	score, lSum := int64(0), int64(0)
	for _, l := range L {
		heap.Push(&H, l.s)
		lSum += int64(l.s)

		if H.Len() > k {
			lSum -= int64(heap.Pop(&H).(int))
		}
		if H.Len() == k {
			score = max(score, lSum*int64(l.m))
		}
	}
	return score
}

// 948
func BagOfTokensScore(tokens []int, power int) int {
	slices.Sort(tokens)
	score := 0
	l, r := 0, len(tokens)-1
	for l <= r {
		log.Printf("|> (%d,%d) %d %d", l, r, score, power)

		if power >= tokens[l] {
			score++
			power -= tokens[l]
			l++
		} else if score > 0 && l < r {
			score--
			power += tokens[r]
			r--
		} else {
			return score
		}

		log.Printf("<| (%d,%d) %d %d", l, r, score, power)
	}
	return score
}
