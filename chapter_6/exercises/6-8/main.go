package main

/*
Design your own: Try to discover a linked-list processing problem that is difficult to solve using
iteration but can be solved directly using recursion.

A problem that is difficult to solve using iteration but simpler with recursion is reversing a linked list in place.

🔹 Why is Iteration Difficult?
In an iterative approach, you'd need to track three pointers (prev, current, and next).

You must manually update links, which can be error-prone.

You have to loop through the entire list while adjusting pointers carefully.

🔹 Why is Recursion Easier?
Recursion naturally unfolds the list, processing each node before assembling it in reverse.

No need for extra tracking variables like prev and next.

The function handles pointer updates automatically as it unwinds.
*/

import "fmt"

// Node represents a linked list node
type Node struct {
	value int
	next  *Node
}

// ReverseLinkedListRecursive reverses a linked list using recursion
func ReverseLinkedListRecursive(head *Node) *Node {
	// Base case: If we reach the end (or empty list)
	if head == nil || head.next == nil {
		return head
	}

	// Recursive call to reverse the rest of the list
	newHead := ReverseLinkedListRecursive(head.next)

	// Adjust pointers to reverse the current node
	head.next.next = head
	head.next = nil

	return newHead
}

// PrintList prints a linked list
func PrintList(head *Node) {
	for head != nil {
		fmt.Print(head.value, " → ")
		head = head.next
	}
	fmt.Println("nil")
}

func main() {
	// Creating linked list: 1 → 2 → 3 → 4 → nil
	head := &Node{1, &Node{2, &Node{3, &Node{4, nil}}}}

	fmt.Println("Original list:")
	PrintList(head)

	// Reverse the linked list recursively
	head = ReverseLinkedListRecursive(head)

	fmt.Println("Reversed list:")
	PrintList(head)
}
