package main

import (
	"fmt"
)

/*
For the variable-length string class of the previous exercise,
replace the characterAt method with an overloaded [] operator. For example, if myString is an object of our class,
then myString[1] should return the same result as myString.characterAt(1)
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
}
