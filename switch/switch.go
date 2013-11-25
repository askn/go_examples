package main

import (
	"fmt"
)

func main() {
	var i = 0

	switch i {
	case 0: //empty case body, nothing is executed when i==0
	case 1, 2, 3, 4:
		fmt.Println("girdi")
	}

	switch i {
	case 0:
		fallthrough
	case 1:
		fmt.Println("2. switch girdi") // çalıştı

	default:
	}

	switch a := 1 * 2; a {
	case 123:

	case 2:
		fmt.Println("3. switch girdi: 2") // çalıştı
	}

}
