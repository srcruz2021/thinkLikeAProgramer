package main

import (
	"fmt"
	"strings"
)

/*
2-2-1. Or how about:

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
	lines := 8
	symbol := "#"
	increment := 2

	// Top of the figure
	for i := 0; i <= 4; i++ {
		// print blank space
		totalBlanks := lines - (increment * i)
		if totalBlanks >= 0 {
			fmt.Print(strings.Repeat(" ", totalBlanks))
			// print symbol
			fmt.Println(strings.Repeat(symbol, increment*i))
		}
	}

	// Bottom of the figure
	for i := 0; i <= 4; i++ {
		// print blank space
		totalBlanks := i * increment
		if totalBlanks >= 0 {
			fmt.Print(strings.Repeat(" ", totalBlanks))
			// print symbol
			fmt.Println(strings.Repeat(symbol, lines-(i*increment)))
		}
	}
}
