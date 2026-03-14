package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é ímpar")
	}

	if 8%4 == 0 {
		fmt.Println("8 é divisível por 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("8 ou 7 é par")
	}

	// declarar e inicializar uma variável antes da condição
	if num := 9; num < 0 {
		fmt.Println("9 é negativo")
	} else if num < 10 {
		fmt.Println("9 tem um dígito")
	} else {
		fmt.Println("9 tem múltiplos dígitos")
	}
}
