package geocode

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
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
