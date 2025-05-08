package main

import (
	"fmt"
	"strings"
)

/*
Here’s a variation on the array of const values. Write a program for creating a substitution cipher problem.
In a substitution cipher problem, all messages are made of uppercase letters and punctuation.
The original message is called the plaintext, and you create the ciphertext by substituting each letter with another
letter (for example, each C could become an X). For this problem, hard-code a const array of 26 char elements for the cipher,
and have your program read a plaintext message and output the equivalent ciphertext.
*/
// Definimos el cifrado como una constante
const cipher = "QWERTYUIOPLKJHGFDSAZXCVBNM"

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

func main() {
	// Solicitamos al usuario que ingrese el mensaje en texto plano
	var plaintext string
	fmt.Println("Ingrese el mensaje en texto plano:")
	fmt.Scanln(&plaintext)

	// Encriptamos el mensaje
	ciphertext := encrypt(plaintext)
	fmt.Println("Texto cifrado:", ciphertext)
}
