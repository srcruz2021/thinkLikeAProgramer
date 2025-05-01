package main

import (
	"fmt"
	"math"
)

/*
func main() {

		var matrix [3][3]int32
		matrix[0][0] = 1
		matrix[0][1] = 2
		matrix[0][2] = 3
		matrix[1][0] = 4
		matrix[1][1] = 5
		matrix[1][2] = 6
		matrix[2][0] = 9
		matrix[2][1] = 8
		matrix[2][2] = 9

		absoluteValue := int32(0)
		sumRightDiagonal := matrix[0][2] + matrix[1][1] + matrix[2][0]
		sumLeftDiagonal := matrix[0][0] + matrix[1][1] + matrix[2][2]
		absoluteValue = sumRightDiagonal - sumLeftDiagonal


	// Sample input
	rows, cols := int32(3), int32(3) // Assuming the first line gave us "3 3"
	matrix := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	// Ensure rows == cols for a square matrix
	if rows != cols {
		fmt.Println("Error: Matrix is not square")
		return
	}

	// Variables to hold diagonal sums
	primaryDiagonalSum := int32(0)
	secondaryDiagonalSum := int32(0)

	// Iterate through the matrix to compute diagonal sums
	for i := int32(0); i < rows; i++ {
		primaryDiagonalSum += matrix[i][i]          // Primary diagonal
		secondaryDiagonalSum += matrix[i][cols-i-1] // Secondary diagonal
	}

	// Calculate absolute difference
	diagonalDifference := int32(math.Abs(float64(primaryDiagonalSum - secondaryDiagonalSum)))

	// Print the result
	fmt.Println("Diagonal Difference:", diagonalDifference)

	fmt.Println(fmt.Sprintf("leftDiagonalSum is: %d \nrightDiagonalSum is: %d \nAbsolute Value: %d", primaryDiagonalSum, secondaryDiagonalSum, diagonalDifference))

}
*/

func diagonalDifference(arr [][]int32) int32 {
	n := int32(len(arr)) // Determine the size of the square matrix

	// Initialize sums for diagonals
	primaryDiagonalSum := int32(0)
	secondaryDiagonalSum := int32(0)

	// Calculate sums of the diagonals
	for i := int32(0); i < n; i++ {
		primaryDiagonalSum += arr[i][i]       // Primary diagonal
		secondaryDiagonalSum += arr[i][n-i-1] // Secondary diagonal
	}

	// Return the absolute difference between the diagonal sums
	return int32(math.Abs(float64(primaryDiagonalSum - secondaryDiagonalSum)))
}

func main() {
	// Example input for testing
	matrix := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	// Call the function and print the result
	result := diagonalDifference(matrix)
	fmt.Println("Absolute difference between the sums of diagonals:", result)
}
