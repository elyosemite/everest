package main

import (
	"fmt"
	"strings"
)

func Sum(a int, b int) int {
	return a+b
}

func main() {
	var firstNumber int
	var secondNumber int
	var name string

	fmt.Scanln(&name)
	name = strings.TrimSpace(name)

	fmt.Scanln(&firstNumber)
	fmt.Scanln(&secondNumber)

	result := Sum(firstNumber, secondNumber)
	
	fmt.Println(result)
	fmt.Printf("Your name is %v\n", name)

	fmt.Printf("Type of \"result\" variable is %T\n", result)
}