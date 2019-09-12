// E xercis e 1.1 tha output all arguments amd  also print os.Args[0],
// the name of t he command that invoked it.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Name of invoking command :", os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
