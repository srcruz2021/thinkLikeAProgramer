package main

import "fmt"

/*
For the variable-length string class of the previous exercises,
add a remove method that takes a starting position and a number of characters and removes that many characters from the middle of
the string. So myString.remove(5,3) would remove three characters starting at the fifth position.
Make sure your method behaves when the value of either of the parameters is invalid.
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

// At devuelve el carácter en una posición específica (simulando sobrecarga de [])
func (vs *VarString) At(index int) (rune, bool) {
	if index < 0 || index >= len(vs.data) {
		return '\000', false // Retorna un carácter nulo si el índice es inválido
	}
	return vs.data[index], true
}

// Clear libera la memoria de la cadena
func (vs *VarString) Clear() {
	vs.data = nil
}

// Print muestra la cadena almacenada
func (vs *VarString) Print() {
	fmt.Println(string(vs.data))
}

func main() {
	// Crear una cadena inicial
	str := NewVarString("Hello")
	str.Print() // Output: Hello

	// Acceder a un carácter usando At (simulando sobrecarga de [])
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

	strRemoved := NewVarString("Hello, World!")
	strRemoved.Remove(5, 3)
	strRemoved.Print()
}

// Remove elimina una cantidad específica de caracteres desde una posición dada
func (vs *VarString) Remove(start, count int) {
	if start < 0 || start >= len(vs.data) || count < 0 {
		return // No hacer nada si los parámetros son inválidos
	}

	end := start + count
	if end > len(vs.data) {
		end = len(vs.data) // Asegurar que no exceda el tamaño del slice
	}

	vs.data = append(vs.data[:start], vs.data[end:]...)
}
