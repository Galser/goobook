// Prints frequncies of each word in an input text file
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanWord(s string) string {
	news := strings.ReplaceAll(s, ",", "")
	news = strings.ReplaceAll(news, ".", "")
	news = strings.ReplaceAll(news, ":", "")
	return strings.ToLower(news)
}

func main() {
	freq := make(map[string]int64) // a set of strings (words) with frequencies

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while trying to access input file: %v\n", err)
			continue
		}
		input := bufio.NewScanner(file)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			//			line := input.Text()
			/*			if _, ok = freq[input.Text()] ; !ok {
						freq[input.Text()] = 0
					} */
			freq[cleanWord(input.Text())]++
		}
	}
	for word, freq := range freq {
		fmt.Println(word, ":", freq)
	}
	/* output example
		[]type : 1
	let’s : 1
	while : 1
	that : 3
	it : 1
	by : 1
	 : 2
	the : 4
	arg2 : 1
	be : 3
	of : 8
	looks : 1
	elem : 3
	an : 2
	argsn) : 1
	can : 1
	s2 : 1
	function : 5
	to : 7
	type : 7
	append : 5
	argument : 2
	or : 1
	instead : 1
	variadic : 1
	this : 2
	from : 1
	above : 1
	with : 1
	we : 4
	want : 1
	(which : 1
	..
	*/
}
