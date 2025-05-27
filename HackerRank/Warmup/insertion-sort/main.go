package main

import (
	"fmt"
)

/*
Sorting One common task for computers is to sort data.
For example, people might want to see all their files on a computer sorted by size.
Since sorting is a simple problem with many different possible solutions, it is often used to introduce the study of algorithms.
Insertion Sort These challenges will cover Insertion Sort, a simple and intuitive sorting algorithm.
We will first start with a nearly sorted list. Insert element into sorted list Given a sorted list with an unsorted number in the rightmost cell,
can you write some simple code to insert into the array so that it remains sorted? Since this is a learning exercise,
it won't be the most efficient way of performing the insertion. It will instead demonstrate the brute-force method in detail.
Function Description Complete the insertionSort1 function below. insertionSort1 has the following parameter(s): n: an integer,
the size of arr arr: an array of integers to sort Returns None: Print the interim and final arrays, each on a new line.
No return value is expected. Input Format The first line contains the integer , the size of the array arr The next line contains n space-separated integers arr[0]...arr[n-1] .
Constraint: 1 <= n <=1000 - 1000 <= arr[i] <= 10000 Output Format Print the array as a row of space-separated integers each time there is a shift or insertion.
input:
5
2 4 6 8 3
Output
2 4 6 8 8
2 4 6 6 8
2 4 4 6 8
2 3 4 6 8
*/

// insertionSort1 inserts the last element into its correct position in a sorted array
func insertionSort1(n int32, arr []int32) {
	// The last element to be inserted
	key := arr[n-1]

	// Traverse backwards from the second last element
	i := n - 2
	for i >= 0 && arr[i] > key {
		arr[i+1] = arr[i] // Shift element to the right
		printArray(arr)   // Print current state
		i--
	}

	// Place the key in its correct position
	arr[i+1] = key
	printArray(arr) // Print final state
}

// Helper function to print the array
func printArray(arr []int32) {
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func main() {
	// Example input
	var n int32 = 5
	arr := []int32{2, 4, 6, 8, 3}

	// Perform insertion sort for the last element
	insertionSort1(n, arr)
}
