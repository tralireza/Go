package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"testing"
)

func TestStart(t *testing.T) {
	log.Print(os.Args[1:])

	p, _ := filepath.Abs(os.Args[0])
	log.Print(p)
}

func TestSignal(t *testing.T) {
	log.Print("...")

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT,
		syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		p, _ := os.FindProcess(os.Getpid())
		if err := p.Signal(syscall.SIGUSR2); err != nil {
			log.Print(err)
		}
	}()

	log.Printf("SIGNAL <- %v", <-c)
}
