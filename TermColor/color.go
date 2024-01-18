package termcolor

import (
	"fmt"
	"io"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
}

type TermColor int

func (c TermColor) Out(v any) { c.Fprintf(os.Stdout, "%v", v) }
func (c TermColor) Fprintf(w io.Writer, format string, args ...any) {
	fmt.Fprintf(w, "\x1b[%dm", c)
	fmt.Fprintf(w, format, args...)
	fmt.Fprintf(w, "\x1b[%dm", 0)
}

const (
	Red     TermColor = 31
	Green   TermColor = 32
	Yellow  TermColor = 33
	Blue    TermColor = 34
	Magenta TermColor = 35
	Cyan    TermColor = 36
	White   TermColor = 37
)
