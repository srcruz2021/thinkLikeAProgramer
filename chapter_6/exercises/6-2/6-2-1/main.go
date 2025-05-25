package main

import "fmt"

// HasOddParityRecursive checks odd parity using recursion
func HasOddParityRecursive(arr []int, index int, count int) bool {
	// Base case: When we reach the end of the array
	if index >= len(arr) {
		return count%2 != 0 // True if count is odd
	}

	// Update the count if the current element is 1
	if arr[index] == 1 {
		count++
	}

	// Recurse to the next element
	return HasOddParityRecursive(arr, index+1, count)
}

func main() {
	binaryString := []int{1, 0, 1, 1, 0, 1}                                           // 4 ones → even parity (false)
	fmt.Println("Odd parity (recursive):", HasOddParityRecursive(binaryString, 0, 0)) // Output: false
}
