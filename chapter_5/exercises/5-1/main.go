package main

import "fmt"

/*
Let’s try implementing a class using the basic framework. Consider a class to store the data for an automobile.
We’ll have three pieces of data: a manufacturer name and model name, both strings, and a model year, an integer.
Create a class with get/set methods for each data member. Make sure you make good decisions concerning details like member names.
It’s not important that you follow my particular naming convention. What’s important is that you think about the choices you make and
are consistent in your decisions
*/

// Automobile representa un automóvil con fabricante, modelo y año
type Automobile struct {
	manufacturer string
	model        string
	year         int
}

// NewAutomobile crea una nueva instancia de Automobile
func NewAutomobile(manufacturer, model string, year int) *Automobile {
	return &Automobile{manufacturer: manufacturer, model: model, year: year}
}

// GetManufacturer devuelve el nombre del fabricante
func (a *Automobile) GetManufacturer() string {
	return a.manufacturer
}

// SetManufacturer establece el nombre del fabricante
func (a *Automobile) SetManufacturer(manufacturer string) {
	a.manufacturer = manufacturer
}

// GetModel devuelve el nombre del modelo
func (a *Automobile) GetModel() string {
	return a.model
}

// SetModel establece el nombre del modelo
func (a *Automobile) SetModel(model string) {
	a.model = model
}

// GetYear devuelve el año del modelo
func (a *Automobile) GetYear() int {
	return a.year
}

// SetYear establece el año del modelo
func (a *Automobile) SetYear(year int) {
	a.year = year
}

// PrintInfo muestra la información del automóvil
func (a *Automobile) PrintInfo() {
	fmt.Printf("Automóvil: %s %s (%d)\n", a.manufacturer, a.model, a.year)
}

func main() {
	// Crear un automóvil
	car := NewAutomobile("Toyota", "Corolla", 2022)

	// Mostrar información inicial
	car.PrintInfo()

	// Modificar atributos
	car.SetManufacturer("Honda")
	car.SetModel("Civic")
	car.SetYear(2023)

	// Mostrar información actualizada
	car.PrintInfo()
}
