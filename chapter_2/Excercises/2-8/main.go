package main

import (
	"fmt"
	"math/big"
)

// Función para convertir un número de una base a decimal
func toDecimal(number string, base int) *big.Int {
	n := new(big.Int)
	n.SetString(number, base)
	return n
}

// Función para convertir un número decimal a una base dada
func fromDecimal(n *big.Int, base int) string {
	return n.Text(base)
}

func main() {
	var number string
	var baseFrom, baseTo int

	// Solicitar entrada al usuario
	fmt.Print("Introduce el número a convertir: ")
	fmt.Scanln(&number)
	fmt.Print("Introduce la base de entrada (2-16): ")
	fmt.Scanln(&baseFrom)
	fmt.Print("Introduce la base de salida (2-16): ")
	fmt.Scanln(&baseTo)

	// Validación de bases
	if baseFrom < 2 || baseFrom > 16 || baseTo < 2 || baseTo > 16 {
		fmt.Println("Las bases deben estar entre 2 y 16.")
		return
	}

	// Convertir entre bases
	decimal := toDecimal(number, baseFrom)
	converted := fromDecimal(decimal, baseTo)

	fmt.Printf("El número %s en base %d es %s en base %d.\n", number, baseFrom, converted, baseTo)
}
