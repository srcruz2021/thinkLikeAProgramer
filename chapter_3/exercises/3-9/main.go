package main

import (
	"fmt"
	"sort"
)

/*
Consider this modification of the sales array: Because salespeople come and go throughout the year,
we are now marking months prior to a sales agent’s hiring, or after a sales agent’s last month, with a –1.
Rewrite your highest sales average, or highest sales median, code to compensate.
*/

// Definimos la cantidad de agentes y meses
const NUM_AGENTS = 3
const NUM_MONTHS = 12

// Función para calcular la mediana de un conjunto de valores, ignorando -1
func median(values []int) float64 {
	validValues := []int{}

	// Filtramos los valores válidos (excluimos -1)
	for _, v := range values {
		if v != -1 {
			validValues = append(validValues, v)
		}
	}

	// Si no hay valores válidos, retornamos 0
	if len(validValues) == 0 {
		return 0
	}

	// Ordenamos los valores filtrados
	sort.Ints(validValues)

	n := len(validValues)
	if n%2 == 0 {
		// Si hay un número par de elementos, la mediana es el promedio de los dos del medio
		return float64(validValues[n/2-1]+validValues[n/2]) / 2.0
	}
	// Si hay un número impar de elementos, la mediana es el valor del medio
	return float64(validValues[n/2])
}

func main() {
	// Definimos la matriz de ventas con meses sin datos (-1)
	sales := [][]int{
		{1856, -1, 30924, 87478, 328, 2653, 387, 3754, 387587, 2873, -1, 32},
		{5865, 5456, 3983, -1, 9957, 4785, -1, 3838, 4959, -1, 7766, 2534},
		{-1, -1, 67, 99, 265, 376, 232, 223, 4546, 564, 4544, 3434},
	}

	highestMedian := 0.0
	bestAgent := -1

	// Calculamos la mediana para cada agente y determinamos el mejor
	for agent := 0; agent < NUM_AGENTS; agent++ {
		med := median(sales[agent])
		fmt.Printf("Agente %d - Mediana de ventas: %.2f\n", agent+1, med)
		if med > highestMedian {
			highestMedian = med
			bestAgent = agent + 1
		}
	}

	if bestAgent != -1 {
		fmt.Printf("\nEl agente con la mediana de ventas más alta es el Agente %d con una mediana de %.2f\n", bestAgent, highestMedian)
	} else {
		fmt.Println("\nNo hay suficientes datos para determinar la mediana de ventas.")
	}
}
