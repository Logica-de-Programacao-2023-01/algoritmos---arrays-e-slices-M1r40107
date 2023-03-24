package main

import "fmt"

func main() {
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	numbers = append(numbers[:3], numbers[4:]...)
	fmt.Println(numbers)

}