package main

import "fmt"

func rotate(s []int, step int) []int {
	out := make([]int, len(s), len(s))
	copy(out, s[step:])
	copy(out[(len(out)-step):], s[:step])
	return out
}

func main() {
	toRotate := []int{0, 1, 2, 3, 4, 5}
	toRotate = rotate(toRotate, 2)
	fmt.Print(toRotate)
}
