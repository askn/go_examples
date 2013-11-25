package main

import (
	"fmt"
	"runtime"
)

var prompt = "Sistemin: %s\n"

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "win :(")
	} else {
		prompt = fmt.Sprintf(prompt, runtime.GOOS+" :)")
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	// fonksiyon return ile bitmezse hata verir
	// return x
	// x = 2

	return x
}

func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

func main() {
	fmt.Printf(prompt)
	max := 20
	if val := 10; isGreater(val, max) {
		fmt.Println("büyük")
	} else {
		fmt.Println("küçük")
	}
}
