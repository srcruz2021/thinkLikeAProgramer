package main

import (
	"fmt"
)

/*
help me in go Reduce a string of lowercase characters in range ascii[‘a’..’z’]by doing a series of operations.
In each operation, select a pair of adjacent letters that match, and delete them.
Delete as many characters as possible using this method and return the resulting string. If the final string is empty,
return Empty String Example. s = "aab" aab shortens to b in one operation: remove the adjacent a characters. s = "abba" Remove the two 'b' characters leaving 'aa'.
Remove the two 'a' characters to leave ''. Return 'Empty String'. Create a Function Description Complete the superReducedString function below.
superReducedString has the following parameter(s): string s: a string to reduce Returns string:
the reduced string or Empty String
Input Format
A single string, s
Constraints 1 <= length s <= 100
*/

func superReducedString(s string) string {
	stack := []rune{}

	for _, char := range s {
		if len(stack) > 0 && stack[len(stack)-1] == char {
			stack = stack[:len(stack)-1] // Remove the matching adjacent character
		} else {
			stack = append(stack, char) // Add non-matching character to the stack
		}
	}

	if len(stack) == 0 {
		return "Empty String"
	}

	return string(stack)
}

func main() {
	s := "abba"
	fmt.Println(superReducedString(s)) // Output: "Empty String"
}
