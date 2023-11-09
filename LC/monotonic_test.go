package lc

import (
	"log"
	"testing"
)

func init() {
	log.Print("> Monotonic Stack")
}

// 901m Online Stock Span
type StockSpanner struct {
	p, s []int
}

func NewStockSpanner() StockSpanner { return StockSpanner{[]int{}, []int{}} }
func (o *StockSpanner) Next(price int) int {
	span := 1
	for len(o.p) > 0 && price >= o.p[len(o.p)-1] {
		span += o.s[len(o.s)-1]
		o.s = o.s[:len(o.s)-1] // Pop
		o.p = o.p[:len(o.p)-1] // Pop
	}

	o.s = append(o.s, span)
	o.p = append(o.p, price)

	return span
}

// [100 80 60 70 60 75 85] -> [1 1 1 2 1 4 6]
func Test901(t *testing.T) {
	o := NewStockSpanner()
	span := []int{}
	price := []int{}
	for _, p := range []int{100, 80, 60, 70, 60, 75, 85} {
		price = append(price, p)
		span = append(span, o.Next(p))
	}
	log.Print(price, " -> ", span)
}
