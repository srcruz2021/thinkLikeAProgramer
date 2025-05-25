package main

/*
Review your variable-length string class for possible refactoring.
For example, is there any common functionality that can be separated into a private support method?
*/

import "fmt"

// VarString representa una cadena de longitud variable
type VarString struct {
	data []rune
}

// NewVarString crea una nueva instancia de VarString
func NewVarString(initial string) *VarString {
	return &VarString{data: []rune(initial)}
}

// validateIndex asegura que el índice está dentro del rango válido
func (vs *VarString) validateIndex(index int) bool {
	return index >= 0 && index < len(vs.data)
}

// updateData actualiza la estructura de datos internamente
func (vs *VarString) updateData(newData []rune) {
	vs.data = newData
}

// Append agrega un nuevo carácter al final de la cadena
func (vs *VarString) Append(ch rune) {
	vs.updateData(append(vs.data, ch))
}

// Concatenate agrega otra VarString al final de la actual
func (vs *VarString) Concatenate(other *VarString) {
	vs.updateData(append(vs.data, other.data...))
}

// At devuelve el carácter en una posición específica (simulando sobrecarga de [])
func (vs *VarString) At(index int) (rune, bool) {
	if !vs.validateIndex(index) {
		return '\000', false
	}
	return vs.data[index], true
}

// Remove elimina una cantidad específica de caracteres desde una posición dada
func (vs *VarString) Remove(start, count int) {
	if !vs.validateIndex(start) || count < 0 {
		return
	}

	end := start + count
	if end > len(vs.data) {
		end = len(vs.data) // Asegurar que no exceda el tamaño del slice
	}

	vs.updateData(append(vs.data[:start], vs.data[end:]...))
}

// Clear libera la memoria de la cadena
func (vs *VarString) Clear() {
	vs.updateData([]rune{})
}

// Print muestra la cadena almacenada
func (vs *VarString) Print() {
	fmt.Println(string(vs.data))
}

func main() {
	// Crear una cadena inicial
	str := NewVarString("Hello, World!")
	str.Print() // Output: Hello, World!

	// Remover tres caracteres desde la posición 5
	str.Remove(5, 3)
	str.Print() // Output: HelloWorld!

	// Acceder a un carácter usando At
	if ch, ok := str.At(1); ok {
		fmt.Println("Carácter en posición 1:", string(ch)) // Output: e
	} else {
		fmt.Println("Índice fuera de rango.")
	}

	// Intentar acceder a un índice inválido
	if ch, ok := str.At(10); ok {
		fmt.Println("Carácter en posición 10:", string(ch))
	} else {
		fmt.Println("Índice fuera de rango.") // Output esperado
	}
}
