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
			log.Print("-> Context done!")
			return
		case <-tkr.C:
			log.Print(time.Since(ts))
		}
	}
}
