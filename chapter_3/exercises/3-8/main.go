package main

import (
	"fmt"
	"sort"
)

/*
Write a program that processes an array of student objects and determines the grade quartiles—that is,
the grade one would need to score as well as or better than 25% of the students, 50% of the students, and 75% of the students.
*/
// Definimos la estructura Student
type Student struct {
	Name  string
	Grade float64
}

// Función para calcular los cuartiles
func quartile(grades []float64, percentile float64) float64 {
	index := int(percentile * float64(len(grades)))
	if index >= len(grades) {
		index = len(grades) - 1
	}
	return grades[index]
}

func main() {
	// Lista de estudiantes con sus calificaciones
	students := []Student{
		{"Ana", 85.5}, {"Luis", 92.3}, {"Carlos", 78.9},
		{"Sofia", 88.0}, {"Miguel", 74.5}, {"Elena", 95.2},
		{"Javier", 81.7}, {"Paula", 89.3}, {"Raúl", 69.8},
		{"Teresa", 76.2},
	}

	// Extraemos las calificaciones y las ordenamos
	var grades []float64
	for _, student := range students {
		grades = append(grades, student.Grade)
	}
	sort.Float64s(grades) // Ordenamos las calificaciones

	// Calculamos los cuartiles
	q1 := quartile(grades, 0.25) // Primer cuartil (25%)
	q2 := quartile(grades, 0.50) // Segundo cuartil (mediana, 50%)
	q3 := quartile(grades, 0.75) // Tercer cuartil (75%)

	// Mostramos los resultados
	fmt.Printf("Primer cuartil (25%%): %.2f\n", q1)
	fmt.Printf("Segundo cuartil (50%% - Mediana): %.2f\n", q2)
	fmt.Printf("Tercer cuartil (75%%): %.2f\n", q3)
}
