package main

import (
	"fmt"
	"math"
)

var Pi float64

func init() {
	Pi = 4 * math.Atan(1)
}

func main() {
	fmt.Println("%g", Pi)
}
