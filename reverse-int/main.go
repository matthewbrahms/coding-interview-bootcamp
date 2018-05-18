package main

import "fmt"

func main() {
	fmt.Println(reverseInt(123456))
}

func reverseInt(i int) int {
	n := 0

	for i > 0 {
		r := i % 10
		n *= 10
		n += r
		i /= 10
	}

	return n
}
