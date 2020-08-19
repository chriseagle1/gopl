package main

import (
	"fmt"
	"math"
)

func main() {
	var a uint8 = 11
	var b uint8 = 13

	fmt.Printf("%b\t%b\n", a, b)

	fmt.Printf("%b\n", a &^ b)

	for x := 0; x < 8 ; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}
