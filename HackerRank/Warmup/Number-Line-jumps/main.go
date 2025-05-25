package main

import (
	"fmt"
	"os"
)

type Kangaroo struct {
	InitialPosition int32
	JumpingDistance int32
}

func (k Kangaroo) CalculateDistance() int32 {
	var distance int32
	if (k.JumpingDistance >= 0 && k.JumpingDistance <= 10000) || (k.InitialPosition >= 0 && k.InitialPosition <= 10000) {
		distance = k.InitialPosition + k.JumpingDistance
		return distance
	}
	os.Exit(0)
	return 0
}

func EvaluateDistance(k1, k2 Kangaroo) string {
	dk1 := k1.CalculateDistance()
	dk2 := k2.CalculateDistance()
	if (k2.JumpingDistance > k1.JumpingDistance) && (k2.InitialPosition > k1.InitialPosition) {
		return "NO"
	}
	if dk1 == dk2 {
		return "YES"
	}
	return "NO"
}

func main() {
	k1 := Kangaroo{
		InitialPosition: 0,
		JumpingDistance: 5,
	}
	k2 := Kangaroo{
		InitialPosition: 2,
		JumpingDistance: 3,
	}
	fmt.Println(EvaluateDistance(k1, k2))
}
