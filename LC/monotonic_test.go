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
	S []int
}

func NewStockSpanner() StockSpanner { return StockSpanner{S: []int{}} }
func (o *StockSpanner) Next(price int) int {
	o.S = append(o.S, price)
	span := 0
	for i := len(o.S) - 1; i >= 0 && price >= o.S[i]; i-- {
		span++
	}
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
