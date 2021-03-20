// Echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main_echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
