package main

/*
Let’s create an implementation for strings that uses a linked list of characters instead of dynamically allocated arrays.
So we’ll have a linked list where the data payload is a single char; this will allow strings to grow without having to recreate
the entire string. We’ll start by implementing the append and characterAt functions.

*/

import "fmt"

// Node representa un nodo en la lista enlazada de caracteres
type Node struct {
	char rune
	next *Node
}

// LinkedString representa una cadena implementada con una lista enlazada
type LinkedString struct {
	head *Node
	size int
}

// append agrega un nuevo carácter al final de la lista
func (ls *LinkedString) append(ch rune) {
	newNode := &Node{char: ch}
	if ls.head == nil {
		ls.head = newNode
	} else {
		current := ls.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ls.size++
}

// characterAt devuelve el carácter en una posición específica
func (ls *LinkedString) characterAt(index int) (rune, bool) {
	if index < 0 || index >= ls.size {
		return '\000', false // Retorna un carácter nulo si el índice es inválido
	}
	current := ls.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.char, true
}

// printString imprime la cadena enlazada
func (ls *LinkedString) printString() {
	current := ls.head
	for current != nil {
		fmt.Print(string(current.char))
		current = current.next
	}
	fmt.Println()
}

func main() {
	ls := &LinkedString{}

	// Agregar caracteres
	ls.append('H')
	ls.append('e')
	ls.append('l')
	ls.append('l')
	ls.append('o')

	// Imprimir la cadena
	ls.printString() // Output: Hello

	// Obtener un carácter en una posición específica
	if ch, ok := ls.characterAt(2); ok {
		fmt.Println("Carácter en posición 2:", string(ch)) // Output: l
	} else {
		fmt.Println("Índice fuera de rango.")
	}
}
