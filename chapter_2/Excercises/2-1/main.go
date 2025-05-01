package main

import (
	"fmt"
	"strings"
)

/*
	using only single-character output statement that output a hash mark a space, or
an end of line, write a program that produces the following shape:

########
 ######
  ####
   ##
*/

func main() {
	symbol := "#"
	repetitions := 8
	removeElements := 2
	lines := 4

	for i := 0; i < lines; i++ {
		// Print leading spaces
		for j := 0; j < i; j++ {
			fmt.Print(" ") // Single-character output for space
		}
		// print symbol
		fmt.Print(strings.Repeat(symbol, repetitions-(i*removeElements)))
		// print new line
		fmt.Print("\n")
	}

}
