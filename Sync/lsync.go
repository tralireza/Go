package lsync

import (
	"log"
	"sync/atomic"
	"time"
)

func init() {
	log.SetFlags(0)
}

type LeakyBucket struct {
	capacity, status uint32
}

func (o *LeakyBucket) Get(quota uint32) uint32 {
	for {
		v := atomic.LoadUint32(&o.status)
		if v == 0 {
			return 0
		}

		if quota > v {
			quota = v
		}
		if atomic.CompareAndSwapUint32(&o.status, v, v-quota) {
			return quota
		}
	}
}

func NewLeakyBucket(capacity uint32, rate time.Duration) *LeakyBucket {
	o := LeakyBucket{capacity: capacity, status: capacity}
	tkr := time.NewTicker(rate)
	go func() {
		for {
			for range tkr.C {
				atomic.StoreUint32(&o.status, o.capacity)
			}
		}
	}()
	return &o
}
