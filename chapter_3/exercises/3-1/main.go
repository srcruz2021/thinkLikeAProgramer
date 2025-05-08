package main

import (
	"fmt"
	"sort"
)

/*
 write code that uses the function to sort an array of our student struct.
First have it sort by grade, and then try it again using the student ID.
*/

// Definimos la estructura Student
type Student struct {
	Grade int
	ID    int
	Name  string
}

func main() {
	// Definimos el arreglo de estudiantes
	students := []Student{
		{Grade: 87, ID: 10001, Name: "Fred"},
		{Grade: 28, ID: 10002, Name: "Tom"},
		{Grade: 100, ID: 10003, Name: "Alistair"},
		{Grade: 78, ID: 10004, Name: "Sasha"},
		{Grade: 84, ID: 10005, Name: "Erin"},
		{Grade: 98, ID: 10006, Name: "Belinda"},
		{Grade: 75, ID: 10007, Name: "Leslie"},
		{Grade: 70, ID: 10008, Name: "Candy"},
		{Grade: 81, ID: 10009, Name: "Aretha"},
		{Grade: 68, ID: 10010, Name: "Veronica"},
	}

	// Ordenamos por calificación (descendente)
	sort.Slice(students, func(i, j int) bool {
		return students[i].Grade > students[j].Grade
	})
	fmt.Println("Ordenado por calificación:")
	for _, student := range students {
		fmt.Println(student)
	}

	// Ordenamos por ID (ascendente)
	sort.Slice(students, func(i, j int) bool {
		return students[i].ID < students[j].ID
	})
	fmt.Println("\nOrdenado por ID:")
	for _, student := range students {
		fmt.Println(student)
	}
}
