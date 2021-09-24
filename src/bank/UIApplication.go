package main

import (
	"fmt"
	"os"
	"bufio"
)

func Menu() {
	fmt.Println("\t1. Ver saldo da conta")
	fmt.Println("\t2. Depositar")
	fmt.Println("\t3. Transferir")
	fmt.Println("\t4. Empréstimo")
	fmt.Println("\t5. Conta Polpança")
	fmt.Println("\t6. Conta corrente")
	fmt.Println("\t7. Pedir Empréstimo\n\n")
}

func Saldo() int {
	return 1000
}
 // Menus

func StartUp() {
	Menu()
	fmt.Println("Escolha sua operação:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
  	
	switch input {
		case "1":
			fmt.Printf("Seu saldo é: %v", Saldo())
			break
		case "2":
			fmt.Printf("\nDepositar\n")
			break
		case "3":
			fmt.Printf("\nTransferir\n")
			break
		case "4":
			fmt.Printf("\nEmpréstimo\n")
			break
		case "5":
			fmt.Printf("\nConta Polpança\n")
			break
		case "6":
			fmt.Printf("\nConta Corrente\n")
			break
		case "7":
			fmt.Printf("\nPedir Empréstimo\n")
			break
	}
}

func main() {
	fmt.Println("Hello, Welcome to my Bank System\n\n")
	StartUp()
}
