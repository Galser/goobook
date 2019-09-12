// Exercise 1.2 Modification of the echo programm taht prints
// index and value of each arguments one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Index of the argument : ", i, ", value of the argument : '", os.Args[i], "'")
	}
}
