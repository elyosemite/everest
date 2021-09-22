package main

import "fmt"

func Menu() {
	fmt.Println("Escolha sua operação")
	fmt.Println("\t1. Ver saldo da conta")
	fmt.Println("\t2. Depositar")
	fmt.Println("\t3. Transferir")
	fmt.Println("\t4. Empréstimo")
	fmt.Println("\t5. Conta Polpança")
	fmt.Println("\t6. Conta corrente")
	fmt.Println("\t7. Pedir Empréstimo\n\n")
}

func main() {
	fmt.Println("Hello, Welcome to my Bank system\n\n")
	Menu()
}
