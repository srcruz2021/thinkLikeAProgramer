package main

import (
	"fmt"
	"strconv"
)

// Función para calcular el dígito de verificación de un ISBN-13
func calculateCheckDigit(isbn string) int {
	sum := 0
	for i, digit := range isbn {
		num, _ := strconv.Atoi(string(digit))
		if i%2 == 0 {
			sum += num
		} else {
			sum += num * 3
		}
	}
	checkDigit := (10 - (sum % 10)) % 10
	return checkDigit
}

// Función para validar un ISBN-13
func validateISBN13(isbn string) bool {
	if len(isbn) != 13 {
		return false
	}
	checkDigit := calculateCheckDigit(isbn[:12])
	return checkDigit == int(isbn[12]-'0')
}

func main() {
	// Prueba generando el dígito de verificación
	isbnBase := "978030640615" // Falta el último dígito
	checkDigit := calculateCheckDigit(isbnBase)
	fmt.Printf("El dígito de verificación es: %d\n", checkDigit)

	// Prueba validando un ISBN completo
	isbnComplete := "9780306406157"
	if validateISBN13(isbnComplete) {
		fmt.Println("El ISBN es válido.")
	} else {
		fmt.Println("El ISBN no es válido.")
	}
}
