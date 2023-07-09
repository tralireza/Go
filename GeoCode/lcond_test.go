package geocode

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	var wtrs []io.Writer
	for i := 0; i < 3; i++ {
		wtrs = append(wtrs, &bytes.Buffer{})
	}
	o := NewMWriter(wtrs...)
	time.Sleep(time.Second)

	for i := 0; i < 3; i++ {
		o.Read([]byte(fmt.Sprintf("Buffer data(%d)|", i)))
		time.Sleep(time.Second)
		for _, w := range wtrs {
			log.Printf("%q", w)
		}
	}
}

func TestCtxCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(time.Second, cancel)
	tkr := time.NewTicker(175 * time.Millisecond)

	ts := time.Now()
	for {
		select {
		case <-ctx.Done():
			tkr.Stop()
			log.Printf("-> Context: %v", ctx.Err())
			return
		case <-tkr.C:
			log.Print(time.Since(ts))
		}
	}
}

func TestCtxDeadline(t *testing.T) {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	lateCtx, _ := context.WithDeadline(ctx, time.Now().Add(time.Minute))

	tkr := time.NewTicker(175 * time.Millisecond)
	ts := time.Now()
	for {
		select {
		case <-lateCtx.Done():
			tkr.Stop()
			log.Printf("-> Context: %v", lateCtx.Err())
			return
		case <-tkr.C:
			log.Print(time.Since(ts))
		}
	}
}

func TestCtxTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
	time.AfterFunc(5*time.Second, cancel)

	ts := time.Now()
	for {
		select {
		case <-time.After(125 * time.Millisecond):
			log.Print(time.Since(ts))
		case <-ctx.Done():
			log.Printf("-> Context: %v (%v)", ctx.Err(), time.Since(ts).Round(time.Second))
			return
		}
	}
}

func TestCtxWithValue(t *testing.T) {
	type key struct{}
	ctx, cancel := context.WithCancel(context.Background())

	n := 3
	var wg sync.WaitGroup
	wg.Add(n)

	for n > 0 {
		go func(ctx context.Context) {
			v := ctx.Value(key{})
			log.Printf("Ctx: --Val-> %v", v)
			wg.Done()
			<-ctx.Done()
			log.Printf("Ctx: %d ->  %v", v, ctx.Err())
		}(context.WithValue(ctx, key{}, n))
		n--
	}

	wg.Wait()
	cancel()
	time.Sleep(time.Millisecond)
}

func TestHttpRqTimeout(t *testing.T) {
	go func() {
		http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) { time.Sleep(3 * time.Second) })
		if err := http.ListenAndServe(":34531", nil); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("%v", time.Duration(1))
	time.Sleep(time.Millisecond)

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	rq, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:34531/", nil)
	if err != nil {
		log.Fatal(err)
	}

	ts := time.Now()
	rsp, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Fatalf("%v (Timeout 1s) -> %v", time.Since(ts), err)
	}
	defer rsp.Body.Close()
}

// Cancels slow replica Requests
func TestMultiRqCancel(t *testing.T) {
	go func() {
		http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
			n := rand.Intn(10)
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Fprintf(w, "%d (s)", n)
		})
		http.ListenAndServe(":45123", nil)
	}()
	time.Sleep(time.Millisecond)

	ctx, cancel := context.WithCancel(context.Background())

	once, wg := sync.Once{}, sync.WaitGroup{}
	rc := make(chan *http.Response)
	wg.Add(4)
	ts := time.Now()
	for i := 0; i < 4; i++ {
		go func(i int) {
			defer wg.Done()

			rq, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:45123", nil)
			if err != nil {
				log.Printf("%v", err)
			}

			rsp, err := http.DefaultClient.Do(rq)
			if err != nil {
				log.Printf("%d. %v", i, err)
			}

			once.Do(func() { rc <- rsp })
		}(i)
	}

	rsp := <-rc
	log.Printf("%v (%v)", rsp.Status, time.Since(ts))
	rsp.Body.Close()

	cancel()
	wg.Wait()
}
