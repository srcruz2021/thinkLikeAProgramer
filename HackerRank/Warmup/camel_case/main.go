package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "itIsCertain"
	numberOFWords := int32(1)
	for _, v := range s {
		if unicode.IsUpper(v) {
			numberOFWords++
		}
	}
	fmt.Println(numberOFWords)
}
