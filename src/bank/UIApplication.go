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

func StartUp(){
	Menu()
	fmt.Println("Escolha sua operação:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
  	
	switch input {
		case "1":
			fmt.Printf("\nopção 1\n")
			break
		case "2":
			fmt.Printf("\nopção 2\n")
			break
		case "3":
			fmt.Printf("\nopção 3\n")
			break
		case "4":
			fmt.Printf("\nopção 4\n")
			break
		case "5":
			fmt.Printf("\nopção 5\n")
			break
		case "6":
			fmt.Printf("\nopção 6\n")
			break
		case "7":
			fmt.Printf("\nopção 7\n")
			break

	}

}

func main() {
	fmt.Println("Hello, Welcome to my Bank system\n\n")
	StartUp()
}
