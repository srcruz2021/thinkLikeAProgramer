package main

import (
	"fmt"
)

/*
Write a program that is given an array of integers and determines the mode, which is the number that appears most frequently in the array. en go
*/

// Función para encontrar la moda en un arreglo de enteros
func findMode(arr []int) int {
	frequency := make(map[int]int)
	maxFreq := 0
	mode := arr[0]

	// Contamos la frecuencia de cada número
	for _, num := range arr {
		frequency[num]++
		if frequency[num] > maxFreq {
			maxFreq = frequency[num]
			mode = num
		}
	}

	return mode
}

func main() {
	// Definimos el arreglo de números
	numbers := []int{2, 3, 4, 2, 7, 2, 3, 3, 3, 8, 8, 8, 8, 8, 5}

	// Calculamos la moda
	mode := findMode(numbers)

	// Mostramos el resultado
	fmt.Printf("La moda del arreglo es: %d\n", mode)
}
