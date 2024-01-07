package lrcp

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

func init() {
	log.SetFlags(0)
	log.Print("lrcp >")
}

type Server struct{}

func (o *Server) Reply(rq int, rp *time.Time) error {
	log.Printf("RCP rq: %v", rq)
	*rp = time.Now()
	return nil
}

func Client() {
	cnn, err := rpc.Dial("tcp", ":19999")
	if err != nil {
		log.Print(err)
		return
	}

	var tsServer time.Time
	if err := cnn.Call("Server.Reply", 1, &tsServer); err != nil {
		log.Print(err)
		return
	}

	log.Printf("TimeStamp on Server (via RPC) -> %v", tsServer)
}

func RunServer() {
	rpc.Register(&Server{})

	lsr, err := net.Listen("tcp", ":19999")
	if err != nil {
		log.Print(err)
		return
	}

	log.Print("...")
	for {
		cnn, err := lsr.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go rpc.ServeConn(cnn)
	}
}
