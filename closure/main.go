package main

import (
	"fmt"
)

/*
oluşturulan fonksiyonda local değişkenler değerlerini korur
*/

func add2() func(x int) int {
	return func(x int) int {
		return x + 2
	}
}

func adder() func(int) int {
	var y int
	return func(d int) int {
		y += d
		return y
	}
}
func main() {
	func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("\t%d\n", i)
		}
	}()

	f := func(x, y int) int { return x + y }
	fmt.Printf("%d\n", f(1, 2))
	fmt.Printf("%d\n\n", f(1, 2))

	for i := 0; i < 4; i++ {
		ff := func(x, y int) int { return x * y }(i, i+1)
		fmt.Printf("%d\t%d\n", i, ff)
	}

	fo := add2()
	fmt.Printf("%d\n\n", fo(2))

	bar := adder()
	bar(5)
	fmt.Printf("%d\n", bar(6))
}
