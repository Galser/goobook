package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(" -7 --> ", Signum(-7))
}

/*
A switch do es not ne e d an op erand; it c an just list t he c as es, e ach of w hich is a b o ole an expression:
This for m is c al le d a tagless sw itch; it’s e quivalent to switch true.
*/
func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}

