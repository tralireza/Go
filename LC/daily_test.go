package lc

import (
	"container/heap"
	"container/list"
	"fmt"
	"log"
	"math"
	"math/rand"
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
		log.Print("14 ?= ", f([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'D'}, 5))
		log.Print("10 ?= ", f([]byte{'A', 'B', 'C', 'D', 'E', 'A', 'B', 'C', 'D', 'E'}, 4))
	}
}

// 1669m Merge In Between Link Lists
func Test1669(t *testing.T) {
	type ListNode struct {
		Val  int
		Next *ListNode
	}

	mergeInBetween := func(list1 *ListNode, a, b int, list2 *ListNode) *ListNode {
		var n, start, end *ListNode

		n = list1
		for i := 0; i < b; i++ {
			if i == a-1 {
				start = n
			}
			n = n.Next
		}
		end = n
		log.Print(start, end)

		start.Next = list2
		n = list2
		for ; n.Next != nil; n = n.Next {
		}
		n.Next = end.Next

		return list1
	}

	type N = ListNode
	ls1 := &N{10, &N{1, &N{13, &N{6, &N{9, &N{5, nil}}}}}}
	ls2 := &N{1000000, &N{1000001, &N{1000002, nil}}}
	for n := mergeInBetween(ls1, 3, 4, ls2); n != nil; n = n.Next {
		fmt.Print(*n, " ")
	}
	fmt.Println("X")
}

// 206 Reserve Linked List
func Test206(t *testing.T) {
	type ListNode struct {
		Val  int
		Next *ListNode
	}

	reverseList := func(head *ListNode) *ListNode {
		var prv *ListNode

		n := head
		for n != nil {
			nxt := n.Next
			n.Next = prv
			prv = n
			n = nxt
		}

		return prv
	}

	type N = ListNode
	var l *N = nil
	for i := 0; i < 9; i++ {
		r := reverseList(l)
		fmt.Print("List: ")
		for n := r; n != nil; n = n.Next {
			fmt.Printf("{%v +} ", n.Val)
		}
		fmt.Println(" X")
		l = &N{i, r}
	}
}

// 143m Reorder List
func Test143(t *testing.T) {
	type ListNode struct {
		Val  int
		Next *ListNode
	}

	// [1 2 3 4 5 6 7 8] -> [1 8 2 7 3 6 4 5]
	reorderList := func(head *ListNode) {
		S := []*ListNode{}

		m, f := head, head
		for ; f != nil && f.Next != nil; f = f.Next.Next {
			m = m.Next
		}
		if f != nil && f.Next == nil {
			m = m.Next
		}
		for n := m; n != nil; n = n.Next {
			S = append(S, n)
		}

		n := head
		for len(S) > 0 {
			r := S[len(S)-1]
			S = S[:len(S)-1]

			nxt := n.Next
			n.Next = r
			r.Next = nxt
			n = nxt
		}
		n.Next = nil
	}

	draw := func(h *ListNode) {
		for n := h; n != nil; n = n.Next {
			fmt.Printf("{%d +} ", n.Val)
		}
		fmt.Println("X")
	}

	type N = ListNode
	l := &N{1, &N{2, &N{3, &N{4, &N{5, &N{6, &N{7, &N{8, nil}}}}}}}}
	reorderList(l)
	fmt.Print("[1 8 2 7 3 6 4 5] -> ")
	draw(l)

	r := &N{0, l}
	reorderList(r)
	fmt.Print("[0 5 1 4 8 6 2 3 7] -> ")
	draw(r)
}

