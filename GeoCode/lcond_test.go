package geocode

import (
	"bytes"
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
