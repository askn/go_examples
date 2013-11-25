package main

import "fmt"

func main() {
	var a int = 1
	var b int
	var c = 1

	var (
		d = 1
		e int
	)

	f := 1

	var g, h, i = 1, true, "askn"
	j, k, l := 0, false, "gdk"

	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g, h, i, j, k, l)

	// swap
	fmt.Println(g, j)
	g, j = j, g
	fmt.Println(g, j)

	// unused
	var _ = "GO"
}
