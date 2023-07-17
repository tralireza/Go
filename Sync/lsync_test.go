package lsync

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {
	Squarer := func(inc <-chan int) <-chan int {
		c := make(chan int)
		go func() {
			defer close(c)
			for i := range inc {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(175)))
				c <- i * i
			}
		}()
		return c
	}

	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < 32; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(75)))
			c <- i
		}
	}()

	FanIn := func(cs ...<-chan int) <-chan int {
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

	for i := range FanIn(Squarer(c), Squarer(c)) {
		fmt.Print(i, ",")
	}
	fmt.Println()
}
