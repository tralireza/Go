package geocode

import (
	"bytes"
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
	o.Read([]byte("Buffer data()"))

	time.Sleep(time.Second)
	for _, w := range wtrs {
		log.Printf("%q", w)
	}
}
