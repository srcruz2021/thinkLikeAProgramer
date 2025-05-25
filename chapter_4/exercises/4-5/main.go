package main

import "fmt"

/*
Write a function removeRecord that takes a pointer to a studentCollection and a student number and that
removes the record with that student number from the collection.
*/

// Student representa un registro de estudiante
type Student struct {
	Number int
	Name   string
	Age    int
}

// StudentCollection representa una colección de estudiantes
type StudentCollection struct {
	Records []Student
}

// removeRecord elimina el registro con el número de estudiante especificado
func removeRecord(collection *StudentCollection, studentNumber int) {
	for i, student := range collection.Records {
		if student.Number == studentNumber {
			// Eliminamos el registro reordenando los elementos
			collection.Records = append(collection.Records[:i], collection.Records[i+1:]...)
			fmt.Println("Registro eliminado:", student)
			return
		}
	}
	fmt.Println("Estudiante con número", studentNumber, "no encontrado.")
}

func main() {
	// Inicializamos una colección de estudiantes
	collection := StudentCollection{
		Records: []Student{
			{Number: 101, Name: "Juan", Age: 20},
			{Number: 102, Name: "María", Age: 22},
			{Number: 103, Name: "Carlos", Age: 19},
		},
	}

	// Mostramos la colección original
	fmt.Println("Colección inicial:", collection.Records)

	// Eliminamos un estudiante
	removeRecord(&collection, 102)

	// Mostramos la colección después de la eliminación
	fmt.Println("Colección actualizada:", collection.Records)
}
