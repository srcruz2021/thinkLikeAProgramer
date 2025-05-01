package main

import (
	"fmt"
	"math"
)

func main() {
	grades := []int32{73, 67, 38, 33}
	roundGrades := []int32{}
	roundRuleValue := int32(3)
	minRoundValue := int32(38)

	if len(grades) <= 60 {
		for _, grade := range grades {
			if grade >= 0 && grade <= 100 {
				if grade >= minRoundValue {
					roundGrades = applyRound(grade, roundRuleValue, roundGrades)
				} else {
					roundGrades = append(roundGrades, grade)
				}
			}

		}
		fmt.Println(roundGrades)
	}
}

func applyRound(grade int32, roundRuleValue int32, roundGrades []int32) []int32 {
	// calculate the next value multiply of 5 given an 'n' number
	nextGrade := nextMultipleOfFive(grade)
	difference := nextGrade - grade
	if difference < roundRuleValue {
		roundGrades = append(roundGrades, nextGrade)
	} else {
		roundGrades = append(roundGrades, grade)
	}
	return roundGrades
}

func nextMultipleOfFive(n int32) int32 {
	return int32(math.Ceil(float64(n)/5.0)) * 5
}
