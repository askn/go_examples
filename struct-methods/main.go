package main

import (
	"fmt"
)

type TwoInts struct {
	a, b int
}

func (t TwoInts) AddItems() int {
	return t.a + t.b
}

func (t TwoInts) AddParam(p int) int {
	return t.a + t.b + p
}

/*
method: belirli türdeki değişkenler üzerine etki eden fonksiyonlar
genel:
func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
*/

type IntVector []int

// sadece struct'lara özgü değil
func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	ti := TwoInts{1, 9}
	fmt.Printf("%d + %d = %d\n", ti.a, ti.b, ti.AddItems())
	fmt.Printf("%d + %d + %d = %d\n", ti.a, ti.b, 4, ti.AddParam(4))

	fmt.Println(IntVector{1, 2, 3}.Sum())

}
