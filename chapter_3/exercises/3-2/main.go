package main

import (
	"fmt"
	"sort"
)

/*
Rewrite the code that finds the agent with the best monthly sales average so that it finds the agent with the highest median sales.
As stated earlier, the median of a set of values is the “one in the middle,” such that half of the other values are higher and half of
the other values are lower. If there is an even number of values, the median is the simple average of the two values in the middle.
For example, in the set 10, 6, 2, 14, 7, 9, the values in the middle are 7 and 9. The average of 7 and 9 is 8, so 8 is the median.
*/

// Definimos la cantidad de agentes y meses
const NUM_AGENTS = 3
const NUM_MONTHS = 12

// Función para calcular la mediana de un conjunto de valores
func median(values []int) float64 {
	sort.Ints(values) // Ordenamos los valores

	n := len(values)
	if n%2 == 0 {
		// Si hay un número par de elementos, la mediana es el promedio de los dos del medio
		return float64(values[n/2-1]+values[n/2]) / 2.0
	}
	// Si hay un número impar de elementos, la mediana es el valor del medio
	return float64(values[n/2])
}

func main() {
	// Definimos la matriz de ventas
	sales := [][]int{
		{1856, 498, 30924, 87478, 328, 2653, 387, 3754, 387587, 2873, 276, 32},
		{5865, 5456, 3983, 6464, 9957, 4785, 3875, 3838, 4959, 1122, 7766, 2534},
		{23, 55, 67, 99, 265, 376, 232, 223, 4546, 564, 4544, 3434},
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

	fmt.Printf("\nEl agente con la mediana de ventas más alta es el Agente %d con una mediana de %.2f\n", bestAgent, highestMedian)
}
