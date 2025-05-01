package main

import (
	"fmt"
	"sort"
)

func main() {
	// Example array of five positive integers
	arr := []int32{100, 500, 45600, 1, 8}

	// Validate constraints: all elements must be positive and within [1, 1,000,000,000]
	for _, val := range arr {
		if val < 1 || val > 1000000000 {
			fmt.Println("Error: Array contains invalid value:", val)
			return
		}
	}

	// Sort the array in ascending order
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Calculate the minimum sum (sum of the first 4 elements)
	var minSum int64
	for i := 0; i < 4; i++ {
		minSum += int64(arr[i]) // Convert to int64 to handle larger sums
	}

	// Calculate the maximum sum (sum of the last 4 elements)
	var maxSum int64
	for i := 1; i < 5; i++ {
		maxSum += int64(arr[i]) // Convert to int64 to handle larger sums
	}

	// Print the results as a single line of space-separated values
	fmt.Printf("%d %d\n", minSum, maxSum)
}
