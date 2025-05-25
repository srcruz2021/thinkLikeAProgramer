package main

/*
Solve exercise 6-1 again, using a linked list instead of an array.
*/

import "fmt"

// Node represents a node in the linked list
type Node struct {
	value int
	next  *Node
}

// SumPositiveRecursive computes the sum of positive numbers recursively
func SumPositiveRecursive(head *Node) int {
	// Base case: If head is nil, return 0
	if head == nil {
		return 0
	}

	// Include head's value if it's positive, then recurse to the next node
	if head.value > 0 {
		return head.value + SumPositiveRecursive(head.next)
	}

	// Otherwise, skip this node and recurse
	return SumPositiveRecursive(head.next)
}

func main() {
	// Creating linked list: -3 → 5 → -1 → 8 → -4 → 10
	head := &Node{-3, &Node{5, &Node{-1, &Node{8, &Node{-4, &Node{10, nil}}}}}}

	fmt.Println("Sum of positive numbers (recursive, linked list):", SumPositiveRecursive(head)) // Output: 23
}
