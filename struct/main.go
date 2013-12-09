package main

import (
	"fmt"
	"strings"
	"unsafe"
)

type human struct {
	name string
	age  int
}

type human2 struct {
	// embedded
	human
	city string
}

func upName(h *human) {
	h.name = strings.ToUpper(h.name)
}

// başka pakette human oluşturup kullanmak için
// x:= new(main.human) derlenmez çünkü human private
// x:= main.NewHuman() derlenir
func NewHuman() *human {
	h := new(human)
	return h
}
func main() {
	var x *human
	x = new(human)
	fmt.Printf("%d\n", x.age)

	// ---
	h := &human{age: 21}
	h1 := &human{"askn", 21}
	fmt.Printf("%d-%d\n", h.age, h1.age)

	// ***
	var p1 human
	p1.name = "Haydar"
	upName(&p1)
	fmt.Printf("%s\n", p1.name)

	// ---

	p2 := new(human)
	p2.name = "Haydar"
	(*p2).name = "Haydar" // bu da geçerli
	upName(p2)
	fmt.Printf("%s\n", p2.name)

	// ---

	p3 := &human{"Haydar", 19}
	upName(p3)
	fmt.Printf("%s\n", p3.name)

	fmt.Printf("sizeof %d\n", unsafe.Sizeof(p3))
	fmt.Printf("sizeof %d\n", unsafe.Sizeof(p3.age))
	fmt.Printf("sizeof %d\n", unsafe.Sizeof(p3.name))

	// embedded struct
	// hx := human2{human{"askn", 19}, "amasya"}
	hx := human2{*p3, "amasya"}
	fmt.Println(hx.name)
	fmt.Println(hx.human)
}
