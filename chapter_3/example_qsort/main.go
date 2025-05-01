package main

import "fmt"

/*
How it works:
The QuickSort function starts by choosing a pivot element (in this case, the first element of the array).

It divides the array into two parts:

Left: Elements smaller than or equal to the pivot.

Right: Elements larger than the pivot.

The function is then called recursively on the left and right slices until they are fully sorted.

Finally, it combines the sorted left slice, the pivot, and the sorted right slice to form the fully sorted array.
*/

// QuickSort function
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0] // Choosing the pivot
	var left, right []int

	// Partitioning into left and right slices
	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// Recursively sorting left and right slices
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

func main() {
	arr := []int{34, 7, 23, 32, 5, 62}
	fmt.Println("Original array:", arr)
	sortedArr := QuickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

/*
How it works:
Recursive Sorting of Left:

QuickSort(left) sorts the left slice recursively using the QuickSort algorithm.

Appending the Pivot:

append(QuickSort(left), pivot) takes the sorted left slice and appends the pivot element to it. Now you have a new slice containing the sorted left elements followed by the pivot.

Recursive Sorting of Right:

QuickSort(right) sorts the right slice recursively.

The ... at the end is the variadic argument operator, which "unpacks" the elements of the sorted right slice so they can be appended individually.

Final Append:

append(... QuickSort(right)) appends the sorted right elements to the previously created slice of sorted left and pivot.

The result is a fully sorted array combining the sorted left slice, pivot, and sorted right slice.

Key takeaway:
This line combines all pieces (sorted left, pivot, and sorted right) into one single, sorted array. It’s efficient and makes use of the recursive nature of QuickSort.
*/
