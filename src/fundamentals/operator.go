package main

import "fmt"

var my_name = "Yuri Melo"

func main() {
	x := 10
	y := "Yuri Melo, Software Developer"

	fmt.Printf("Name: %v, %T\n", my_name, my_name)
	fmt.Printf("X: %v, %T\n", x, x)
	fmt.Printf("Y: %v, %T\n", y, y)

	fmt.Printf("%T", x)

	x, second_variable := 20, 9132

	fmt.Printf("X: %v, %T\n", x, x)
	fmt.Printf("Second_Variable: %v, %T\n", second_variable, second_variable)
}