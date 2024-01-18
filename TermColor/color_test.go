package termcolor

import (
	"log"
	"testing"
)

func init() {
	log.Print("TermColor >")
}

func TestColor(t *testing.T) {
	Red.Out("*** Red ***\n")
	Green.Out("*** Green ***\n")
	Yellow.Out("*** Yellow ***\n")
}
