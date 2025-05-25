package main

import "fmt"

/*
For the digit list of the previous exercise, write a function that takes two such lists and produces a new list representing their sum.
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

// append agrega un nuevo dígito al final de la lista
func (ln *LinkedNumber) append(digit int) {
	newNode := &Node{digit: digit}
	if ln.head == nil {
		ln.head = newNode
	} else {
		current := ln.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// sumLists suma dos listas enlazadas y devuelve una nueva con el resultado
func sumLists(l1, l2 *LinkedNumber) *LinkedNumber {
	result := &LinkedNumber{}
	carry := 0

	n1, n2 := l1.head, l2.head

	for n1 != nil || n2 != nil || carry > 0 {
		sum := carry

		if n1 != nil {
			sum += n1.digit
			n1 = n1.next
		}
		if n2 != nil {
			sum += n2.digit
			n2 = n2.next
		}

		result.append(sum % 10) // Guardamos solo el dígito de las unidades
		carry = sum / 10        // Calculamos el acarreo
	}

	return result
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
	// Crear dos números como listas enlazadas
	l1 := &LinkedNumber{}
	l1.append(3)
	l1.append(4)
	l1.append(2) // Representa el número 342

	l2 := &LinkedNumber{}
	l2.append(4)
	l2.append(6)
	l2.append(5) // Representa el número 465

	fmt.Print("Número 1: ")
	l1.printList()

	fmt.Print("Número 2: ")
	l2.printList()

	// Sumar las listas
	sum := sumLists(l1, l2)

	fmt.Print("Suma: ")
	sum.printList() // Output esperado: 807
}
