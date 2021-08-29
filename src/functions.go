package main

import "fmt"

func Soma(a int, b int) int {
	return a+b
}

func main() {
	result := Soma(1,1)

	fmt.Printf("Type of \'result\' variable is %T\n", result)
}