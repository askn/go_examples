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

func uniq(str string) string {
	arr := []byte(str)
	uniq_arr := []byte{}

	for _, v := range arr {
		i := 0
		for _, u := range uniq_arr {
			if v == u {
				i++
			}
			if i > 1 {
				break
			}
		}
		if i == 0 {
			uniq_arr = append(uniq_arr, v)
		}
	}
	return string(uniq_arr)
}

func main() {
	fmt.Printf("%s\n", reverse("askin"))
	fmt.Printf("%s\n", uniq("asskin  gedikk"))
}
