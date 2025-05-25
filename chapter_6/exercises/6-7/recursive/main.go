package main

import "fmt"

// Recursive way

// Node represents a node in a linked list storing an integer value
type Node struct {
	value int
	next  *Node
}

// CountOccurrencesRecursive counts occurrences using recursion
func CountOccurrencesRecursive(head *Node, target int) int {
	// Base case: If we reach the end of the list
	if head == nil {
		return 0
	}

	// If the current node matches the target, add 1 to the result
	if head.value == target {
		return 1 + CountOccurrencesRecursive(head.next, target)
	}

	// Otherwise, continue checking the next node
	return CountOccurrencesRecursive(head.next, target)
}

func main() {
	// Creating linked list: 4 → 2 → 7 → 2 → 8 → 2 → 3 → 2
	head := &Node{4, &Node{2, &Node{7, &Node{2, &Node{8, &Node{2, &Node{3, &Node{2, nil}}}}}}}}

	target := 2
	fmt.Println("Occurrences (recursive):", CountOccurrencesRecursive(head, target)) // Output: 4
}
