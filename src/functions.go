package src

import (
	"errors"
	"fmt"
)

func Sum(a int, b int) int {
	return a + b
}

func CalcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator cannot be zero.")
	}
	return numerator / denominator, numerator % denominator, nil
}

func PlayWithPointers() {
	var creature string = "dragon"
	var saphira string = "saphira"
	var pointer *string = &creature

	fmt.Println("createure:", creature)
	fmt.Println("pointer:", pointer)
	fmt.Println("pointer value:", *pointer)

	*pointer = "unicorn"
	fmt.Println("createure:", creature)
	fmt.Println("pointer:", pointer)
	fmt.Println("pointer value:", *pointer)
	pointer = &saphira
	fmt.Println("pointer", pointer)
}
