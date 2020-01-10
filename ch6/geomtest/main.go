package main

import (
	"fmt"
	"mycode/ch6/geometry"
)

func main() {
	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(geometry.Path.Distance(perim)) // "12", standalone function
	fmt.Println(perim.Distance())              // "12", method of geometry.Path
}

// There is an error in the book, page 157, end of the page
// missing point so it looks like this :
// fmt.Println(geometry.PathDistance(perim)) // "12", standalone function
