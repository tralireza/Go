package lrcp

import (
	"log"
	"net/rpc"
	"testing"
	"time"
)

func init() {
	log.Print("> rpc.Library")
}

func TestISBN(t *testing.T) {
	for i := 0; i < 15; i++ {
		log.Print(NewISBN())
	}
}

func TestAdd(t *testing.T) {
	l := Library{}

	b := Book{}
	for i := 0; i < 5; i++ {
		b.ISBN = NewISBN()
		log.Printf("%v -> %v", b, l.Add(b))
	}
	b.Title = "New Title"
	log.Printf("%v -> %v", b, l.Add(b))

	if _, err := l.Get("ISBN0"); err == nil {
		t.Fatalf("Wrong error: must be %v != %v", ErrMissing, err)
	}
}

func TestServerRPC(t *testing.T) {
	go RunRPCServer()
	time.Sleep(150 * time.Millisecond)

	client, err := rpc.DialHTTP("tcp", ":18080")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	index := -1
	b := Book{ISBN: NewISBN(), Title: "Title1", Author: "Author1"}
	if err := client.Call("LibrarySvc.Add", b, &index); err != nil {
		log.Fatal(err)
	}
	log.Printf("Book: %d", index)
}
