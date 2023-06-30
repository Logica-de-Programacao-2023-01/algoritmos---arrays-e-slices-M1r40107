package main

import "fmt"

func main() {
	array := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceRes := []float64{}

	for i := 0; i < len(array); i++ {
		if array[i] > 5 {
			sliceRes = append(sliceRes, array[i])
		}
	}
	fmt.Println("Elementos que são maiores do que 5:", sliceRes)

}
