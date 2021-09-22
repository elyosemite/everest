package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
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

func Enter(){
	Menu()
	fmt.Println("Escolha sua operação:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
  	
	switch text {
		case "1":
			fmt.Println("opção 1")
			break
		case "2":
			fmt.Println("opção 2")
			break
	}

}

func main() {
	fmt.Println("Hello, Welcome to my Bank system\n\n")
	Enter()
}
