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

func Worker(inc <-chan int) <-chan int {
	worker := func(i int) int { return i * i }

	c := make(chan int)
	go func() {
		defer close(c)
		for i := range inc {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(250)))
			c <- worker(i)
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
		for i := 1; i <= 32; i++ {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(75)))
			c <- i
		}
	}()

	fanOut := func(factor int, inc <-chan int) (cs []<-chan int) {
		for factor > 0 {
			cs = append(cs, Worker(inc))
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

// N*M Producers & Consumers
func TestProdCons(t *testing.T) {
	c := make(chan struct{})

	N, M := 7, 9
	pWg, cWg := sync.WaitGroup{}, sync.WaitGroup{}
	pWg.Add(N)
	cWg.Add(M)

	for i := 0; i < N; i++ {
		go func() {
			tasks := rand.Intn(32)
			defer func(i int) {
				pWg.Done()
				log.Printf("P: %2d tasks", i)
			}(tasks)
			for tasks > 0 {
				c <- struct{}{}
				tasks--
			}
		}()
	}

	go func() {
		pWg.Wait()
		log.Print("No more Producers: closing chan.")
		close(c)
	}()

	for i := 0; i < M; i++ {
		go func() {
			w := 0
			defer func() {
				cWg.Done()
				log.Printf("C: %2d tasks complete", w)
			}()
			for range c {
				w++
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(75)))
			}
		}()
	}

	cWg.Wait()
	log.Print("All done!")
}
