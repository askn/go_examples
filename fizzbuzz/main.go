package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		switch fizz, buzz := i%3, i%5; {
		case fizz == 0 && buzz == 0:
			fmt.Printf("fizzbuzz: %d\n", i)
		case fizz == 0:
			fmt.Printf("fizz: %d\n", i)
		case buzz == 0:
			fmt.Printf("buzz: %d\n", i)
		}
	}
}
