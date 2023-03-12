package main

import "fmt"

/*
Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y" e "z".

	1. 42
	2. James Bond
	3. true

- Agora demonstre os valores nestas variáveis, com:
	1. a única declaração print
	2. Múltiplas declarações print
*/

func main() {
	x := 42
	y := "James Bond"
	z := true
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x = %v\ny = %v\nz = %v", x, y, z)
}