// Calculates nunmber of different bits in 2 Sha256
// hashes
package main

import (
	"crypto/sha256"
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func DiffBitsInSHA256(c1 [32]byte, c2 [32]byte) uint64 {
	if c1 == c2 {
		return 0
	} else {
		bc1, bc2 := 0, 0
		var bitsSumm uint64
		var bitDiff uint64
		for i := 0; i < len(c1); i++ {
			bc1 = PopCount(uint64(c1[i]))
			bc2 = PopCount(uint64(c2[i]))
			bitDiff = 0
			if bc1 > bc2 {
				bitDiff = uint64(bc1 - bc2)
			} else {
				bitDiff = uint64(bc2 - bc1)
			}
			bitsSumm += bitDiff
		} // for evey byte in hash
		return bitsSumm
	}
}

func main() {
	c1 := sha256.Sum256([]byte("HELLO"))
	c2 := sha256.Sum256([]byte("HELLo"))
	// 3733cd977ff8eb18b987357e22ced99f46097f31ecb239e878ae63760e83e4d5
	// 7591eb0d9700bcb8ced5f5f0053c4b3df559a19aa855f92ba054f56ea6b2ec44
	fmt.Printf("Hashes : %x\n%x\n", c1, c2)
	fmt.Printf("Set bit difference  : %d\n", DiffBitsInSHA256(c1, c2))
	// Output:
}
