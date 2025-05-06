package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Función para contar las palabras en el texto
func countWords(text string) int {
	words := strings.Fields(text) // Divide el texto en palabras usando espacios
	return len(words)
}

// Función para encontrar la longitud de la palabra más larga
func longestWordLength(text string) int {
	words := strings.Fields(text)
	maxLength := 0
	for _, word := range words {
		length := len(word)
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

// Función para contar el número de vocales en una palabra
func countVowels(word string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range word {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

// Función para encontrar la palabra con más vocales
func wordWithMostVowels(text string) (string, int) {
	words := strings.Fields(text)
	maxVowelCount := 0
	wordWithMaxVowels := ""

	for _, word := range words {
		vowelCount := countVowels(word)
		if vowelCount > maxVowelCount {
			maxVowelCount = vowelCount
			wordWithMaxVowels = word
		}
	}
	return wordWithMaxVowels, maxVowelCount
}

// Función para contar caracteres alfabéticos
func countAlphabeticCharacters(text string) int {
	count := 0
	for _, char := range text {
		if unicode.IsLetter(char) {
			count++
		}
	}
	return count
}

func main() {
	// Solicitar entrada al usuario
	fmt.Print("Introduce una línea de texto: ")
	var input string
	getline(&input) // Leer línea completa

	// Obtener estadísticas
	wordCount := countWords(input)
	longestWord := longestWordLength(input)
	wordMostVowels, maxVowelCount := wordWithMostVowels(input)
	letterCount := countAlphabeticCharacters(input)

	// Mostrar resultados
	fmt.Printf("\nNúmero total de palabras: %d\n", wordCount)
	fmt.Printf("Longitud de la palabra más larga: %d\n", longestWord)
	fmt.Printf("Palabra con más vocales: \"%s\" (%d vocales)\n", wordMostVowels, maxVowelCount)
	fmt.Printf("Cantidad total de caracteres alfabéticos: %d\n", letterCount)
}

// Función para leer una línea completa de entrada
func getline(input *string) {
	_, err := fmt.Scanln(input)
	if err != nil {
		fmt.Println("Error al leer la entrada.")
	}
}
