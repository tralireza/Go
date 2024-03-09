package geocode

import (
	"bytes"
	"sync"
)

var BfrPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 4096))
	},
}

func Get() *bytes.Buffer {
	return BfrPool.Get().(*bytes.Buffer)
}

func Put(bfr *bytes.Buffer) {
	bfr.Reset()
	BfrPool.Put(bfr)
}
