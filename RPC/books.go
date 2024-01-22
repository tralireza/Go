package lrcp

import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetFlags(0)
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
