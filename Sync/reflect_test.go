package lsync

import (
	"log"
	"math"
	"reflect"
	"slices"
	"sync"
	"testing"
	"time"
)

func TestReflect(t *testing.T) {
	func() {
		n := 42
		v, p := reflect.ValueOf(n), reflect.ValueOf(&n)
		log.Printf("%T %[1]v %v | %T %[3]v %v", v, v.Kind(), p, p.Kind())
	}()

	var i interface{}
	i = new(int)
	log.Printf("interface{} = new(int) -> %v, %v, %v", reflect.TypeOf(i), reflect.TypeOf(i).Kind(), reflect.ValueOf(i).Elem())

	i = new(struct{ a, b int })
	log.Printf("interface{} = new(struct{}) -> %v, %v, %v", reflect.TypeOf(i), reflect.TypeOf(i).Kind(), reflect.ValueOf(i).Elem())
	log.Printf("%v | %+v", reflect.ValueOf(i).Type(), reflect.ValueOf(i).Interface())

	log.Printf("CanSet: %v", reflect.ValueOf(i).CanSet())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			log.Print("[panic] -> ", recover())
		}()
		reflect.ValueOf(i).SetInt(1)
	}()
	wg.Wait()

	func() {
		i := int64(0)
		var e reflect.Value = reflect.ValueOf(&i).Elem()
		log.Printf("CanSet of (%v) -> ValueOf(): %v | ValueOf().Elem(): %v", reflect.TypeOf(&i), reflect.ValueOf(&i).CanSet(), e.CanSet())

		e.SetInt(42)
		log.Printf("SetInt: %v", i)
	}()

	r := reflect.TypeOf(int64(24))
	log.Printf("%T %[1]v %T %[2]v", r, reflect.New(r))
}

type RefPerson struct {
	Name string `json:"name" xml:"-" tst:"name,qfr1,qfr2"`
	Year int    `json:"year,omitempty" xml:"-"`
}

func TestRefFieldTags(t *testing.T) {
	v := reflect.ValueOf(RefPerson{"Mr Reflection", 2006})
	r := v.Type()

	log.Printf("%+v %v", v, r)
	for i := 0; i < r.NumField(); i++ {
		log.Printf("- %v: %v", r.Field(i).Name, v.Field(i))

		f := r.Field(i)
		log.Printf(":Tag:  %v | %v | %v", f.Tag.Get("json"), f.Tag.Get("xml"), f.Tag.Get("tst"))
	}
}

func TestRefMap(t *testing.T) {
	m := map[string]int{"a": 0, "b": 1, "c": 2, "z": 25}
	v := reflect.ValueOf(m)
	r := reflect.TypeOf(m)

	log.Printf("%v, Kind: %v, %v, %v", m, v.Kind(), v, r)

	for _, k := range v.MapKeys() {
		log.Print(k, v.MapIndex(k))
	}

	v.SetMapIndex(reflect.ValueOf("y"), reflect.ValueOf(24))

	itr := v.MapRange()
	log.Printf("%T -> %[1]v", itr)
	for itr.Next() {
		log.Print(itr.Key(), itr.Value())
	}
}

// 791
func TestCustomSortString(t *testing.T) {
	customSortString := func(order string, s string) string {
		m := map[byte]int{}
		for _, r := range s {
			m[byte(r)]++
		}
		log.Print(m)

		k, bs := 0, make([]byte, len(s))
		for _, r := range order {
			for f := m[byte(r)]; f > 0; f-- {
				bs[k] = byte(r)
				k++
			}
			m[byte(r)] = 0
		}
		log.Print(m)

		for b, f := range m {
			for ; f > 0; f-- {
				bs[k] = b
				k++
			}
		}
		return string(bs)
	}

	log.Print("abcd -> ", customSortString("bcafg", "abcd"))
}

// BinSearch template
/*
  func Cond(int) bool
  func BinSearch() int {
    l, r := 0, len(Space)
    for l < r {
      m := l + (r-l)>>1
      if Cond(m) {
        r = m
      } else {
        l = m+1
      }
    }
    return l
  }
*/

// 69: 4->2, 10->3, ... (x > 1)
func TestIntSqrt(t *testing.T) {
	// k^2 <= x -> k is int sqrt of x
	intSqrt := func(x int) int {
		l, r := 1, x
		for l < r {
			m := l + (r-l)>>1
			if m*m > x {
				r = m
			} else {
				l = m + 1
			}
		}
		return l - 1
	}

	for _, x := range []int{4, 8, 10, 35, 81, 83, 3448230483} {
		log.Print(x, " -> ", intSqrt(x))
	}
}

