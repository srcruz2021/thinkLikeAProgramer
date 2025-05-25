package main

import "fmt"

// iterative approach

// Node represents a node in a linked list storing an integer value
type Node struct {
	value int
	next  *Node
}

// CountOccurrencesIterative counts occurrences of the target using iteration
func CountOccurrencesIterative(head *Node, target int) int {
	count := 0
	current := head

	// Traverse the linked list
	for current != nil {
		if current.value == target {
			count++
		}
		current = current.next
	}

	return count
}

func main() {
	// Creating linked list: 4 → 2 → 7 → 2 → 8 → 2 → 3 → 2
	head := &Node{4, &Node{2, &Node{7, &Node{2, &Node{8, &Node{2, &Node{3, &Node{2, nil}}}}}}}}

	target := 2
	fmt.Println("Occurrences (iterative):", CountOccurrencesIterative(head, target)) // Output: 4
}
