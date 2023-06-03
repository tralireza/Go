package geocode

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestReverseGeocode(t *testing.T) {
	c := NewClient(time.Second, os.Getenv("GCMAP_API"))

	ts := time.Now()
	ls, err := c.ReverseGeocode(51.44, -0.34)
	if err != nil {
		log.Fatal(err)
	}
	et := time.Since(ts)

	for i, l := range ls {
		log.Printf("%d -> %T", i, l)
		for i, v := range l.AddressComponents {
			log.Printf("  %d -> %T", i, v)
		}
	}
	log.Print(et)
}
