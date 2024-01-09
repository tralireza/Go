package lrcp

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestCall(t *testing.T) {
	go RunServer()
	time.Sleep(150 * time.Millisecond)

	errc := make(chan error)
	for i := 0; i < 15; i++ {
		q := i
		go func() { errc <- Client(q) }()
	}
	for i := 0; i < 15; i++ {
		<-errc
	}
}

func Knapsack(ksCapacity int) int {
	type Fruit struct {
		Size, Price int
		Name        string
	}

	fruits := []Fruit{{4, 4500, "Plum"},
		{5, 5700, "Apple"},
		{2, 2250, "Orange"},
		{1, 1100, "Strawberry"},
		{6, 6700, "Melon"}}

	items := make([]int, ksCapacity+1)
	values := make([]int, ksCapacity+1)

	for i := 0; i < len(fruits); i++ {
		for j := fruits[i].Size; j <= ksCapacity; j++ {
			left := j - fruits[i].Size
			newValue := values[left] + fruits[i].Price
			if newValue > values[j] {
				values[j] = newValue
				items[j] = i
			}
		}
	}

	for j := ksCapacity; j > 0; {
		fmt.Printf("%v, ", fruits[items[j]].Name)
		j -= fruits[items[j]].Size
	}
	fmt.Println()
	return values[ksCapacity]
}

func TestKnapsack(t *testing.T) {
	for _, v := range []int{8, 13} {
		log.Printf("+ %d: %d", v, Knapsack(v))
	}
}
