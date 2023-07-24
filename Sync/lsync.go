package lsync

import (
	"log"
	"sync"
	"time"
)

func init() {
	log.SetFlags(0)
}

type LeakyBucket struct {
	capacity, status int
	mx               sync.Mutex
}

func (o *LeakyBucket) Get(quota int) int {
	defer o.mx.Unlock()
	o.mx.Lock()
	v := o.status - quota
	if v < 0 {
		quota += v
	}
	o.status -= quota
	return quota
}

func NewLeakyBucket(capacity int, rate time.Duration) *LeakyBucket {
	o := LeakyBucket{capacity: capacity, status: capacity}
	go func() {
		for {
			time.Sleep(rate)
			o.mx.Lock()
			o.status = o.capacity
			o.mx.Unlock()
		}
	}()
	return &o
}
