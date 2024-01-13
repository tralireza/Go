package lrcp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestIOBuf(t *testing.T) {
	f, err := os.Open("rpc.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rdr := bufio.NewReader(f)
	for {
		if l, err := rdr.ReadString('\n'); err == nil {
			log.Printf("| %d", len(l))
		} else {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func TestIO(t *testing.T) {
	f, err := os.Open("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bs := make([]byte, 16)
	bfs := bytes.Buffer{}
	for {
		if n, err := f.Read(bs); err == nil {
			log.Printf("> %d\n| %q\n| %[2]v\n| %[2]x", n, bs[:n])
			bfs.Write(bs[:n])
			log.Printf("+ %d %d", bfs.Len(), bfs.Cap())
			bfs.Reset()
		} else {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func TestCall(t *testing.T) {
	go RunServer()
	time.Sleep(150 * time.Millisecond)

	errc := make(chan error)
	for i := 0; i < 15; i++ {
		q := i
		go func() { errc <- Client(q) }()
	}
	for i := 0; i < 15; i++ {
		<-errc
	}
}

func Knapsack(ksCapacity int) int {
	type Fruit struct {
		Size, Price int
		Name        string
	}

	fruits := []Fruit{{4, 4500, "Plum"},
		{5, 5700, "Apple"},
		{2, 2250, "Orange"},
		{1, 1100, "Strawberry"},
		{6, 6700, "Melon"}}

	items := make([]int, ksCapacity+1)
	values := make([]int, ksCapacity+1)

	for i := 0; i < len(fruits); i++ {
		for j := fruits[i].Size; j <= ksCapacity; j++ {
			left := j - fruits[i].Size
			newValue := values[left] + fruits[i].Price
			if newValue > values[j] {
				values[j] = newValue
				items[j] = i
			}
		}
	}

	for j := ksCapacity; j > 0; {
		fmt.Printf("%v, ", fruits[items[j]].Name)
		j -= fruits[items[j]].Size
	}
	fmt.Println()
	return values[ksCapacity]
}

func TestKnapsack(t *testing.T) {
	for _, v := range []int{8, 9, 18} {
		log.Printf("+ %d: %d", v, Knapsack(v))
	}
}

type NullWriter struct{}

func (NullWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func BenchmarkBufIO(b *testing.B) {
	b.StopTimer()
	bfr := bytes.NewBuffer(make([]byte, 128))
	bs := make([]byte, 1024)

	b.StartTimer()
	var wtr NullWriter
	for i := 0; i < b.N; i++ {
		bfr.Read(bs)
		bfr.WriteTo(wtr)
	}
}

func BenchmarkIO(b *testing.B) {
	var wtr NullWriter
	for i := 0; i < b.N; i++ {
		bs := make([]byte, 1024)
		wtr.Write(bs)
	}
}

func TestGenBooks(t *testing.T) {
	var wtr bytes.Buffer
	GenBooks(3, &wtr)
	log.Printf("%v", wtr.String())
}

func GenBooks(n int, wtr io.Writer) {
	type Book struct {
		Author     string `json:"author"`
		Title      string `json:"title"`
		Year       int    `json:"year,omitempty"`
		Rating     int    `json:"rating,omitempty"`
		OnlineCopy bool   `json:"online_copy,string"`
	}

	authors := []string{"Author 1", "Author 2", "Author 3"}
	tWords := []string{"a", "of", "the", "Games", "Pride", "Story", "Adventure", "to", "Kill", "Runaway", "Plain", "House", "Lake"}

	var bfr bytes.Buffer
	jenc := json.NewEncoder(&bfr)
	jenc.SetIndent("", "  ")

	wtr.Write([]byte{'['})
	for n > 0 {
		book := Book{Author: authors[rand.Intn(len(authors))]}

		k := 3 + rand.Intn(len(tWords)-3)
		for _, i := range rand.Perm(len(tWords)) {
			book.Title += tWords[i]
			if k > 0 {
				book.Title += " "
			}
			if k == 0 {
				break
			}
			k--
		}

		switch rand.Intn(2) {
		case 0:
			book.Year = 1900 + rand.Intn(124)
		}
		switch rand.Intn(2) {
		case 0:
			book.OnlineCopy = true
		}
		switch r := rand.Intn(6); r {
		case 0:
		default:
			book.Rating = r
		}

		jenc.Encode(book)
		if n > 1 {
			bfr.WriteByte(',')
		}
		bfr.WriteTo(wtr)

		n--
	}
	wtr.Write([]byte{']'})

}
