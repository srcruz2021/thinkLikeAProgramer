package main

import "fmt"

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}

	positiveRatio := 0
	negativeRatio := 0
	neutro := 0

	for _, v := range arr {
		if v < 0 && v >= -100 {
			negativeRatio++
		}
		if v > 0 && v <= 100 {
			positiveRatio++
		}
		if v == 0 {
			neutro++
		}
	}

	n := len(arr)
	nPositivo := float64(positiveRatio) / float64(n)
	nNegativo := float64(negativeRatio) / float64(n)
	nNeutro := float64(neutro) / float64(n)
	fmt.Println(fmt.Sprintf("positive ratio: %0.6f", nPositivo))
	fmt.Println(fmt.Sprintf("Negative ratio: %0.6f", nNegativo))
	fmt.Println(fmt.Sprintf("Neutro ratio: %0.6f", nNeutro))

}
