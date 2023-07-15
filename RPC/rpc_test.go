package lrcp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"strings"
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

func BenchmarkBufIO(b *testing.B) {
	b.StopTimer()
	bfr := bytes.NewBuffer(make([]byte, 128))
	bs := make([]byte, 1024)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		bfr.Read(bs)
		io.Discard.Write(bfr.Bytes())
	}
}

func BenchmarkIO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := make([]byte, 1024)
		io.Discard.Write(bs)
	}
}

func TestGenBookToFile(t *testing.T) {
	f, err := os.OpenFile("books.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fs.FileMode(0644))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	wtr := bufio.NewWriter(f)
	defer wtr.Flush()

	GenBooks(1000, wtr)
	log.Printf("+ %d", wtr.Size())
}

func TestGenBooks(t *testing.T) {
	var wtr bytes.Buffer
	GenBooks(3, &wtr)
	log.Printf("%v", wtr.String())
}

func GenBooks(n int, wtr io.Writer) {
	type Book struct {
		Author     string    `json:"author"`
		Title      string    `json:"title"`
		Year       int       `json:"year,omitempty"`
		Rating     int       `json:"rating,omitempty"`
		OnlineCopy bool      `json:"online_copy,string"`
		TStamp     time.Time `json:"timestamp"`
	}

	authors := []string{"Author 1", "Author 2", "Author 3"}
	tWords := []string{"a", "of", "the", "Games", "Pride", "Story", "Adventure", "to", "Kill", "Runaway", "Plain", "House", "Lake"}

	var bfr bytes.Buffer
	jenc := json.NewEncoder(&bfr)
	jenc.SetIndent("", "  ")

	wtr.Write([]byte{'['})
	for n > 0 {
		book := Book{Author: authors[rand.Intn(len(authors))], TStamp: time.Now()}

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

// O(n)
func PeekFinderLinear(A []int) int {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return i
		}
	}
	return len(A) - 1
}

// O(lg n)
func PeekFinder(A []int) int {
	l, r := 0, len(A)-1
	for l <= r {
		m := l + (r-l)/2
		log.Printf("(%d %d %d) -> %d %d %d", l, m, r, A[l], A[m], A[r])
		if A[m] > A[m+1] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func TestPeekFinder(t *testing.T) {
	A := []int{1, 2, 3, 4, 7, 8, 9, 3, 1}
	log.Print(A)

	i := PeekFinder(A)
	log.Printf("+ %d: %d | O(n): %d", i, A[i], PeekFinderLinear(A))
}

func TestQueryWriter(t *testing.T) {
	hlColor := 31
	rdr := strings.NewReader(`This is going to test Go/go query writer!
At least one go should trun RED :-)
No show line.`)

	io.ReadAll(io.TeeReader(rdr, NewQueryWriter(os.Stdout, "go", hlColor)))

	f, _ := os.Open("rpc_test.go")
	defer f.Close()
	io.ReadAll(io.TeeReader(bufio.NewReader(f), NewQueryWriter(os.Stdout, "go", hlColor)))
}

func NewQueryWriter(w io.Writer, query string, hlightColor int) io.Writer {
	return &QueryWriter{w, []byte(query), hlightColor}
}

type QueryWriter struct {
	io.Writer
	q      []byte
	hlCode int
}

func (q QueryWriter) Write(p []byte) (int, error) {
	for _, lb := range bytes.Split(p, []byte{'\n'}) {
		if i := bytes.Index(lb, q.q); i >= 0 {
			for _, b := range [][]byte{lb[0:i],
				[]byte(fmt.Sprintf("\x1b[%dm", q.hlCode)), q.q, []byte("\x1b[0m"),
				lb[i+len(q.q):]} {
				n, err := q.Writer.Write(b)
				if err != nil {
					return n, err
				}
			}
			q.Writer.Write([]byte{'\n'})
		}
	}
	return len(p), nil
}
