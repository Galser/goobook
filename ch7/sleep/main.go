package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

/* Output :
 ~/gobook/src/mycode   flag-value  ./gsleep
Sleeping for 1s...
 ~/gobook/src/mycode   flag-value  ./gsleep -period 5s
Sleeping for 5s...

*/
