package main

import (
	"fmt"
	"strings"
)

/*
	using only single-character output statement that output a hash mark a space, or
an end of line, write a program that produces the following shape:
   ##
  ####
 ######
########
########
 ######
  ####
   ##
*/

func main() {
	symbol := "#"
	finalRepetitions := 8
	initialRepetitions := 2
	removeElements := 2
	lines := 4

	// Top of pyramid
	for i := 0; i < lines; i++ {
		for j := lines; i < j; j-- {
			fmt.Print(" ") // Single-character output for space
		}
		// print symbol
		fmt.Print(strings.Repeat(symbol, initialRepetitions*i))
		// print white space
		fmt.Print("\n")
	}

	// bottom of pyramid
	for i := 0; i < lines; i++ {
		// Print leading spaces
		for j := 0; j < i; j++ {
			fmt.Print(" ") // Single-character output for space
		}
		// print symbol
		fmt.Print(strings.Repeat(symbol, finalRepetitions-(i*removeElements)))
		// print new line
		fmt.Print("\n")
	}

}