// 278
func Test278(t *testing.T) {
	isBad := func(n int) bool { return n > 7 }
	firstBadVersion := func(x int) int {
		l, r := 1, x
		for l < r {
			m := l + (r-l)>>1
			if isBad(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	ts := time.Now()
	log.Print(firstBadVersion(math.MaxInt), time.Since(ts))
}

// 35: search insert position
func Test35(t *testing.T) {
	searchInsert := func(nums []int, x int) int {
		l, r := 0, len(nums)
		for l < r {
			m := l + (r-l)>>1
			if nums[m] >= x {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	nums := []int{1, 3, 3, 3, 5, 5, 6}
	log.Print(nums)
	for _, n := range []int{0, 1, 2, 5, 6, 7} {
		log.Print(n, " -> ", searchInsert(nums, n))
	}
}

// 1011m -> minimum ship capacity
func Test1011(t *testing.T) {
	shipWithinDays := func(weights []int, days int) int {
		canShip := func(capacity int) bool {
			d, c := 1, 0
			for _, w := range weights {
				c += w
				if c > capacity {
					c = w
					d++
					if d > days {
						return false
					}
				}
			}
			return true
		}

		xCap := 0
		for _, w := range weights {
			xCap += w
		}

		l, r := slices.Max(weights), xCap
		for l < r {
			m := l + (r-l)>>1
			if canShip(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	ws := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, days := range []int{5, 7} {
		log.Print("Minimum Ship Capacity -> ", shipWithinDays(ws, days))
	}
}

// 410h Split Array Largest Sum
/*
[7 2 5 10 8] -2-> [7 2 5] [10 8] : 18
*/
func Test410(t *testing.T) {
	splitArray := func(nums []int, m int) int {
		sm := 0
		for _, n := range nums {
			sm += n
		}

		isValid := func(aSum int) bool {
			lSum, count := 0, 1
			for _, n := range nums {
				lSum += n
				if lSum > aSum {
					count++
					lSum = n
					if count > m {
						return false
					}
				}
			}
			return true
		}

		l, r := slices.Max(nums), sm
		for l < r {
			m := l + (r-l)>>1
			log.Printf("%2d %2d %2d", l, m, r)

			if isValid(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	log.Printf("18 -> %t", splitArray([]int{7, 2, 5, 10, 8}, 2) == 18)
}

// 875m Koko Eating Bananas
func Test875(t *testing.T) {
	minEatingSpeed := func(piles []int, hours int) int {
		fastEnough := func(speed int) bool {
			h := 0
			for _, p := range piles {
				h += (p-1)/speed + 1
				if h > hours {
					return false
				}
			}
			return true
		}

		l, r := 1, slices.Max(piles)
		for l < r {
			m := l + +(r-l)>>1
			log.Printf("%2d %2d %2d", l, m, r)

			if fastEnough(m) {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	log.Print("+ 4 -> ", minEatingSpeed([]int{3, 6, 7, 11}, 8))
	log.Print("+ 23 -> ", minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
	log.Print("+ 30 -> ", minEatingSpeed([]int{30}, 1))
}

// 1482m
func Test1482(t *testing.T) {
	minDays := func(bloomDay []int, m int, k int) int {
		if len(bloomDay) < m*k {
			return -1
		}

		flowers := func(day int) int {
			b, fadj := 0, 0
			for _, d := range bloomDay {
				if d <= day {
					b += (fadj + 1) / k
					fadj = (fadj + 1) % k
				} else {
					fadj = 0
				}
			}
			return b
		}

		l, r := slices.Min(bloomDay), slices.Max(bloomDay)
		for l < r {
			md := l + (r-l)>>1
			if flowers(md) >= m {
				r = md
			} else {
				l = md + 1
			}
		}
		return l
	}

	log.Print("3 =? ", minDays([]int{1, 10, 3, 10, 2}, 3, 1))
	log.Print("-1 =? ", minDays([]int{1, 10, 3, 10, 2}, 3, 2))
	log.Print("12 =? ", minDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3))
}

// 668h
func Test668(t *testing.T) {
	findKthNumber := func(m int, n int, k int) int {
		kth := func(v int) int {
			x := 0
			for i := 1; i <= m; i++ {
				x += min(n, v/i)
			}
			return x
		}

		l, r := 1, m*n
		for l < r {
			m := l + (r-l)>>1
			if kth(m) >= k {
				r = m
			} else {
				l = m + 1
			}
		}
		return l
	}

	log.Print("3 =? ", findKthNumber(3, 3, 5))
	log.Print("6 =? ", findKthNumber(2, 3, 6))
	log.Print("? =? ", findKthNumber(102, 394, 19299))
}
