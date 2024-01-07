package lrcp

import (
	"testing"
	"time"
)

func TestCall(t *testing.T) {
	go RunServer()

	time.Sleep(150 * time.Millisecond)
	for i := 0; i < 15; i++ {
		q := i
		go Client(q)
	}

	time.Sleep(time.Second)
}