// 287m Find Duplicate Number
func Test287(t *testing.T) {
	// BinSearch
	findDuplicate := func(nums []int) int {
		lessOrEqual := func(v int) bool {
			f := 0
			for _, n := range nums {
				if n <= v {
					f++
				}
			}
			return f > v
		}

		l, r := 1, len(nums)-1
		for l < r {
			m := l + (r-l)>>1
			if lessOrEqual(m) {
				r = m
			} else {
				l = m + 1
			}
		}

		return l
	}

	log.Print("2 ?= ", findDuplicate([]int{1, 3, 4, 2, 2}))
	log.Print("1 ?= ", findDuplicate([]int{1, 3, 4, 2, 1}))
	log.Print("4 ?= ", findDuplicate([]int{4, 3, 1, 4, 2}))
	log.Print("1 ?= ", findDuplicate([]int{1, 1}))
}

// 80m Remove Duplicates from Sorted Array II
func Test80(t *testing.T) {
	// at most 2 instances of any duplicates
	removeDuplicates := func(nums []int) int {
		l := 2
		for _, n := range nums[2:] {
			if n > nums[l-1] {
				nums[l] = n
				l++
			} else if n > nums[l-2] {
				nums[l] = n
				l++
			}
		}
		return l
	}

	sarr := []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 5, 6}
	log.Print("[1 1 2 2 3 4 4 5 6] ?= ", sarr[:removeDuplicates(sarr)])
}

// 172m Factorial Leading Zeroes
func Test172(t *testing.T) {
	trailingZeroes := func(n int) int {
		z := 0
		for n > 0 {
			n /= 5
			z += n
		}
		return z
	}

	log.Print("0 ?= ", trailingZeroes(3))
	log.Print("1 ?= ", trailingZeroes(7))
	log.Print("2 ?= ", trailingZeroes(12)) // 12! = 479001600
	log.Print(" ?= ", trailingZeroes(120))
}

// 442m Find all Duplicates in an Array
func Test442(t *testing.T) {
	// all elements between 1..n with duplicates
	// O(n) time & O(1) space
	findDuplicates := func(nums []int) []int {
		dups := []int{}
		for _, n := range nums {
			if n < 0 {
				n = -n
			}
			if nums[n-1] < 0 {
				dups = append(dups, n)
			}

			nums[n-1] = -nums[n-1]
		}
		return dups
	}

	log.Print("[2 3] ?= ", findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

// 637 Average of Levels in Binary Tree
func Test637(t *testing.T) {
	type TreeNode struct {
		Val         int
		Left, Right *TreeNode
	}

	averageOfLevels := func(root *TreeNode) []float64 {
		rst := []float64{}

		queue := []*TreeNode{root}
		for len(queue) > 0 {
			l := len(queue)
			var v float64
			for range l {
				n := queue[0]
				queue = queue[1:]
				v += float64(n.Val)

				if n.Left != nil {
					queue = append(queue, n.Left)
				}
				if n.Right != nil {
					queue = append(queue, n.Right)
				}
			}

			rst = append(rst, v/float64(l))
		}

		return rst
	}

	type T = TreeNode
	log.Print(averageOfLevels(&T{3, &T{Val: 9}, &T{20, &T{Val: 15}, &T{Val: 7}}}))
}

// 74m Search in a 2D Matrix
func Test74(t *testing.T) {
	searchMatrix := func(matrix [][]int, target int) bool {
		p, q := len(matrix), len(matrix[0])
		l, r := 0, p*q
		for l < r {
			m := l + (r-l)>>1
			if matrix[m/q][m%q] >= target {
				r = m
			} else {
				l = m + 1
			}
		}
		return l < p*q && matrix[l/q][l%q] == target
	}

	log.Print("true ?= ", searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 11))
	log.Print("false ?= ", searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 0))
	log.Print("false ?= ", searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	log.Print("false ?= ", searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 61))
}

