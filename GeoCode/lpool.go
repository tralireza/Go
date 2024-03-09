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

func GetBfr() *bytes.Buffer {
	return BfrPool.Get().(*bytes.Buffer)
}

func PutBfr(bfr *bytes.Buffer) {
	bfr.Reset()
	BfrPool.Put(bfr)
}
