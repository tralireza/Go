package lsync

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Squarer(inc <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := range inc {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(250)))
			c <- i * i
		}
	}()
	return c
}

// Demux from channels into one
func FanIn(cs ...<-chan int) <-chan int {
	outc := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			defer wg.Done()
			for i := range c {
				outc <- i
			}
		}(c)
	}
	go func() {
		defer close(outc)
		wg.Wait()
	}()
	return outc
}

func TestFanIn(t *testing.T) {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < 32; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(75)))
			c <- i
		}
	}()

	fanOut := func(factor int, inc <-chan int) (cs []<-chan int) {
		for factor > 0 {
			cs = append(cs, Squarer(inc))
			factor--
		}
		return cs
	}

	ts := time.Now()
	for i := range FanIn(fanOut(5, c)...) {
		fmt.Print(i, ",")
	}
	log.Printf("\nExec Time: %v", time.Since(ts))
}
