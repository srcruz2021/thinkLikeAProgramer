package main

import "fmt"

/*
 */
// un arreglo con 2 valores, el primero para almacenar el valor, y el segundo para almacenar la suma de repeticiones.
func main() {
	candles := []int32{3, 2, 1, 3, 5, 5}
	highestCandle := []int32{0, 0}

	for i := 0; i < len(candles); i++ {
		if candles[i] >= 0 && candles[i] <= 10000000 {
			if candles[i] == highestCandle[0] {
				// Valido si el valor anterior es menor al valor actual entonces reinicio el contador.
				if highestCandle[0] < candles[i] {
					highestCandle[1] = 0
				}
				highestCandle[0] = candles[i]
				highestCandle[1]++
			}
			if candles[i] > highestCandle[0] {
				highestCandle[0] = candles[i]
				highestCandle[1] = 0
				highestCandle[1]++
			}
		}
	}

	fmt.Println(highestCandle[1])
	fmt.Println(fmt.Sprintf("highest candle value: %d, number of repetitions: %d", highestCandle[0], highestCandle[1]))

}
