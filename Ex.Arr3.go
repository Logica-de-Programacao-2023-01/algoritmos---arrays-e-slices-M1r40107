package main

import "fmt"

func main() {
	var w float64
	var x float64
	var y float64
	var z float64
	fmt.Print("Escolha um número.")
	fmt.Scan(&w)
	fmt.Print("Escolha um número.")
	fmt.Scan(&x)
	fmt.Print("Escolha um número.")
	fmt.Scan(&y)
	fmt.Print("Escolha um número.")
	fmt.Scan(&z)

	var numeros = [4]float64{w, x, y, z}
	fmt.Println(numeros)
	fmt.Println(w * x * y * z)
}
