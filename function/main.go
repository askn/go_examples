package main

import (
	"fmt"
)

func foo() (x int) {
	x = 19
	return
}

func bar() int {
	x := 19
	return x
}

func hop(a ...int) {
	for _, c := range a {
		// defer LİFO mantığında bir stack
		// bulunduğu blocktaki işlemler bitene kadar beklet
		defer fmt.Printf("%d\n", c)
		defer fmt.Printf("%d\n", c)
	}
	// 4 4 3 3 2 2 1 1
}

func main() {
	fmt.Printf("%d\n", foo())
	fmt.Printf("%d\n", bar())

	hop(1, 2, 3, 4)

	a := new(int)
	b := make([]int, 2)
	fmt.Printf("%d\t%p\n", *a, a)
	fmt.Printf("%d\t\n", b)
}

/*

new(type) belirtilen tipte bellekte yer ayırır ve onun adresini döner
a:= new(int) a bir pointer
*a = 0
a 0x123ıasdj

make(type)
ise ilkler ve döndürür sadece map, slice, channel
*/
