package main

import "fmt"

func mmin(inputs ...int) int {
	if len(inputs) > 0 {
		minv := inputs[0]
		fmt.Println("Values :", inputs)
		if len(inputs) > 1 {
			for _, val := range inputs {
				if val < minv {
					minv = val
				}
			}
		} // if
		return minv
	}
	return 0
}

func mmax(inputs ...int) int {
	if len(inputs) > 0 {
		maxv := inputs[0]
		fmt.Println("Values :", inputs)
		if len(inputs) > 1 {
			for _, val := range inputs {
				if val > maxv {
					maxv = val
				}
			}
		} // if
		return maxv
	}
	return 0
}

func main() {
	fmt.Println("test")
	fmt.Println("Minimum ", mmin(10, 3, 5, 6))
	fmt.Println("Maximum ", mmax(10, 3, 5, 6))
}

/* Output

go run ch5/minmax/main.go
test
Values : [10 3 5 6]
Minimum  3
Values : [10 3 5 6]
Maximum  10

*/
