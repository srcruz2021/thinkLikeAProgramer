package main

import "fmt"

/*
Write a bool function that is passed an array and the number of elements in that array
and determines whether the data in the array is sorted. This should require only one pass
*/
// isSorted verifica si un arreglo está ordenado en orden ascendente
func isSorted(arr []int, n int) bool {
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	n := len(arr)

	if isSorted(arr, n) {
		fmt.Println("El arreglo está ordenado.")
	} else {
		fmt.Println("El arreglo no está ordenado.")
	}
}
