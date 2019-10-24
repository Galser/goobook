package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"mycode/ch4/treesort"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func main() {
	data := make([]int, 20)
	for i := range data {
		data[i] = rand.Int() % 20
		fmt.Print(" ", data[i])
	}
	fmt.Println("\n Tree is gogin to be next")
	treesort.Sort(data)
	fmt.Println(data)

/* 
what this code is not doing - deduplication : 
Output : 
 10 11 1 11 17 0 18 8 16 9 4 7 14 16 15 13 8 11 10 11
 Tree is gogin to be next
  11     1  
  11    8     0    10  
  17    9     4  
  18    7     8    16  
  16     14  
  15     13  
   11  
  11   
[0 1 4 7 8 8 9 10 10 11 11 11 11 13 14 15 16 16 17 18]
*/
}
