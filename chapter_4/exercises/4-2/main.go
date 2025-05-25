package main

import (
	"fmt"
)

/*
For our dynamically allocated strings, create a function substring that takes three parameters: an arrayString,
a starting position integer, and an integer length of characters. The function returns a pointer to a new dynamically allocated
string array. This string array contains the characters in the original string, starting at the specified position for the specified
length. The original string is unaffected by the operation. So if the original string was abcdefg, the position was 3,
and the length was 4, then the new string would contain cdef
*/
// substring crea una nueva cadena dinámica a partir de otra
func substring(original string, start, length int) *string {
	// Validar parámetros
	if start < 0 || start >= len(original) || length < 0 {
		return nil // Retornar nil si los parámetros son inválidos
	}

	end := start + length
	if end > len(original) {
		end = len(original) // Ajustar si la longitud excede el tamaño de la cadena original
	}

	// Crear una nueva cadena y devolver un puntero a ella
	newString := original[start:end]
	return &newString
}

func main() {
	originalString := "abcdefg"

	// Ejemplo de uso
	start := 3
	length := 4
	sub := substring(originalString, start, length)

	if sub != nil {
		fmt.Printf("Original: %s\n", originalString)
		fmt.Printf("Subcadena desde posición %d con longitud %d: %s\n", start, length, *sub)
	} else {
		fmt.Println("Error: parámetros inválidos.")
	}
}
