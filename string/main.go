package main

import (
	"fmt"
)

func reverse(str string) string {
	chars := []byte(str)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	fmt.Printf("%s\n", reverse("askin"))
}
