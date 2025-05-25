package main

/*
Write a function that is passed an array of integers and a “target” number and that returns the number of occurrences of the target
in the array. Solve the problem first using iteration, then recursion.
*/

import "fmt"

// CountOccurrencesIterative counts occurrences of a target number using iteration
func CountOccurrencesIterative(arr []int, target int) int {
	count := 0
	for _, num := range arr {
		if num == target {
			count++
		}
	}
	return count
}

func main() {
	numbers := []int{4, 2, 7, 2, 8, 2, 3, 2}
	target := 2
	fmt.Println("Occurrences (iterative):", CountOccurrencesIterative(numbers, target)) // Output: 4
}
