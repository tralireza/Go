package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func init() {
	log.SetFlags(0)
}

func main() {
	pid := os.Getpid()
	defer log.Printf("[%d] done.", pid)
	log.Printf("[%d] start...", pid)
	log.Printf("[%d] ppid: %d", pid, os.Getppid())

	if len(os.Args) != 1 {
		for {
			log.Printf("[%d] daemon: %v", pid, time.Now())
			time.Sleep(5 * time.Second)
		}
		return
	}

	cwd, _ := os.Getwd()
	log.Printf("[%d] -> %v", pid, filepath.Join(cwd, os.Args[0]))

	cmd := exec.Command(filepath.Join(cwd, os.Args[0]), "--daemon")
	cmd.Stdout, cmd.Stderr, cmd.Dir = os.Stdout, os.Stderr, "/"
	if err := cmd.Start(); err != nil {
		log.Printf("[%d] <FORK> %v", pid, err)
		return
	}

	p, err := os.FindProcess(pid)
	if err != nil {
		log.Printf("[%d] <RELEASE> %v", pid, err)
		return
	}
	p.Release()
}
