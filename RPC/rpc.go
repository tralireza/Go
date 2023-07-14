package lrcp

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

func init() {
	log.SetFlags(0)
	log.Print("> rpc")
}

type Server struct{}

func (o *Server) Reply(rq int, rp *time.Time) error {
	log.Printf("RCP rq: %v", rq)
	*rp = time.Now()
	return nil
}

func Client(q int) error {
	cnn, err := rpc.Dial("tcp", ":19999")

	if err != nil {
		log.Print(err)
		return err
	}

	var tsServer time.Time
	if err := cnn.Call("Server.Reply", q, &tsServer); err != nil {
		log.Print(err)
		return err
	}

	log.Printf("%2d. TimeStamp on Server (via RPC) -> %v", q, tsServer)
	return err
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
