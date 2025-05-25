package main

import "fmt"

/*
Following up on the previous exercise, implement the concatenate function.
Note that if we make a call concatenate(s1, s2), where both parameters are pointers to the first nodes of their respective linked lists,
the function should create a copy of each of the nodes in s2 and append them to the end of s1.
That is, the function should not simply point the next field of the last node in s1’s list to the first node of s2’s list
*/
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

// concatenate copia los nodos de s2 y los agrega al final de s1
func concatenate(s1, s2 *LinkedString) {
	if s2.head == nil {
		return
	}

	current := s1.head
	// Encontrar el último nodo de s1
	for current != nil && current.next != nil {
		current = current.next
	}

	// Copiar los nodos de s2 y agregarlos al final de s1
	s2Current := s2.head
	for s2Current != nil {
		newNode := &Node{char: s2Current.char}
		if current == nil {
			s1.head = newNode
			current = s1.head
		} else {
			current.next = newNode
			current = current.next
		}
		s1.size++
		s2Current = s2Current.next
	}
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
	// Crear primera lista enlazada
	s1 := &LinkedString{}
	s1.append('H')
	s1.append('e')
	s1.append('l')
	s1.append('l')
	s1.append('o')

	// Crear segunda lista enlazada
	s2 := &LinkedString{}
	s2.append(' ')
	s2.append('W')
	s2.append('o')
	s2.append('r')
	s2.append('l')
	s2.append('d')

	// Concatenar s2 en s1 sin modificar s2
	concatenate(s1, s2)

	// Imprimir la nueva cadena en s1
	s1.printString() // Output: Hello World

	// Imprimir s2 para verificar que no ha cambiado
	s2.printString() // Output: World
}
