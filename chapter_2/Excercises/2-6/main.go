package main

import (
	"fmt"
	"math/big"
)

/*
If you’ve learned about binary numbers and how to convert from decimal to binary and the reverse,
try writing programs to do those conversions with unlimited length numbers (but you can assume the
numbers are small enough to be stored in a standard C++ int).
*/
// Convertir de decimal a binario
func decimalToBinary(n *big.Int) string {
	return n.Text(2) // Representación en base 2
}

// Convertir de binario a decimal
func binaryToDecimal(bin string) *big.Int {
	n := new(big.Int)
	n.SetString(bin, 2) // Convierte desde base 2
	return n
}

func main() {
	// Convertir un número decimal a binario
	decimalNumber := big.NewInt(12345) // Ejemplo
	binaryRepresentation := decimalToBinary(decimalNumber)
	fmt.Printf("Decimal: %d -> Binario: %s\n", decimalNumber, binaryRepresentation)

	// Convertir un número binario a decimal
	binaryNumber := "11000000111001"
	decimalResult := binaryToDecimal(binaryNumber)
	fmt.Printf("Binario: %s -> Decimal: %d\n", binaryNumber, decimalResult)
}
