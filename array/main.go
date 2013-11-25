package main

import (
	"fmt"
)

func main() {
	var ar [5]int

	// var ar1 = new([5]int)
	// ar1 type = *[5]int

	for i := 0; i < len(ar); i++ {
		ar[i] = i * i
	}

	for i := range ar {
		fmt.Printf("%d\n", ar[i])
	}

	var intArr = [5]int{0, 1, 2, 3}
	fmt.Printf("%d\n", intArr[4]) // 0
	// diğerleri 0 olur
	// aynı şekilde
	var arrKeyValue = [6]string{4: "foo", 5: "bar"}
	arrKeyValue[0] = "joo"

	sumArr := [...]int{1, 2, 3, 4}
	fmt.Printf("%d\n", Sum(&sumArr))
}

func Sum(a *[4]int) (sum int) {
	for _, n := range a {
		sum += n
	}
	return sum
}
