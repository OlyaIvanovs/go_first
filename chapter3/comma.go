// Inserts commas in nonnegative decimal integer string
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := 3
	if n < 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// None-recursive version of the func above
func comma2(s string) string {
	var buf bytes.Buffer
	n := len(s)
	commaStep := 3
	comma := 3
	if n%commaStep != 0 {
		comma = n % commaStep
	}
	if n < 3 {
		return s
	}
	for i := 0; i < n; i++ {
		fmt.Fprintf(&buf, "%c", s[i])
		if (i+1) == comma && (i+1) < n {
			buf.WriteByte(',')
			comma += commaStep
		}
	}
	return buf.String()
}

func main() {
	// s := "Mama myla ramy"
	// n := len(s)
	// fmt.Print(s[:n-3])
	// fmt.Print(s[:3])fmt.Print(s[:3])
	s := comma2("134567890")
	fmt.Println(s)
}
