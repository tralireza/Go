package geocode

import (
	"bytes"
	"fmt"
	"sync"
)

type bfrPool struct {
	sync.Pool
	sync.Map
}

func NewBfrPool() *bfrPool {
	return &bfrPool{
		sync.Pool{New: func() interface{} { return bytes.NewBuffer(make([]byte, 0, 4096)) }},
		sync.Map{},
	}
}

func (o *bfrPool) Get() *bytes.Buffer {
	v := o.Pool.Get().(*bytes.Buffer)
	o.Map.LoadOrStore(v, struct{}{})
	return v
}

func (o *bfrPool) Put(bfr *bytes.Buffer) {
	bfr.Reset()
	o.Pool.Put(bfr)
}

func (o *bfrPool) Len() int {
	n := 0
	o.Map.Range(func(k, v interface{}) bool {
		fmt.Printf("| %p %v ", k.(*bytes.Buffer), v)
		n++
		return true
	})
	fmt.Println("|")
	return n
}
