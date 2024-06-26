package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Content-Type: text/html\n\n")
	fmt.Printf("<h4>%s</h4>\n", os.Getenv("REMOTE_ADDR"))

	fmt.Println("<pre>\n")
	for i, s := range os.Environ() {
		v := strings.SplitN(s, "=", 2)
		fmt.Printf("%2d %s %s\n", i+1, s, v)
	}
	fmt.Println("</pre>")
}
