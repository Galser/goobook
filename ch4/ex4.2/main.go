// Prints hash for standart input, defaul is 256
// but you an specify two flags "sha384" and "sha512" to
// print extended hashes respectfully
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var showSHA384 = flag.Bool("sha384", false, "Generate SHA384 hash")
var showSHA512 = flag.Bool("sha512", false, "Generate SHA512 hash")

func main() {
	var printed = false
	var inputBytes []byte
	input := bufio.NewScanner(os.Stdin)
	inputBytes = input.Bytes()

	flag.Parse()
	if *showSHA384 {
		hash := sha512.Sum384(inputBytes)
		printed = true
		fmt.Printf("SHA384 hash for input is : %x\n", hash)
	}
	if *showSHA512 {
		hash := sha512.Sum512(inputBytes)
		printed = true
		fmt.Printf("SHA512 hash for input is : %x\n", hash)
	}
	// and print SHA256 by default
	if !printed {
		hash := sha256.Sum256(inputBytes)
		fmt.Printf("Default SHA256 hash for input is : %x\n", hash)
	}
	/* Output for "test":
	➜ ex4.2 git:(ex4-2) ✗ go run main.go < input.txt
	Default SHA256 hash for input is : e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
	➜ ex4.2 git:(ex4-2) ✗ go run main.go -sha384 < input.txt
	SHA384 hash for input is : 38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b
	➜ ex4.2 git:(ex4-2) ✗ go run main.go -sha512 < input.txt
	SHA512 hash for input is : cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e
	*/
}
