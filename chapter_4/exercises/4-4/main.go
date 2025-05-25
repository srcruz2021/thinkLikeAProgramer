package main

import "fmt"

/*
Change the implementation of our strings such that location[0] in the array stores the size of the array (and therefore location[1]
stores the first actual character in the string), rather than using a null-character terminator.
Implement each of the three functions, append, concatenate, and charactertAt, taking advantage of the stored size information whenever
possible. Because we’ll no longer be using the null-termination convention expected by the standard output stream, you’ll need to write
your own output function that loops through its string parameter, displaying the characters.
*/

// arrayString representa una cadena donde el índice 0 almacena el tamaño.
type arrayString []rune

// append agrega un carácter al final de la cadena.
func appendString(str *arrayString, ch rune) {
	*str = append(*str, ch)
	(*str)[0] = rune(len(*str) - 1) // Actualizamos el tamaño almacenado en la primera posición
}

// concatenate concatena dos cadenas.
func concatenate(str1, str2 *arrayString) {
	*str1 = append(*str1, (*str2)[1:]...) // Agregamos los caracteres de str2 a str1
	(*str1)[0] = rune(len(*str1) - 1)     // Actualizamos el tamaño de str1
}

// characterAt devuelve el carácter en una posición específica.
func characterAt(str arrayString, index int) rune {
	if index < 0 || index >= int(str[0]) {
		return '\000' // Retorna un carácter nulo si el índice es inválido
	}
	return str[index+1] // Ajustamos por el desplazamiento
}

// printString imprime la cadena sin la información del tamaño inicial.
func printString(str arrayString) {
	for i := 1; i < len(str); i++ {
		fmt.Print(string(str[i]))
	}
	fmt.Println()
}

func main() {
	// Inicializamos una cadena con su tamaño en la primera posición
	source := arrayString{4, 'a', 'b', 'c', 'd'}
	printString(source) // Output: abcd

	// Agregamos un carácter
	appendString(&source, 'e')
	printString(source) // Output: abcde

	// Concatenamos con otra cadena
	target := arrayString{2, 'x', 'y'}
	concatenate(&source, &target)
	printString(source) // Output: abcdexy

	// Obtenemos un carácter en una posición específica
	fmt.Println("Caracter en posición 2:", string(characterAt(source, 2))) // Output: c
}
