package main

import (
	"fmt"
	"mycode/ch2/lenconvert"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uc: %v\n", err)
			os.Exit(1)
		}
		f := lenconvert.Feet(t)
		m := lenconvert.Meters(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, lenconvert.FTToM(f), m, lenconvert.MToFT(m))
	}
}
