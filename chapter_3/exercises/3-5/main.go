package main

import (
	"fmt"
	"strings"
)

/*
Have the previous program convert the ciphertext back to the plaintext to verify the encoding and decoding
*/
// Definimos el cifrado y el descifrado
const cipher = "QWERTYUIOPLKJHGFDSAZXCVBNM"

var decipher map[byte]byte

// Función para inicializar el mapa de descifrado
func initDecipher() {
	decipher = make(map[byte]byte)
	for i := 0; i < len(cipher); i++ {
		decipher[cipher[i]] = byte('A' + i)
	}
}

// Función para cifrar el texto
func encrypt(plaintext string) string {
	plaintext = strings.ToUpper(plaintext) // Convertimos a mayúsculas
	var ciphertext strings.Builder

	for _, char := range plaintext {
		if char >= 'A' && char <= 'Z' {
			ciphertext.WriteByte(cipher[char-'A']) // Sustitución basada en el índice
		} else {
			ciphertext.WriteByte(byte(char)) // Mantener los caracteres que no son letras
		}
	}
	return ciphertext.String()
}

// Función para descifrar el texto
func decrypt(ciphertext string) string {
	var plaintext strings.Builder

	for _, char := range ciphertext {
		if val, exists := decipher[byte(char)]; exists {
			plaintext.WriteByte(val) // Sustitución inversa
		} else {
			plaintext.WriteByte(byte(char)) // Mantener caracteres que no son letras
		}
	}
	return plaintext.String()
}

func main() {
	// Inicializamos el mapa de descifrado
	initDecipher()

	// Solicitamos al usuario que ingrese el mensaje en texto plano
	var plaintext string
	fmt.Println("Ingrese el mensaje en texto plano:")
	fmt.Scanln(&plaintext)

	// Ciframos y desciframos el mensaje
	ciphertext := encrypt(plaintext)
	decodedText := decrypt(ciphertext)

	// Mostramos los resultados
	fmt.Println("Texto cifrado:", ciphertext)
	fmt.Println("Texto descifrado:", decodedText)
}
