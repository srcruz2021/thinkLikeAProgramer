package main

/*
Consider an array representing a binary string, where every element’s data value is 0 or 1.
Write a bool function to determine whether the binary string has odd parity (an odd number of 1 bits).
Hint: Remember that the recursive function is going to return true (odd) or false (even), not the count of 1 bits.
Solve the problem first using iteration, then recursion.
*/

import "fmt"

// HasOddParityIterative checks if a binary string has odd parity using iteration
func HasOddParityIterative(arr []int) bool {
	count := 0
	for _, bit := range arr {
		if bit == 1 {
			count++
		}
	}
	return count%2 != 0 // True if count of 1s is odd
}

func main() {
	binaryString := []int{1, 0, 1, 1, 0, 1}                                     // 4 ones → even parity (false)
	fmt.Println("Odd parity (iterative):", HasOddParityIterative(binaryString)) // Output: false
}
