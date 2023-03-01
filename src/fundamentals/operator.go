package main

import (
	"fmt"
)

func main() {
	x := 10
	y := "Yuri Melo, Software Developer"

	fmt.Printf("X: %v, %T\n", x, x)
	fmt.Printf("Y: %v, %T\n", y, y)

	x = 20

	fmt.Printf("X: %v, %T\n", x, x)
}