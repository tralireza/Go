package lrcp

import (
	"testing"
	"time"
)

func TestCall(t *testing.T) {
	go RunServer()

	time.Sleep(150 * time.Millisecond)
	go Client()

	time.Sleep(time.Second)
}
