package lrcp

import (
	"log"
	"testing"
)

func init() {
	log.Print(">>> rpc.Library")
}

func TestAdd(t *testing.T) {
	l := Library{}

	b := Book{}
	for _, v := range []string{"ISBN1", "ISBN2", "ISBN3", "ISBN2"} {
		b.ISBN = v
		log.Printf("%v -> %v", b, l.Add(b))
	}

	log.Print(l.Get("ISBN0"))
}
