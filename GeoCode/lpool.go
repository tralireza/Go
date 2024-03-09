package geocode

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

type bfrPool struct {
	sync.Pool
	sync.Map
	rpTime time.Duration
	calls  int
}

func NewBfrPool() *bfrPool {
	return &bfrPool{
		sync.Pool{
			New: func() interface{} { return bytes.NewBuffer(make([]byte, 0, 4096)) },
		},
		sync.Map{},
		0,
		0,
	}
}

func (o *bfrPool) Get() *bytes.Buffer {
	o.calls++
	ts := time.Now()
	v := o.Pool.Get().(*bytes.Buffer)
	o.rpTime = (time.Duration(o.calls-1)*o.rpTime + time.Since(ts)) / time.Duration(o.calls)

	o.Map.LoadOrStore(v, struct{}{})
	return v
}

func (o *bfrPool) Put(bfr *bytes.Buffer) {
	bfr.Reset()
	o.Pool.Put(bfr)
}

func (o *bfrPool) Stat() (size int, tcalls int, rpTime time.Duration) {
	o.Map.Range(func(k, v interface{}) bool {
		fmt.Printf("|%p%v", k, v)
		size++
		return true
	})
	fmt.Println("|")
	return size, o.calls, o.rpTime
}
