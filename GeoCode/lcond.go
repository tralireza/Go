package geocode

import (
	"io"
	"log"
	"sync"
)

type mWrite struct {
	sync.Mutex
	cond *sync.Cond
	bfr  []byte
	wtrs []io.Writer
}

func NewMWriter(wtrs ...io.Writer) *mWrite {
	v := mWrite{}
	v.cond = sync.NewCond(&v)
	v.wtrs = append(v.wtrs, wtrs...)
	return &v
}

func (o *mWrite) Run() {
	for _, wtr := range o.wtrs {
		go func(w io.Writer) {
			for {
				o.Lock()
				o.cond.Wait()
				n, err := w.Write(o.bfr)
				o.Unlock()
				log.Printf("W: write complete! %d %v", n, err)
			}
		}(wtr)
	}
}

func (o *mWrite) Read(data []byte) {
	o.Lock()
	o.bfr = make([]byte, len(data))
	copy(o.bfr, data)
	o.Unlock()
	o.cond.Broadcast()
	log.Print("W broadcast...")
}
