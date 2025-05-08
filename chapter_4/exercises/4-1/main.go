package main

/*
Design your own: Take a problem that you already know how to solve using an array but that is limited by the size of the array.
Rewrite the code to remove that limitation using a dynamically allocated array
*/

import (
	"fmt"
)

// Función para agregar un número a un array dinámico
func appendElement(slice []int, element int) []int {
	return append(slice, element) // Usa `append` para expandir dinámicamente
}

func main() {
	// Inicializamos un slice vacío
	var numbers []int

	// Agregamos elementos dinámicamente
	numbers = appendElement(numbers, 1)
	numbers = appendElement(numbers, 2)
	numbers = appendElement(numbers, 3)

	fmt.Println("Slice dinámico:", numbers) // Output: [1 2 3]
}
