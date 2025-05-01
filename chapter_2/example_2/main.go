package main

import (
	"fmt"
	"os"
)

/*
PROBLEM: LUHN CHECKSUM VALIDATION

The Luhn formula is a widely used system for validating identification numbers. Using the original number, double the value of every other digit. Then add the values of the individual digits together (if a doubled value now has two digits, add the digits individually). The identification number is valid if the sum is divisible by 10.

Write a program that takes an identification number of arbitrary length and determines whether the number is valid under the Luhn formula. The program must process each character before reading the next one.
*/

func main() {
	var char byte
	var sum int
	var double bool

	fmt.Println("Enter the identification number:")

	for {
		_, err := fmt.Scanf("%c", &char)
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		if char == '\n' {
			break
		}

		digit := int(char - '0')

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	if sum%10 == 0 {
		fmt.Println("The identification number is valid.")
	} else {
		fmt.Println("The identification number is invalid.")
	}
}

/*
Explanation:
Reading Input: The program reads each character of the identification number one by one.
Doubling Every Other Digit: The double variable is used to keep track of whether the current digit should be doubled. If the doubled digit is greater than 9, subtract 9 from it.
Summing Digits: The sum of all processed digits is calculated.
Validation: The identification number is valid if the sum is divisible by 10.
*/
