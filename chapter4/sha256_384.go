package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

// Counts the number of bits that are different in two SHA256 hashes
func main() {
	var n = flag.String("hash", "SHA256", "SHA256 on default, other options: SHA384 or SHA512")
	flag.Parse()
	fmt.Println(*n)

	switch *n {
	case "SHA384":
		s := sha512.Sum384([]byte("x"))
		fmt.Printf("%x\n", s)
	case "SHA512":
		s := sha512.Sum512([]byte("x"))
		fmt.Printf("%x\n", s)
	default:
		s := sha256.Sum256([]byte("x"))
		fmt.Printf("%x\n", s)
	}
	return
}
