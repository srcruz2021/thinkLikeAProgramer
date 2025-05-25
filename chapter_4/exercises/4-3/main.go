package main

/*
For our dynamically allocated strings, create a function replaceString that takes three parameters,
each of type arrayString: source, target, and replaceText. The function replaces every occurrence of target in source with replaceText.
For example, if source points to an array containing abcdabee, target points to ab, and replaceText points to xyz,
then when the function ends, source should point to an array containing xyzcdxyzee
*/

import (
	"fmt"
	"strings"
)

// replaceString reemplaza cada ocurrencia de target en source con replaceText.
func replaceString(source *string, target, replaceText string) {
	*source = strings.ReplaceAll(*source, target, replaceText)
}

func main() {
	source := "abcdabee"
	target := "ab"
	replaceText := "xyz"

	replaceString(&source, target, replaceText)

	fmt.Println("Resultado:", source) // Salida esperada: "xyzcdxyzee"
}
