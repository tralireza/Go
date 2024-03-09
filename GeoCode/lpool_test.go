package geocode

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestBfrPool(t *testing.T) {
	bfrPl := NewBfrPool()
	wg := sync.WaitGroup{}

	n := 16
	wg.Add(n)
	ts := time.Now()
	for n > 0 {
		go func(v int) {
			bfr := bfrPl.Get()
			defer func() {
				bfrPl.Put(bfr)
				wg.Done()
			}()

			fmt.Fprintf(bfr, "%2d -> %p   %v", v, bfr, time.Since(ts))
			log.Printf("%s", bfr.Bytes())
		}(n)
		n--
	}

	wg.Wait()
	size, calls, rpTime := bfrPl.Stat()
	log.Printf("Stat (size:calls): %d:%d,   Eff: %%%d,   Get Rsp. (avg): %v", size, calls, 100*(calls-size)/calls, rpTime)
}
