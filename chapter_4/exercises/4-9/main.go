package main

import (
	"fmt"
)

/*
 Imagine a linked list where instead of the node storing a character, the node stores a digit: an int in the range 0–9.
We could represent positive numbers of any size using such a linked list; the number 149, for example, would be a linked list in which
the first node stores a 1, the second a 4, and the third and last a 9. Write a function intToList that takes an integer value and
produces a linked list of this sort. Hint: You may find it easier to build the linked list backward, so if the value were 149,
you would create the 9 node first.
*/
// Node representa un nodo en la lista enlazada de dígitos
type Node struct {
	digit int
	next  *Node
}

// LinkedNumber representa un número como una lista enlazada
type LinkedNumber struct {
	head *Node
}

// intToList convierte un entero en una lista enlazada de dígitos
func intToList(value int) *LinkedNumber {
	if value < 0 {
		fmt.Println("Solo números positivos son soportados.")
		return nil
	}

	ln := &LinkedNumber{}

	// Manejo especial para el caso de valor 0
	if value == 0 {
		ln.head = &Node{digit: 0}
		return ln
	}

	var prev *Node
	for value > 0 {
		digit := value % 10
		newNode := &Node{digit: digit, next: prev}
		prev = newNode
		value /= 10
	}

	ln.head = prev
	return ln
}

// printList imprime el número representado por la lista enlazada
func (ln *LinkedNumber) printList() {
	current := ln.head
	for current != nil {
		fmt.Print(current.digit)
		current = current.next
	}
	fmt.Println()
}

func main() {
	number := 149
	ln := intToList(number)

	fmt.Print("Número representado como lista enlazada: ")
	ln.printList() // Output esperado: 149
}
