package main

import (
	"fmt"
)

func main() {
	// performans açısından slice
	var m map[string]int

	m = map[string]int{"one": 1, "two": 2}
	m["three"] = 3

	// ok: var mı yok mu
	value, ok := m["one"]
	fmt.Printf("%d-%t\n", value, ok) // 1-true
	value2, ok2 := m["ten"]
	fmt.Printf("%d-%t\n", value2, ok2) // 0-false

	if v, ok := m["three"]; ok {
		fmt.Printf("\t%d\n", v)
	}

	// map_name, key
	delete(m, "three")

	for k, v := range m {
		fmt.Printf("%s - %d\n", k, v)
	}

	// mm := make(map[int]string)
	// mmm := map[int]string{}

	/* ----------------------------- */

	mf := map[int]func() int{
		1: func() int { return 1 },
		2: func() int { return 2 },
	}
	fmt.Printf("%d\n", mf[1])   // 0x400ed3
	fmt.Printf("%d\n", mf[1]()) // 1

	mf2 := map[int]func(int) int{
		1: func(x int) int { return 1 * x },
		2: func(x int) int { return 2 * x },
	}
	fmt.Printf("%d\n", mf2[2](2))

	/* ----------------------------- */

	// map tutan slice
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Value of items: %v\n", items)

	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}
	// istenilen olmadı çünkü item sadece slice'ın elemanını kopyaladı
	// bir sonraki işlemde uçtu
	fmt.Printf("Value of items2: %v\n", items2)
}
