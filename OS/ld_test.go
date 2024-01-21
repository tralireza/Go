package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"testing"
	"time"
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
		p, err := os.FindProcess(os.Getpid())
		if err != nil {
			log.Print(err)
			return
		}
		if err := p.Signal(syscall.SIGUSR2); err != nil {
			log.Print(err)
		}
	}()

	log.Printf("SIGNAL <- %v", <-c)
}

// ls -l | grep w | wc -l
func TestPipeIO(t *testing.T) {
	ws := []string{".go", "mod"}
	cmds := make([][2]*exec.Cmd, len(ws))
	wtrs := make([]io.Writer, len(ws))
	outs := make([]bytes.Buffer, len(ws))

	for i := range ws {
		var err error
		cmds[i][0] = exec.Command("grep", ws[i])
		if wtrs[i], err = cmds[i][0].StdinPipe(); err != nil {
			log.Fatal(err)
		}
		cmds[i][1] = exec.Command("wc", "-l")
		if cmds[i][1].Stdin, err = cmds[i][0].StdoutPipe(); err != nil {
			log.Fatal(err)
		}
		cmds[i][1].Stdout = &outs[i]

		log.Printf("%d| %v | %v", i, cmds[i][0], cmds[i][1])
	}

	cmd := exec.Command("ls", "-l")
	cmd.Stdout = io.MultiWriter(wtrs...)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	for i := range ws {
		if err := cmds[i][0].Start(); err != nil {
			log.Fatalf("grep Start %d: %v", i, err)
		}
		if err := cmds[i][1].Start(); err != nil {
			log.Fatalf("ws Start %d: %v", i, err)
		}
	}

	for i := range ws {
		if err := wtrs[i].(io.Closer).Close(); err != nil {
			log.Fatalf("close 0 (Stdin) %d: %v", i, err)
		}
		if err := cmds[i][0].Wait(); err != nil {
			log.Fatalf("grep Wait %d: %v", i, err)
		}
		if err := cmds[i][1].Wait(); err != nil {
			log.Fatalf("wc Wait %d: %v", i, err)
		}
		out := bytes.TrimSpace(outs[i].Bytes())
		log.Printf("%9q -> %s", cmds[i][0].Args, out)
	}
}

type logTrip struct{ http.RoundTripper }

func (t logTrip) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Print("-> ", r.URL)
	r.Header.Set("X-Log-Time", time.Now().String())
	return t.RoundTripper.RoundTrip(r)
}

func TestGLog(t *testing.T) {
	client := http.Client{Transport: logTrip{http.DefaultTransport}}
	rq, err := http.NewRequest("GET", "http://127.0.0.1:10080/search?q=Golang", nil)
	if err != nil {
		t.Fatal(err)
	}

	rsp, err := client.Do(rq)
	if err != nil {
		t.Fatal(err)
	}
	defer rsp.Body.Close()

	log.Print(rsp.StatusCode, rsp.Status)
}
