package lrcp

import (
	"bytes"
	"encoding/gob"
	"encoding/xml"
	"log"
	"net/rpc"
	"strings"
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

	b := Book{ISBN: NewISBN(), Title: "Title1", Author: "Author1"}
	for i := 0; i < 2; i++ {
		index := -1
		if err := client.Call("LibrarySvc.Add", b, &index); err != nil {
			log.Printf("%d. error: %v", i, err)
			continue
		}
		log.Printf("%d. Book: %d", i, index)
	}

	count := -1
	if err := client.Call("LibrarySvc.Count", 0, &count); err != nil {
		log.Fatal(err)
	}
	log.Printf("Books: %d", count)
}

type lS struct {
	XMLName   struct{} `xml:"treasure"`
	ID        iD       `xml:",attr"`
	Name      string   `xml:"name"`
	Job       string   `xml:"details>job,omitempty"`
	BirthYear int      `xml:"birth_year,omitempty"`
}
type iD string

func (o iD) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  xml.Name{Local: "id"},
		Value: strings.ToUpper(string(o)),
	}, nil
}

func TestXML(t *testing.T) {
	c := lS{Name: "Hakim-e-Tos", Job: "Poet", BirthYear: 550, ID: "1"}

	bfr := bytes.Buffer{}
	e := xml.NewEncoder(&bfr)
	e.Indent("", "  ")
	if err := e.Encode(c); err != nil {
		t.Fatal(err)
	}
	log.Printf("%d:%d -> %v", bfr.Len(), len(bfr.String()), bfr.String())
}

func TestGOB(t *testing.T) {
	luke := lS{Name: "Luke", Job: "Jedi"}

	var s strings.Builder
	e := gob.NewEncoder(&s)
	if err := e.Encode(luke); err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v\n|--->>>\n%s\n<<<---|\n", s, s.String())
}
