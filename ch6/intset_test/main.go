package main

import (
	"fmt"
	"mycode/ch6/intset"
)

func main() {
	var x, y intset.IntSet
	var z *intset.IntSet
	x.Add(1)
	fmt.Println(x.String())      // "{1 9 144}"
	fmt.Println("len:", x.Len()) //
	x.Add(144)
	fmt.Println(x.String())      // "{1 9 144}"
	fmt.Println("len:", x.Len()) //
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())    //

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())    //

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Len())    //

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	z = x.Copy()
	fmt.Println(z.String())
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}
