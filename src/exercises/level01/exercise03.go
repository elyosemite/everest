package main

import "fmt"

/*
- Utilizando a solução do exercicio anterior:
Em package-level scope, atribua os seguintes valores às variáveis:

1. para "x" atribua 42
2. para "y" atribua "James Bond"
3. para "z" atribua true

Na função main:
1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tip
string a uma variável de nome "s" utilizando o operador curto de declaração.
2. Demonstre a variável "s".
*/

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}