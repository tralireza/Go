package lrcp

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	log.SetFlags(0)
}

func GetRandomISBN() string {
	B := []byte{}
	for i := 0; i < 10; i++ {
		B = append(B, '0'+byte(rand.Intn(9)))
	}
	rand.Shuffle(len(B), func(i, j int) { B[i], B[j] = B[j], B[i] })
	switch v := rand.Intn(7); v {
	case 0:
		B[len(B)-1] = 'X'
	}
	return string(B)
}

type Book struct {
	ISBN       string
	Author     string
	Title      string
	Year       int
	OnlineCopy bool
	Pages      int
	Timestamp  time.Time
}

var (
	ErrDuplicate = fmt.Errorf("duplicate ISBN")
	ErrMissing   = fmt.Errorf("ISBN not found")
)

type Library struct {
	B []*Book
}

func (l *Library) Add(b Book) error {
	for _, v := range l.B {
		if v.ISBN == b.ISBN {
			return ErrDuplicate
		}
	}
	l.B = append(l.B, &b)
	return nil
}

func (l *Library) Get(isbn string) (Book, error) {
	for _, v := range l.B {
		if v.ISBN == isbn {
			return *v, nil
		}
	}

	return Book{}, ErrMissing
}
