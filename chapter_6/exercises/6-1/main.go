package main

/*
Write a function to compute the sum of just the positive numbers in an array of integers.
First, solve the problem using iteration. Then, using the technique shown in this chapter,
convert your iterative function to a recursive function.
*/

import "fmt"

// SumPositiveIterative calculates the sum of positive numbers using iteration
func SumPositiveIterative(arr []int) int {
	sum := 0
	for _, num := range arr {
		if num > 0 {
			sum += num
		}
	}
	return sum
}

func main() {
	numbers := []int{-3, 5, -1, 8, -4, 10}
	fmt.Println("Sum of positive numbers (iterative):", SumPositiveIterative(numbers)) // Output: 23
}
