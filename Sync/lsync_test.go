package lsync

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
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

	var works, completes atomic.Int32
	go func() {
		for {
			w, c := works.Load(), completes.Load()
			fmt.Printf("\r%3d : %3d  (%d)", w, c, w-c)
			time.Sleep(time.Millisecond * 50)
		}
	}()

	for i := 0; i < N; i++ {
		go func() {
			tasks := rand.Intn(64)
			defer func(i int) {
				pWg.Done()
			}(tasks)
			for tasks > 0 {
				works.Add(1)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(55)))
				c <- struct{}{}
				tasks--
			}
		}()
	}

	go func() {
		pWg.Wait()
		close(c)
	}()

	for i := 0; i < M; i++ {
		go func() {
			w := 0
			defer func() {
				cWg.Done()
			}()
			for range c {
				w++
				completes.Add(1)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(75)))
			}
		}()
	}

	cWg.Wait()
	time.Sleep(time.Millisecond * 50)
	log.Printf("\n+ %v out of %v done.", completes.Load(), works.Load())
}

func TestLeakyBucket(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	o := NewLeakyBucket(ctx, 7, time.Duration(time.Second))
	for {
		q := o.Get(3)
		log.Printf("Got: %d", q)
		time.Sleep(time.Millisecond * 150)
		if q == 0 {
			log.Print("Bucket overflowing!")
			break
		}
	}

	time.Sleep(time.Second + time.Millisecond)
	log.Printf("Bucket ready! Got: %d", o.Get(7))
}
