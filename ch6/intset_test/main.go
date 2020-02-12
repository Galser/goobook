package main

import (
	"fmt"
	"mycode/ch6/intset"
)

func main() {
	var x, y intset.IntSet
	var z *intset.IntSet
	x.Add(1)
	fmt.Println(x.String())         // "{1 9 144}"
	fmt.Println("x len: ", x.Len()) //
	x.Add(144)
	fmt.Println(x.String())            // "{1 9 144}"
	fmt.Println("x len now:", x.Len()) //
	x.Add(9)
	fmt.Println(x.String())                 // "{1 9 144}"
	fmt.Println("and x len now :", x.Len()) //

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())         // "{9 42}"
	fmt.Println("y len :", y.Len()) //

	x.UnionWith(&y)
	fmt.Println(x.String())                     // "{1 9 42 144}"
	fmt.Println("x len after union :", x.Len()) //

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	z = x.Copy()
	fmt.Println("COPY of x ", z.String())
	z.Clear()
	fmt.Println("z after clear :", z.String())

	x.Remove(42)
	fmt.Println("x after Remove(42) : ", x.String())
	//!-main

	// Output:
	// {1}
	// x len:  1
	// {1 144}
	// x len now: 2
	// {1 9 144}
	// and x len now : 3
	// {9 42}
	// y len : 2
	// {1 9 42 144}
	// x len after union : 4
	// true false
	// COPY of x  {1 9 42 144}
	// z after clear : {}
	// x after Remove(42) :  {1 9 144}
	//
}
