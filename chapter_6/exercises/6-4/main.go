package main

/*
Design your own: Find a problem processing a one-dimension array that you have already solved or that is trivial for you at
your current skill level, and solve the problem (or solve it again) using recursion.
*/

import "fmt"

// FindMaxRecursive finds the maximum value in an array using recursion
func FindMaxRecursive(arr []int, index int, currentMax int) int {
	// Base case: If we've checked all elements, return the max found
	if index >= len(arr) {
		return currentMax
	}

	// Update currentMax if the current element is larger
	if arr[index] > currentMax {
		currentMax = arr[index]
	}

	// Recursive call for the next element
	return FindMaxRecursive(arr, index+1, currentMax)
}

func main() {
	numbers := []int{12, 45, 7, 89, 23, 56}
	maxValue := FindMaxRecursive(numbers, 0, numbers[0])
	fmt.Println("Maximum value (recursive):", maxValue) // Output: 89
}