// 41h First Missing Positive
func Test41(t *testing.T) {
	firstMissingPositive := func(nums []int) int {
		posCount := 0
		for i, n := range nums {
			if n <= 0 {
				nums[i] = math.MaxInt
			} else {
				posCount++
			}
		}

		for _, n := range nums {
			if n < 0 {
				n = -n
			}
			if n <= len(nums) && nums[n-1] > 0 {
				nums[n-1] = -nums[n-1]
			}
		}

		for i, n := range nums {
			if n > 0 {
				return i + 1
			}
		}
		return posCount + 1
	}

	log.Print("2 ?= ", firstMissingPositive([]int{-1, -3, -7, -8, -9, 1}))
	log.Print("1 ?= ", firstMissingPositive([]int{-6, 4, -6, 4, 3, 3, -6, 4, 0, 9, 7}))
}

// Cycle Sort
func TestCycleSort(t *testing.T) {
	// Sort [1..n] in-place ignoring zeroes, negatives && duplicates
	cycleSort := func(nums []int) {
		i := 0
		for i < len(nums) {
			n := nums[i]
			if 0 < n && n <= len(nums) && nums[n-1] != n {
				nums[n-1], nums[i] = n, nums[n-1]
			} else {
				i++
			}
		}
	}

	nums := []int{1, 3, 4, 9, 0, 9, -1, 4, 7, -3, 17, 2, 0, 10}
	log.Print(nums)
	cycleSort(nums)
	log.Print(nums)
}

// Moving all Non-Positive Numbers to Right
func TestAdjust(t *testing.T) {
	adjust := func(nums []int) int {
		l, r := 0, len(nums)-1
		for ; l < r; l++ {
			if nums[l] <= 0 {
				nums[l], nums[r] = nums[r], nums[l]
				r--
				if nums[l] <= 0 {
					l--
				}
			}
		}
		if l < len(nums) && nums[l] > 0 {
			l++
		}
		return l
	}

	for i := 0; i < 1000; i++ {
		nums := make([]int, rand.Intn(25))
		frq := 0
		positives := []int{}
		for i := range nums {
			n := rand.Intn(len(nums))
			if rand.Intn(2) == 1 {
				n = -n
			}
			if n > 0 {
				positives = append(positives, n)
				frq++
			}
			nums[i] = n
		}

		log.Print(nums)
		l := adjust(nums)
		log.Print(nums)
		log.Print(frq, " ?= ", len(nums[:l]), nums[:l])
		if frq != l {
			t.Fatalf("Wrong length: %d != %d", l, frq)
		}
		log.Print(positives)
		log.Print("---")
	}
}

