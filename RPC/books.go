package lrcp

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

func init() {
	log.SetFlags(0)
}

func NewISBN() ISBN {
	B := []byte{}
	for i := 0; i < 10; i++ {
		B = append(B, '0'+byte(rand.Intn(9)))
	}
	rand.Shuffle(len(B), func(i, j int) { B[i], B[j] = B[j], B[i] })
	switch v := rand.Intn(7); v {
	case 0:
		B[len(B)-1] = 'X'
	}
	return ISBN(string(B))
}

type ISBN string
type Book struct {
	ISBN       ISBN
	Author     string
	Title      string
	Year       int
	OnlineCopy bool
	Pages      int
	Timestamp  time.Time
}

var (
	ErrDuplicate = fmt.Errorf("duplicate ISBN")
	ErrMissing   = fmt.Errorf("book not found")
)

type Library struct{ B []*Book }

func (l *Library) Add(b Book) error {
	for _, v := range l.B {
		if v.ISBN == b.ISBN {
			return ErrDuplicate
		}
	}
	l.B = append(l.B, &b)
	return nil
}

func (l *Library) Get(isbn ISBN) (Book, error) {
	for _, v := range l.B {
		if v.ISBN == isbn {
			return *v, nil
		}
	}
	return Book{}, ErrMissing
}

type LibrarySvc struct{ Library }

func (o *LibrarySvc) Add(b Book, index *int) error {
	if err := o.Library.Add(b); err != nil {
		return err
	}
	*index = len(o.Library.B)
	return nil
}

func RunRPCServer() {
	if err := rpc.Register(&LibrarySvc{}); err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()

	lsr, err := net.Listen("tcp", ":18080")
	if err != nil {
		log.Fatal(err)
	}
	if err := http.Serve(lsr, nil); err != nil {
		log.Fatal(err)
	}
}
