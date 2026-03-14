package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// o underscore (_) está ali porque o range retorna dois valores quando percorre slices/arrays:
	// - O índice (posição do elemento).
	// - O valor (conteúdo do elemento).
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
