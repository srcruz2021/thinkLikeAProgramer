package main

import (
	"fmt"
	"time"
)

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

// GetDescription devuelve una descripción completa del automóvil
func (a *Automobile) GetDescription() string {
	return fmt.Sprintf("%d %s %s", a.year, a.manufacturer, a.model)
}

// GetAge devuelve la edad del automóvil en años
func (a *Automobile) GetAge() int {
	currentYear := time.Now().Year()
	return currentYear - a.year
}

// PrintInfo muestra la información del automóvil
func (a *Automobile) PrintInfo() {
	fmt.Println("Descripción:", a.GetDescription())
	fmt.Println("Edad:", a.GetAge(), "años")
}

func main() {
	// Crear un automóvil
	car := NewAutomobile("Chevrolet", "Impala", 1957)

	// Mostrar información del automóvil
	car.PrintInfo()
}
