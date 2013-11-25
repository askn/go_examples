package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		a := rand.Int()
		fmt.Printf("%d\n", a)
	}

	fmt.Println("---")

	for i := 0; i < 4; i++ {
		b := rand.Intn(100)
		fmt.Printf("%d\n", b)
	}

	fmt.Println("---")

	for i := 0; i < 4; i++ {
		timens := int64(time.Now().Nanosecond())
		rand.Seed(timens)
		fmt.Printf("%d\n", rand.Int())
	}

}
