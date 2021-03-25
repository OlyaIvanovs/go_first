// Inserts commas in nonnegative decimal integer string
package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func comma(s string) string {
	n := 3
	if n < 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// None-recursive version of the func above. It deals also with floating-point numbers
func comma2(s string) string {
	var buf bytes.Buffer

	index := strings.Index(s, ".")
	var endString string
	if index != -1 {
		endString = s[index:]
		s = s[:index]
		fmt.Println(endString)
	}

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
	if endString != "" {
		fmt.Fprintf(&buf, "%s", endString)
	}
	return buf.String()
}

func sortString(s string) string {
	snew := strings.Split(s, "")
	sort.Strings(snew)
	return strings.Join(snew, "")
}

func anagramCheck(s1, s2 string) bool {
	sortedS1 := sortString(s1)
	sortedS2 := sortString(s2)
	if sortedS1 == sortedS2 {
		return true
	}
	return false
}

func main() {
	// s := "Mama myla ramy"
	// n := len(s)
	// fmt.Print(s[:n-3])
	// fmt.Print(s[:3])fmt.Print(s[:3])
	s := comma2("134567.890")
	s2 := comma2("134567890")
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(anagramCheck("2a3", "a23"))
}