// 713m Subarray Product Less Than K
func Test713(t *testing.T) {
	numSubarrayProductLessThanK := func(nums []int, k int) int {
		count := 0

		l := 0
		p := 1
		for r, n := range nums {
			p *= n

			for p >= k {
				p /= nums[l]
				l++
			}

			count += r - l + 1
		}

		return count
	}

	log.Print("8 ?= ", numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}

// 2958m Length Of Longest Subarray With at Most K Frequency
func Test2958(t *testing.T) {
	maxSubarrayLength := func(nums []int, k int) int {
		frq := map[int]int{}
		x := 0

		l := 0
		for r, n := range nums {
			frq[n]++
			if frq[n] <= k {
				x = max(x, r-l+1)
			} else {
				for l < r && frq[n] > k {
					if frq[l] > 0 {
						frq[nums[l]]--
					}
					l++
				}
			}
		}

		return x
	}

	maxSubarrayLength2 := func(nums []int, k int) int {
		frq := map[int]int{}
		x := 0

		l := 0
		countKVals := 0
		for r, n := range nums {
			frq[n]++
			if frq[n] <= k {
				if countKVals == 0 {
					x = max(x, r-l+1)
				}
			} else {
				countKVals++
				if frq[nums[l]] > 0 {
					frq[nums[l]]--
					if frq[nums[l]] <= k {
						countKVals--
					}
				}
				l++
			}
		}

		return x
	}

	for _, f := range []func([]int, int) int{maxSubarrayLength, maxSubarrayLength2} {
		log.Print("6 ?= ", f([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2))
		log.Print("4 ?= ", f([]int{5, 5, 5, 5, 5, 5, 5}, 4))
	}
}

// 1971 Find if Path Exists in Graph
func Test1971(t *testing.T) {
	validPath := func(n int, edges [][]int, source int, destination int) bool {
		DJS := make([][2]int, n) // [2]int: 0: Parent, 1: Rank
		for i := range DJS {
			DJS[i][0] = i
		}

		var find func([][2]int, int) int
		find = func(DJS [][2]int, x int) int {
			if DJS[x][0] != x {
				DJS[x][0] = find(DJS, DJS[x][0])
			}
			return DJS[x][0]
		}

		link := func(DJS [][2]int, x, y int) {
			X, Y := find(DJS, x), find(DJS, y)
			if X != Y {
				if DJS[X][1] < DJS[Y][1] {
					DJS[X][0] = Y
				} else {
					DJS[Y][0] = X
					if DJS[X][1] == DJS[Y][1] {
						DJS[X][1]++
					}
				}
			}
		}

		for _, p := range edges {
			link(DJS, p[0], p[1])
		}
		log.Print(DJS)

		return find(DJS, source) == find(DJS, destination)
	}

	log.Print("true ?= ", validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	log.Print("false ?= ", validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
}

// 752m Open Lock
func Test752(t *testing.T) {
	openLock := func(deadends []string, target string) int {
		Space := make([][][][]byte, 10)
		for x := range Space {
			Space[x] = make([][][]byte, 10)
			for y := range Space[x] {
				Space[x][y] = make([][]byte, 10)
				for z := range Space[x][y] {
					Space[x][y][z] = make([]byte, 10)
				}
			}
		}

		b := []byte(target)
		Space[b[0]-'0'][b[1]-'0'][b[2]-'0'][b[3]-'0'] = '*'
		if Space[0][0][0][0] == '*' {
			return 0
		}

		for _, deadend := range deadends {
			b := []byte(deadend)
			Space[b[0]-'0'][b[1]-'0'][b[2]-'0'][b[3]-'0'] = 'X'
		}
		if Space[0][0][0][0] == 'X' {
			return -1
		}

		Q, lvl := [][4]byte{}, 1
		var x [4]byte
		for len(Q) > 0 {
			for range len(Q) {
				x, Q = Q[0], Q[1:]
				for i := 0; i < 4; i++ {
					v := x[i]

					for _, m := range []byte{1, 9} {
						x[i] = (x[i] + m) % 10
						if Space[x[0]][x[1]][x[2]][x[3]] == '*' {
							return lvl
						}

						if Space[x[0]][x[1]][x[2]][x[3]] == 0 {
							Space[x[0]][x[1]][x[2]][x[3]] = '-'
							Q = append(Q, x)
						}
						x[i] = v
					}
				}
			}
			lvl++
		}

		return -1
	}

	mapAndList := func(deadends []string, target string) int {
		Visited := map[string]struct{}{}

		if target == "0000" {
			return 0
		}

		for _, n := range deadends {
			Visited[n] = struct{}{}
		}
		if _, ok := Visited["0000"]; ok {
			return -1
		}
		if _, ok := Visited[target]; ok {
			return -1
		}

		Q := list.New()
		Q.PushBack("0000")

		Visited["0000"] = struct{}{}

		lvl := 1
		for Q.Len() > 0 {
			for range Q.Len() {
				n := Q.Remove(Q.Front()).(string)

				B := []byte(n)
				for i := 0; i < 4; i++ {
					for _, m := range []byte{1, 9} {
						B[i] = '0' + (B[i]-'0'+m)%10
						v := string(B)

						if v == target {
							log.Print(len(Visited))
							return lvl
						}
						if _, ok := Visited[v]; !ok {
							Visited[v] = struct{}{}
							Q.PushBack(v)
						}

						B[i] = n[i]
					}
				}
			}

			lvl++
		}

		log.Print(len(Visited))
		return -1
	}

	for _, f := range []func([]string, string) int{openLock, mapAndList} {
		log.Print(" 6 ?= ", f([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
		log.Print("-1 ?= ", f([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
		log.Print("-1 ?= ", f([]string{"0000"}, "1234"))
		log.Print(" 0 ?= ", f([]string{}, "0000"))
		log.Print("20 ?= ", f([]string{}, "5555"))
	}
}

// 310m Minimum Height Trees
func Test310(t *testing.T) {
	findMinHeightTrees := func(n int, edges [][]int) []int {
		tree := make([][]byte, n)
		for i := range tree {
			tree[i] = make([]byte, n)
		}
		for _, edge := range edges {
			v, u := edge[0], edge[1]
			tree[v][u] = 1
			tree[u][v] = 1
		}

		heights, mh := make([]int, n), n
		for v := 0; v < n; v++ {
			Vis := make([]bool, n)
			Vis[v] = true

			Q, h := []int{v}, 0
			for len(Q) > 0 {
				for range len(Q) {
					v := Q[0]
					Q = Q[1:]
					for u, y := range tree[v] {
						if y == 1 && !Vis[u] {
							Vis[u] = true
							Q = append(Q, u)
						}
					}
				}
				h++
			}

			mh = min(mh, h)
			heights[v] = h
		}

		roots := []int{}
		for v, h := range heights {
			if h == mh {
				roots = append(roots, v)
			}
		}
		return roots
	}

	graphAdj := func(n int, edges [][]int) []int {
		G := map[int]map[int]struct{}{}
		for _, edge := range edges {
			v, u := edge[0], edge[1]
			if _, ok := G[v]; !ok {
				G[v] = map[int]struct{}{}
			}
			G[v][u] = struct{}{}
			if _, ok := G[u]; !ok {
				G[u] = map[int]struct{}{}
			}
			G[u][v] = struct{}{}
		}

		heights := make([]int, n)
		for v := 0; v < n; v++ {
			Vis := make([]bool, n)
			Vis[v] = true

			Q, h := list.New(), 0
			Q.PushBack(v)
			for Q.Len() > 0 {
				for range Q.Len() {
					v := Q.Remove(Q.Front()).(int)
					for u := range G[v] {
						if !Vis[u] {
							Vis[u] = true
							Q.PushBack(u)
						}
					}
				}
				h++
			}

			heights[v] = h
		}

		roots, hMin := []int{}, slices.Min(heights)
		for v, h := range heights {
			if h == hMin {
				roots = append(roots, v)
			}
		}
		return roots
	}

	onlyTwo := func(n int, edges [][]int) []int {
		if n == 1 {
			return []int{0}
		}

		G := map[int][]int{}
		Degree := make([]int, n)
		for _, edge := range edges {
			v, u := edge[0], edge[1]
			Degree[u]++
			Degree[v]++
			G[v] = append(G[v], u)
			G[u] = append(G[u], v)
		}

		var Vis []bool
		var dG []int

		var dfs func(int, []int)
		dfs = func(v int, D []int) {
			if len(D) > len(dG) {
				dG = D
			}
			for _, u := range G[v] {
				if !Vis[u] {
					Vis[u] = true
					dfs(u, append(D, u))
				}
			}
		}

		for v := range Degree {
			if Degree[v] > 1 {
				continue
			}

			Vis = make([]bool, n)
			Vis[v] = true

			dfs(v, []int{v})
		}

		log.Print(dG)

		if len(dG)&1 == 1 {
			return []int{dG[len(dG)>>1]}
		}
		return []int{dG[len(dG)>>1-1], dG[len(dG)>>1]}
	}

	degreeCheck := func(n int, edges [][]int) []int {
		if n == 1 {
			return []int{0}
		}

		G, D := make([][]int, n), make([]int, n)
		for _, edge := range edges {
			v, u := edge[0], edge[1]
			G[v] = append(G[v], u)
			G[u] = append(G[u], v)
			D[v]++
			D[u]++
		}

		leaves := []int{}
		for v := range D {
			if D[v] == 1 {
				leaves = append(leaves, v)
			}
		}

		for n > 2 {
			n -= len(leaves)
			nQueue := []int{}
			for _, leaf := range leaves {
				for _, u := range G[leaf] {
					D[u]--
					if D[u] == 1 {
						nQueue = append(nQueue, u)
					}
				}
			}
			leaves = nQueue
		}

		return leaves
	}

	for _, f := range []func(int, [][]int) []int{findMinHeightTrees, graphAdj, onlyTwo, degreeCheck} {
		log.Print("[0] ?= ", f(1, [][]int{}))
		log.Print("[0 1] ?= ", f(2, [][]int{{1, 0}}))
		log.Print("[0] ?= ", f(3, [][]int{{1, 0}, {0, 2}}))
		log.Print("[1] ?= ", f(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
		log.Print("[3 4] ?= ", f(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
		log.Print("[1 2] ?= ", f(7, [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}))
		log.Print("[0] ?= ", f(8, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 4}, {4, 5}, {4, 6}, {6, 7}}))
		log.Print("[3] ?= ", f(10, [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 3}, {4, 6}, {4, 7}, {5, 8}, {5, 9}}))
		log.Print("===")
	}
}

// 2370m Longest Ideal Subsequence
func Test2370(t *testing.T) {
	longestIdealString := func(s string, k int) int {
		D := [26]int{}
		D[s[0]-'a'] = 1

		for i := 1; i < len(s); i++ {
			cur := int(s[i] - 'a')
			curX := 0
			for prv := 0; prv < 26; prv++ {
				if prv-cur <= k && -k <= prv-cur {
					curX = max(curX, D[prv])
				}
			}
			D[cur] = curX + 1
		}

		return slices.Max(D[:])
	}

	bottomUp := func(s string, k int) int {
		D := make([][26]int, len(s))
		D[0][s[0]-'a'] = 1
		log.Print(D[0])

		for i := 1; i < len(s); i++ {
			for prv := range 26 {
				D[i][prv] = D[i-1][prv]
			}

			l, r := max(0, int(s[i]-'a')-k), min(25, int(s[i]-'a')+k)
			for prv := l; prv <= r; prv++ {
				D[i][s[i]-'a'] = max(D[i][s[i]-'a'], D[i-1][prv]+1)
			}

			log.Print(D[i])
		}

		return slices.Max(D[len(s)-1][:])
	}

	topDown := func(s string, k int) int {
		D := make([][26]int, len(s))
		for i := range len(s) {
			for c := range 26 {
				D[i][c] = -1
			}
		}

		var findMax func(i, c int) int
		findMax = func(i, c int) int {
			if D[i][c] != -1 {
				return D[i][c]
			}

			D[i][c] = 0
			if c == int(s[i]-'a') {
				D[i][c] = 1
			}

			if i > 0 {
				D[i][c] = findMax(i-1, c)
				if c == int(s[i]-'a') {
					for p := max(0, c-k); p <= min(25, c+k); p++ {
						D[i][c] = max(D[i][c], findMax(i-1, p)+1)
					}
				}
			}

			return D[i][c]
		}

		x := 0
		for c := 0; c < 26; c++ {
			x = max(x, findMax(len(s)-1, c))
		}
		return x
	}

	for _, f := range []func(string, int) int{longestIdealString, bottomUp, topDown} {
		log.Print("4 ?= ", f("acfgbd", 2))
		log.Print("4 ?= ", f("abcd", 4))
		log.Print("2 ?= ", f("pvjcci", 4))
		log.Print("42 ?= ", f("dyyonfvzsretqxucmavxegvlnnjubqnwrhwikmnnrpzdovjxqdsatwzpdjdsneyqvtvorlwbkb", 7))
	}
}
