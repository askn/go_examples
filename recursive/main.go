package main

import (
	"fmt"
)

func fac(n int) int {
	if n == 0 {
		return 1
	}
	return fac(n-1) * n
}

func fib(n int) int {
	switch n {
	case 0, 1:
		return n
	}
	return fib(n-1) + fib(n-2)

}

func main() {
	fmt.Printf("factoriel(%d):\t%d\n", 3, fac(3))
	fmt.Printf("fibonacci(%d):\t%d\n", 5, fib(5))
}
