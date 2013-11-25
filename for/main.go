package main

import (
	"fmt"
)

func main() {
	i := 1
	for i < 10 {
		fmt.Printf("%d\n", i)
		i++
	}

	// sonsuz döngü
	// for {
	// }

	for pos, c := range "askn" {
		fmt.Printf("%d-%c\n", pos, c)
	}

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	// goto kullanmak önerilmez
HERE:
	if i == 10 {
		goto HERE
	}
}
