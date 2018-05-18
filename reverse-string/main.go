package main

import "fmt"

func main() {
	fmt.Println(reverse("Hello"))
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
