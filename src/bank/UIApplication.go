package main

import (
	"fmt"
	"os"
	"bufio"
)

func Menu() {
	fmt.Println("1. Ver saldo da conta")
	fmt.Println("2. Depositar")
	fmt.Println("3. Transferir")
	fmt.Println("4. Empréstimo")
	fmt.Println("5. Conta Polpança")
	fmt.Println("6. Conta corrente")
	fmt.Println("7. Pedir Empréstimo\n")
}

var limite float64 = 15000.00
var saldo float64 = 0.0

func Saldo() float64 {
	return saldo
}

func Depositar() {
	valor := 0.0
	fmt.Println("Valor para depósito: ")
	fmt.Scanf("%f", &valor)

	if saldo + valor > limite {
		fmt.Printf("Você atingiu seu limite!")
	} else {
		saldo += valor
		fmt.Printf("Você depositou: R$ %f\n", valor)
		fmt.Printf("Saldo atual: R$ %f\n", saldo)
	}
}

func StartUp() {
	Menu()
	fmt.Println("Escolha sua operação:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
  	
	switch input {
		case "1":
			fmt.Printf("Seu saldo é: %f\n", Saldo())
			StartUp()
			break
		case "2":
			Depositar()
			StartUp()
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
