package geocode

import (
	"log"
	"net"
	"os"
	"sync"
	"testing"
	"time"
)

func TestSync(t *testing.T) {
	m := sync.Mutex{}
	c := make(chan struct{}, 1000)
	a := 0
	ts := time.Now()
	for i := 0; i < cap(c); i++ {
		go func(i int) {
			defer func() { c <- struct{}{} }()
			m.Lock()
			if i&1 == 0 {
				a++
			} else {
				a--
			}
			m.Unlock()
		}(i)
	}
	for i := 0; i < cap(c); i++ {
		<-c
	}
	duration := time.Since(ts)
	log.Printf("[%d]   %v", a, duration)
}

func TestNets(t *testing.T) {
	I := []net.IP{}

	ifcs, err := net.Interfaces()
	if err != nil {
		t.Fatal(err)
	}
	for _, ifc := range ifcs {
		addrs, err := ifc.Addrs()
		if err != nil {
			t.Fatal(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip = ip.To4(); ip != nil {
				I = append(I, ip)
			}
		}
	}

	log.Printf("+ IPs: %+v", I)
}

func TestReverseGeocode(t *testing.T) {
	c := NewClient(time.Second, os.Getenv("GCMAP_API"))

	ts := time.Now()
	ls, err := c.ReverseGeocode(51.44, -0.34)
	if err != nil {
		log.Fatal(err)
	}
	et := time.Since(ts)

	for i, l := range ls {
		log.Printf("%d -> %T", i, l)
		for i, v := range l.AddressComponents {
			log.Printf("  %d -> %T", i, v)
		}
	}
	log.Print(et)
}
