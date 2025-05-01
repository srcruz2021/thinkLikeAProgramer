package main

import (
	"fmt"
	"strings"
)

/*
PROBLEM: HALF OF A SQUARE

Write a program that uses only two output statements, cout << "#" and cout << "\n", to produce a pattern
of hash symbols shaped like half of a perfect 5 x 5 square (or a right triangle):

#####
####
###
##
#
*/

/*
func main() {
	symbol := "#"
	maxValue := 5

	for i := 0; i < maxValue; i++ {
		fmt.Println(strings.Repeat(symbol, maxValue-i))
	}

}
*/

/*
PROBLEM: A SQUARE (HALF OF A SQUARE REDUCTION)

Write a program that uses only two output statements, cout << "#" and cout << "\n", to produce a pattern of hash symbols shaped like a perfect 5x5 square:

#####
#####
#####
#####
#####


func main() {
	symbol := "#"
	maxRepetitions := 5

	for i := 0; i < maxRepetitions; i++ {
		fmt.Println(strings.Repeat(symbol, maxRepetitions))
	}
}
*/

/*
PROBLEM: A LINE (HALF OF A SQUARE FURTHER REDUCTION)

Write a program that uses only two output statements, cout << "#" and cout << "\n",
to produce a line of five hash symbols:

#####


func main() {
	symbol := "#"
	maxValue := 5

	for i := 1; i <= 1; i++ {
		fmt.Println(strings.Repeat(symbol, maxValue))
	}

}
*/

/*
PROBLEM: COUNT DOWN BY COUNTING UP

Write a line of code that goes in the designated position in the loop in the listing below.
The program displays the numbers 5 through 1, in that order, with each number on a separate line.


func main() {
	val := 5
	for i := 0; i < val; i++ {
		fmt.Println(strconv.Itoa(val - i))
	}
}

*/

/*
 PROBLEM: A SIDEWAYS TRIANGLE

Write a program that uses only two output statements, cout << "#" and cout << "\n", to produce a pattern of hash symbols shaped like a sideways triangle:

#
##
###
####
###
##
#
*/

func main() {
	symbol := "#"
	repetitions := 8
	half := repetitions / 2

	for i := 1; i < repetitions; i++ {
		if i <= half {
			fmt.Println(strings.Repeat(symbol, i))
		}
		if i > half {
			fmt.Println(strings.Repeat(symbol, half-(i-half)))
		}
	}
}
