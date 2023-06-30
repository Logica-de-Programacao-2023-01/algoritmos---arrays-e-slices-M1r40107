package main

import "fmt"

func main() {
	var numeros = [6]float64{1.5, 2.4, 3.8, 4.7, 5.5, 6.9}
	var x float64
	var i int = 0
	fmt.Println(numeros)

	fmt.Print("Escolha um número para ser adicionado ao array em todas as posições.")
	fmt.Scan(&x)

	for i < len(numeros) {
		numeros[i] = numeros[i] + x
		i++
	}
	fmt.Println("O array final é: ", numeros)

}
