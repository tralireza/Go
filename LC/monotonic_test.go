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
	price, span []int
}

func NewStockSpanner() StockSpanner { return StockSpanner{[]int{}, []int{}} }
func (o *StockSpanner) Next(price int) int {
	span := 1
	for len(o.price) > 0 && price >= o.price[len(o.price)-1] {
		span += o.span[len(o.span)-1]
		o.span = o.span[:len(o.span)-1]    // Pop
		o.price = o.price[:len(o.price)-1] // Pop
	}

	o.span = append(o.span, span)
	o.price = append(o.price, price)

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
	log.Print(o.price, o.span)
	log.Print(price, " -> ", span)
}
