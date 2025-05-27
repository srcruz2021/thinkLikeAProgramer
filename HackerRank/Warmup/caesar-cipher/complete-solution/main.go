package main

import (
	"fmt"
	"unicode"
)

// caesarCipher encrypts a string by shifting each letter by k positions
func caesarCipher(s string, k int32) string {
	// Normalize shift in case k is greater than 26
	k = k % 26

	// Store the encrypted result
	encrypted := []rune{}

	// Iterate through each character in s
	for _, char := range s {
		if unicode.IsLetter(char) {
			base := 'A'
			if unicode.IsLower(char) {
				base = 'a'
			}
			// Apply shift while keeping the letter within bounds
			shifted := base + (char-base+int32(k))%26
			encrypted = append(encrypted, shifted)
		} else {
			// Keep non-letter characters unchanged
			encrypted = append(encrypted, char)
		}
	}

	return string(encrypted)
}

func main() {
	s := "middle-Outz"
	k := int32(2)
	fmt.Println("Encrypted:", caesarCipher(s, k)) // Expected output: "okffng-Qwvb"
}
