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

// Counts the number of bits that are different in two SHA256 hashes
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))
	b := 0
	for i := 0; i < 32; i++ {
		difBits := c1[i] ^ c2[i]
		b += int(pc[difBits])
	}
	fmt.Printf("The number of different bits in 2 hashes %d\n", b)
}
