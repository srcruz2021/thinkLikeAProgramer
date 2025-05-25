package main

import "fmt"

/*
 Add a function to the linked-list string implementation called removeChars to remove a section of characters from
a string based on the position and length. For example, removeChars(s1, 5, 3) would remove the three characters starting at
the fifth character in the string. Make sure the removed nodes are properly deallocated.
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

// removeChars elimina caracteres desde una posición dada con una longitud específica
func (ls *LinkedString) removeChars(position, length int) {
	if position < 0 || position >= ls.size || length <= 0 {
		fmt.Println("Posición o longitud inválida.")
		return
	}

	// Localizar el nodo antes de la sección a eliminar
	var prev *Node
	current := ls.head
	for i := 0; i < position; i++ {
		prev = current
		current = current.next
	}

	// Eliminar los nodos deseados
	for i := 0; i < length && current != nil; i++ {
		temp := current
		current = current.next
		temp.next = nil // Liberamos la referencia del nodo eliminado
	}

	// Conectar el nodo anterior con el siguiente
	if prev != nil {
		prev.next = current
	} else {
		ls.head = current // Si estamos eliminando desde el inicio
	}

	// Ajustar el tamaño de la cadena
	ls.size -= length
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
	// Crear y poblar una lista enlazada
	ls := &LinkedString{}
	text := "Hello, world!"
	for _, ch := range text {
		ls.append(ch)
	}

	// Imprimir cadena original
	fmt.Println("Original:")
	ls.printString() // Output: Hello, world!

	// Eliminar caracteres desde la posición 5, eliminando 3 caracteres
	ls.removeChars(5, 3)

	// Imprimir cadena modificada
	fmt.Println("Después de eliminar:")
	ls.printString() // Output: Hello world!
}
