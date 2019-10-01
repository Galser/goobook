// Fetch prints the content found at each specified URL.
// exercise 1.9 with HTTP status code
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "HTTP Status code : %d and status : %s \n", resp.StatusCode, resp.Status)
		//b, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: Error copying %s: %v\n", url, err)
			os.Exit(1)
		}
		//		fmt.Printf("%s", b)
	}
}