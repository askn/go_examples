package main

import (
	"fmt"
)

func BubbleSort(list []byte) []byte {
	for n := len(list); n != 0; {
		newn := 0
		for i := 1; i < n; i++ {
			if list[i-1] > list[i] {
				list[i-1], list[i] = list[i], list[i-1]
				newn = i
			}
		}
		n = newn
		// fmt.Printf("%d\n", n)
	}
	return list
}

func BSort(list []byte) []byte {
	for i := 1; i < len(list); i++ {
		for j := len(list) - 1; j >= 1; j-- {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}
	return list
}

func main() {
	fmt.Printf("%d\n", BubbleSort([]byte("askingedik")))
	fmt.Printf("%s\n", BubbleSort([]byte("askingedik")))

	fmt.Printf("%d\n", BubbleSort([]byte{5, 1, 4, 2, 8}))

	fmt.Printf("%d\n", BSort([]byte{5, 1, 4, 2, 8}))
}
