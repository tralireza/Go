package lc

import (
	"log"
	"testing"
)

func init() {
	log.Print("> Basic DS")
}

// 1614 Maximum Nesting Depth of Parentheses
func Test1614(t *testing.T) {
	maxDepth := func(s string) int {
		S := []byte{}
		x := 0
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '(':
				S = append(S, '(')
				x = max(x, len(S))
			case ')':
				S = S[:len(S)-1]
			}
		}
		return x
	}

	log.Print("3 ?= ", maxDepth("(1+(2*3)+((8)/4))+1"))
	log.Print("3 ?= ", maxDepth("(1)+((2))+(((3)))"))
}

// 1544 Make String Great
func Test1544(t *testing.T) {
	makeGood := func(s string) string {
		i := 0
		for i < len(s)-1 {
			log.Printf("%q %3[1]d %q %3[2]d | %3d", s[i], s[i+1], s[i]-s[i+1])

			if s[i]-s[i+1] == 'a'-'A' || s[i]-s[i+1] == 256-('a'-'A') {
				s = s[:i] + s[i+2:]
				i = 0
			} else {
				i++
			}
		}
		return s
	}

	fmt.Print("byte 'Overflow' -> ")
	b := byte(253)
	for range 7 {
		fmt.Printf("%d,", b)
		b++
	}
	fmt.Print("\n")

	log.Print(" ?= ", makeGood("abBACc"))
	log.Print(" ?= ", makeGood("Pp"))
}
