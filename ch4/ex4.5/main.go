// Function that eliminates adjacent elements
// in-place from []string
package main

import "fmt"

func RemoveDups(strings []string) []string {
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] == strings[i+1] {
			copy(strings[i:], strings[i+1:])
			strings = strings[:len(strings)-1]
		}
	} // strings iteration
	return strings
}

func main() {
	s := []string{"one", "two", "two", ".", ".", "three", "four", "four", "three"}
	fmt.Printf("String before cleaning : %v \n", s)
	s = RemoveDups(s)
	fmt.Printf("String after cleaning : %v \n", s)
	/* output :
	String before cleaning : [one two two . . three four four three]
	String after cleaning : [one two . three four three]
	*/

}
