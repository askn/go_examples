package main

import (
	"fmt"
)

func fib() func() uint64 {
	x, y := 0, 1
	return func() uint64 {
		x, y = y, x+y
		return uint64(x)
	}
}

func Fibonacci(c int) uint64 {
	f := fib()
	for i := 0; i < c-1; i++ {
		f()
	}
	return uint64(f())
}

func main() {
	// fmt.Printf("%d\n", Fibonacci(7))
	// fmt.Printf("%d\n", fibRec(7))
	fmt.Printf("%d\n", fibGen(10000000))
	fmt.Printf("%d\n", Fibonacci(10000000))
	fmt.Printf("%d\n", projectEuler2())
}

func fibRec(n int) uint64 {
	switch n {
	case 0, 1:
		return uint64(n)
	}
	return fibRec(n-1) + fibRec(n-2)
}

func fibGen(n int) uint64 {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return uint64(a)
}

func projectEuler2() int {
	var total = 0

	a, b := 0, 1
	for {
		a, b = b, a+b
		if a > 4000000 {
			break
		} else if a%2 == 0 {
			total += a
		}

	}
	return total
}
