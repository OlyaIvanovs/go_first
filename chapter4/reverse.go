package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(ar *[7]int) {
	arLen := 7 - 1
	for i, v := range *ar {
		ar[i], ar[arLen-i] = ar[arLen-i], v
		if (i + 1) >= (arLen - 1 - i) {
			break
		}
	}
}

func main() {
	s := [...]int{91, 92, 93, 94, 95, 96, 97}
	reverse2(&s)
	fmt.Println(s)
}
