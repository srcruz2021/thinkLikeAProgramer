package main

import "fmt"

/*
Compute Statistics
*/

func main() {
	arrSize := 10
	grades := []int{87, 76, 100, 97, 64, 83, 88, 92, 74, 95}
	sum := 0

	for i := 0; i < arrSize; i++ {
		sum += grades[i]
	}

	average := float64(sum) / float64(arrSize)

	fmt.Println(fmt.Sprintf("Average: %.2f", average))
	fmt.Println(fmt.Sprintf("Median value: %v", grades[(arrSize/2)-1]))
	printHistogram()
}

func printHistogram() {
	// Example data
	data := []int{4, 7, 7, 9, 9, 9, 8, 3, 3, 3, 10}

	// Create a map to store frequencies
	frequency := make(map[int]int)
	for _, value := range data {
		frequency[value]++
	}

	// Display the histogram
	fmt.Println("Histogram:")
	for key, value := range frequency {
		fmt.Printf("%d: %s\n", key, string(make([]rune, value, '*')))
	}
}
