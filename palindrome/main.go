package main

import "fmt"

func main() {
	a := "abcba"
	b := reverse(a) == a

	fmt.Println(b)
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
