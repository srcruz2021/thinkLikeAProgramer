package main

import (
	"fmt"
	"math/big"
	"strings"
)

// Función para detectar el formato de entrada
func detectFormat(input string) (string, *big.Int) {
	n := new(big.Int)

	if strings.HasPrefix(input, "0b") { // Binario
		n.SetString(input[2:], 2)
		return "binario", n
	} else if strings.HasPrefix(input, "0x") { // Hexadecimal
		n.SetString(input[2:], 16)
		return "hexadecimal", n
	} else { // Decimal por defecto
		n.SetString(input, 10)
		return "decimal", n
	}
}

// Función para imprimir el número en los tres formatos
func printFormats(n *big.Int) {
	fmt.Printf("Decimal: %s\n", n.Text(10))
	fmt.Printf("Binario: 0b%s\n", n.Text(2))
	fmt.Printf("Hexadecimal: 0x%s\n", n.Text(16))
}

func main() {
	var input string

	// Solicitar entrada al usuario
	fmt.Print("Introduce un número (decimal, binario: 0b..., hexadecimal: 0x...): ")
	fmt.Scanln(&input)

	// Detectar formato y convertir
	format, number := detectFormat(input)

	fmt.Printf("Formato detectado: %s\n", format)
	printFormats(number)
}
