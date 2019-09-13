// Echo2 prints its comand-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// variant 1
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	// variant 2
	fmt.Println(strings.Join(os.Args[1:], " "))
}
