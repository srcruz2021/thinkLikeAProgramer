package main

import "fmt"

// Node represents a node in a linked list storing a binary value (0 or 1)
type Node struct {
	value int
	next  *Node
}

// HasOddParityIterative checks for odd parity using iteration
func HasOddParityIterative(head *Node) bool {
	count := 0
	current := head

	// Traverse the linked list
	for current != nil {
		if current.value == 1 {
			count++
		}
		current = current.next
	}

	// True if count of 1s is odd, otherwise false
	return count%2 != 0
}

func main() {
	// Creating linked list: 1 → 0 → 1 → 1 → 0 → 1
	head := &Node{1, &Node{0, &Node{1, &Node{1, &Node{0, &Node{1, nil}}}}}}

	fmt.Println("Odd parity (iterative):", HasOddParityIterative(head)) // Output: true
}
