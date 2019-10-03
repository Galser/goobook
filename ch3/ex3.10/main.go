package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	for i := 0; i <= len(s)-1; i++ {
		buf.WriteByte(s[i])
		if (i+1)%3 == 0 {
			buf.WriteString(", ")
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("6523525234235"))
	//go run ex3.10/main.go
	//652, 352, 523, 423, 5
}
