package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
READING A NUMBER WITH THREE OR FOUR DIGITS - Now it's available with all type of length

Write a program to read a number character by character and convert it to an integer,
using just one char variable and one int variable.
The number will have either three or four digits.
*/
func main() {
	var char byte
	var number int

	fmt.Println("Enter a number with three or four digits:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	input := scanner.Text()

	for i := 0; i < len(input); i++ {
		// %c	the character represented by the corresponding Unicode code point
		//_, err := fmt.Scanf("%c", &char)
		//if err != nil {
		//	fmt.Println("Error reading input:", err)
		//	os.Exit(1)
		//}

		char = input[i]

		if char == '\n' {
			break
		}

		// Obtengo el digito y lo convierto de char a una representacion numerica.
		digit := int(char - '0')
		number = number*10 + digit
	}

	fmt.Println("The number is:", number)
}
