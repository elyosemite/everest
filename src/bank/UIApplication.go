package bank

import (
	"bufio"
	"fmt"
	"os"
)

func Menu() {
	fmt.Println("1. Ver saldo da conta")
	fmt.Println("2. Depositar")
	fmt.Println("3. Transferir")
	fmt.Println("4. Empréstimo")
	fmt.Println("5. Conta Polpança")
	fmt.Println("6. Conta corrente")
	fmt.Println("7. Pedir Empréstimo")
}

var limite float64 = 1_500.00
var saldo float64 = 0.0

func GetSaldo() {
	fmt.Printf("\n====== Seu saldo é: %.2f ======\n\n", saldo)
}

func Depositar() {
	valor := 0.0
	fmt.Print("Valor para depósito :> ")
	fmt.Scanf("%f", &valor)

	bufio.NewReader(os.Stdin).ReadBytes('\n')

	if saldo+valor > limite {
		fmt.Printf("\n====== Você atingiu seu limite! ======\n\n")
	} else {
		saldo += valor
		fmt.Printf("\n====== Você depositou: R$ %.2f ======\n\n", valor)
	}
}

func StartUp() {
	Menu()
	fmt.Print("\nEscolha sua operação :> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	switch input {
	case "1":
		GetSaldo()
		StartUp()
	case "2":
		Depositar()
		StartUp()
	case "3":
		fmt.Printf("\nTransferir\n")
	case "4":
		fmt.Printf("\nEmpréstimo\n")
	case "5":
		fmt.Printf("\nConta Polpança\n")
	case "6":
		fmt.Printf("\nConta Corrente\n")
	case "7":
		fmt.Printf("\nPedir Empréstimo\n")
	}
}
