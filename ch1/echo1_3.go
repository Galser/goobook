// Echo2 prints its comand-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// variant 1
	fmt.Println("Starting variant 1 (for loop)")
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("End of run for variant 1 : %dMs elapsed \n", time.Since(start).Microseconds())
	// variant 2
	fmt.Println("Starting variant 2 (Stings.join)")
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("End of run for variant 2 : %dMs elapsed \n", time.Since(start).Microseconds())
}
