package main

import (
	"fmt"
)

/*
Take the variable-length string functions from Chapter 4 (append, concatenate, and characterAt) and use them to create a class for
variable-length strings, making sure to implement all necessary constructors, a destructor, and an overloaded assignment operator
*/

// VarString representa una cadena de longitud variable
type VarString struct {
	data []rune
}

// NewVarString crea una nueva instancia de VarString
func NewVarString(initial string) *VarString {
	return &VarString{data: []rune(initial)}
}

// Append agrega un nuevo carácter al final de la cadena
func (vs *VarString) Append(ch rune) {
	vs.data = append(vs.data, ch)
}

// Concatenate agrega otra VarString al final de la actual
func (vs *VarString) Concatenate(other *VarString) {
	vs.data = append(vs.data, other.data...)
}

// CharacterAt devuelve el carácter en una posición específica
func (vs *VarString) CharacterAt(index int) (rune, bool) {
	if index < 0 || index >= len(vs.data) {
		return '\000', false // Retorna un carácter nulo si el índice es inválido
	}
	return vs.data[index], true
}

// Clear libera la memoria de la cadena (similar a un destructor)
func (vs *VarString) Clear() {
	vs.data = nil
}

// Print muestra la cadena almacenada
func (vs *VarString) Print() {
	fmt.Println(string(vs.data))
}

func main() {
	// Crear una cadena inicial
	str1 := NewVarString("Hello")
	str1.Print() // Output: Hello

	// Agregar un carácter
	str1.Append('!')
	str1.Print() // Output: Hello!

	// Crear otra cadena y concatenarla
	str2 := NewVarString(" World")
	str1.Concatenate(str2)
	str1.Print() // Output: Hello! World

	// Obtener un carácter en una posición específica
	if ch, ok := str1.CharacterAt(6); ok {
		fmt.Println("Carácter en posición 6:", string(ch)) // Output: W
	} else {
		fmt.Println("Índice fuera de rango.")
	}

	// Limpiar la memoria de la cadena
	str1.Clear()
	fmt.Println("Después de limpiar:", str1.data) // Output: []

}
