package geocode

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestBfrPool(t *testing.T) {
	n, ts := 16, time.Now()

	wg := sync.WaitGroup{}
	wg.Add(n)

	for n > 0 {
		go func(v int) {
			bfr := GetBfr()
			defer func() {
				PutBfr(bfr)
				wg.Done()
			}()

			fmt.Fprintf(bfr, "%2d -> %p   %v", v, bfr, time.Since(ts))
			log.Printf("%s", bfr.Bytes())
		}(n)
		n--
	}

	wg.Wait()
}
