// Dup2, version for ex.1.4 prints the count and text of lines that appear more
// than once in the input.  It reads from stdin or from a list of named files.
// Ex 1.4 - print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	dupfiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin", dupfiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, dupfiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("repeats : %d\t, line : %s\n", n, line)
			fmt.Printf("  and seen in files : %s \n", strings.Join(dupfiles[line], ","))
		}
	}
}

func countLines(f *os.File, counts map[string]int, fName string, dupfiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 0 {
			dupfiles[input.Text()] = append(dupfiles[input.Text()], fName)
		}
	}
	// NOTE: again - ignoring potential errors from input.Err()
}
