// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Fetchall fetches URLs in parallel and reports their times and sizes.
// Modification n fro exercise 1.19
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	f, err := os.Create("fetch_all_output.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	//var s string
	for range os.Args[1:] {
		s := <-ch // receive from channel ch
		fmt.Println(s)
		fmt.Fprintln(f, s)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
