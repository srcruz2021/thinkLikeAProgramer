package main

import "fmt"

// Node represents a node in a linked list storing a binary value (0 or 1)
type Node struct {
	value int
	next  *Node
}

// HasOddParityRecursive checks for odd parity using recursion
func HasOddParityRecursive(head *Node, count int) bool {
	// Base case: If we reach the end of the list
	if head == nil {
		return count%2 != 0 // True if count of 1s is odd
	}

	// Update count if current node's value is 1
	if head.value == 1 {
		count++
	}

	// Recurse to the next node
	return HasOddParityRecursive(head.next, count)
}

func main() {
	// Creating linked list: 1 → 0 → 1 → 1 → 0 → 1
	head := &Node{1, &Node{0, &Node{1, &Node{1, &Node{0, &Node{1, nil}}}}}}

	fmt.Println("Odd parity (recursive):", HasOddParityRecursive(head, 0)) // Output: true
}
