package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
To make the ciphertext problem even more challenging, have your program randomly generate the cipher array instead of a hard-coded
const array. Effectively, this means placing a random character in each element of the array, but remember that you can’t substitute a
letter for itself. So the first element can’t be A, and you can’t use the same letter for two substitutions—that is, if the first
element is S, no other element can be S
*/
// Genera un cifrado aleatorio sin repeticiones ni sustituciones idénticas
func generateCipher() map[byte]byte {
	rand.Seed(time.Now().UnixNano()) // Inicializamos la semilla aleatoria

	letters := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	shuffled := make([]byte, 26)
	copy(shuffled, letters)

	// Mezclamos aleatoriamente los caracteres evitando sustituciones idénticas
	for i := 0; i < len(shuffled); i++ {
		swapIndex := rand.Intn(26)
		if shuffled[i] == letters[i] || shuffled[swapIndex] == letters[swapIndex] {
			swapIndex = (swapIndex + 1) % 26 // Aseguramos que el cambio sea válido
		}
		shuffled[i], shuffled[swapIndex] = shuffled[swapIndex], shuffled[i]
	}

	// Mapeamos cada letra original con su reemplazo
	cipher := make(map[byte]byte)
	for i := 0; i < 26; i++ {
		cipher[letters[i]] = shuffled[i]
	}

	return cipher
}

// Función para cifrar el texto
func encrypt(plaintext string, cipher map[byte]byte) string {
	ciphertext := ""
	for i := 0; i < len(plaintext); i++ {
		char := plaintext[i]
		if char >= 'A' && char <= 'Z' {
			ciphertext += string(cipher[char])
		} else {
			ciphertext += string(char) // Mantener los caracteres que no son letras
		}
	}
	return ciphertext
}

// Función para descifrar el texto
func decrypt(ciphertext string, cipher map[byte]byte) string {
	decipher := make(map[byte]byte)
	for key, value := range cipher {
		decipher[value] = key
	}

	plaintext := ""
	for i := 0; i < len(ciphertext); i++ {
		char := ciphertext[i]
		if val, exists := decipher[char]; exists {
			plaintext += string(val)
		} else {
			plaintext += string(char)
		}
	}
	return plaintext
}

func main() {
	// Generamos el cifrado aleatorio
	cipher := generateCipher()

	// Mostramos el cifrado generado
	fmt.Println("Cifrado generado:")
	for k, v := range cipher {
		fmt.Printf("%c -> %c\n", k, v)
	}

	// Solicitamos al usuario el mensaje en texto plano
	var plaintext string
	fmt.Println("\nIngrese el mensaje en texto plano:")
	fmt.Scanln(&plaintext)

	// Convertimos a mayúsculas para garantizar compatibilidad
	plaintext = strings.ToUpper(plaintext)

	// Ciframos y desciframos el mensaje
	ciphertext := encrypt(plaintext, cipher)
	decodedText := decrypt(ciphertext, cipher)

	// Mostramos los resultados
	fmt.Println("\nTexto cifrado:", ciphertext)
	fmt.Println("Texto descifrado:", decodedText)
}
