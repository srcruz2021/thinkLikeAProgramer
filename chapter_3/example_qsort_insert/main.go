package main

import "fmt"

/*
Easy to modify sorting with insertion sort

The insertion sort algorithm is simple yet effective for sorting small arrays. Conceptually, it works like organizing a hand of playing cards:

Start with the first element as a "sorted" section of the array.

Pick the next element (this is the "key") from the unsorted portion.

Compare the key with the elements in the sorted portion, moving elements larger than the key one position to the right.

Insert the key into the correct position in the sorted section.

Repeat these steps for every element in the array until the unsorted section is empty.

At the end of the process, all elements are in the correct order, and the array is sorted. It's like building the sorted sequence step by step!
*/

// InsertionSort function
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Move elements that are greater than the key one position ahead
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	arr := []int{34, 8, 64, 51, 32, 21}
	fmt.Println("Original array:", arr)
	sortedArr := InsertionSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
