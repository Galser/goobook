package main

import (
	"fmt"
)

func reverse(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func main() {
	test0 := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(test0)
	//reverse(test0[:2])
	//reverse(test0[2:])
	reverse(&test0) // in the book this is written as "reverse(s)" on page 86, which is not going to work
	// with modern Go

	fmt.Println(test0)
	/* Output :
	[0 1 2 3 4 5]
	[2 3 4 5 0 1]
	*/

}
