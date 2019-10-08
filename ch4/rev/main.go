package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	test0 := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(test0)
	reverse(test0[:])
	fmt.Println(test0)
	/* Output : 
	[0 1 2 3 4 5]
	[5 4 3 2 1 0]
	*/

}

