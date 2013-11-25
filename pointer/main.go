package main

import (
	"fmt"
)

// pointer değişkeni başka bir değerin bellek adresini içerir
func main() {
	i := 1
	fmt.Printf("%d\t%p\n", i, &i)

	var ip *int
	ip = &i
	fmt.Printf("%d\t%p\n", &ip, ip)

	var it *int
	fmt.Printf("%p\n", it)

	// %t boolean
	fmt.Printf("%t\n", *ip == i)

	*ip = 3
	fmt.Printf("%d\t%d\n", *ip, i)
	// 3 3 i'de değişti

	const c = 1
	// ip = &c
	// hata verir.

	var n *int
	// var n *int = nil
	// *n = 1
	fmt.Printf("%p\n", n)
}
