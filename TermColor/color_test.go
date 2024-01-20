package termcolor

import (
	"bytes"
	"log"
	"math/rand"
	"testing"
)

func init() {
	log.Print("TermColor >")
}

func TestOut(t *testing.T) {
	var bfr bytes.Buffer
	Red.Fprintf(&bfr, "%d", Red)
	if !bytes.Equal(bfr.Bytes(), []byte{27, 91, 51, 49, 109, 51, 49, 27, 91, 48, 109}) {
		t.Fatalf(`Wrong escape! "\x1b[31m31\x1b[0m" != %q`, bfr.String())
	}
}

func TestColor(t *testing.T) {
	Red.Out("--- Red ---\n")
	Green.Out("--- Green ---\n")
	Yellow.Out("--- Yellow ---\n")

	TermColor(31 + rand.Intn(8)).Out("[ Random Color ]\n")
}
