package main

import "fmt"

func main() {
	var numeros = []int{1, 2, 3, 4, 5}
	var x int

	fmt.Print("Escolha um número inteiro.")
	fmt.Scan(&x)

	if x >= 1 && x <= 5 {
		fmt.Println("Esse número já está na lista.")
		fmt.Println(numeros)
	} else {
		fmt.Println("Esse número não está na lista, mas será adicionado.")
		numeros = append(numeros, x)

		fmt.Println(numeros)
	}

}
