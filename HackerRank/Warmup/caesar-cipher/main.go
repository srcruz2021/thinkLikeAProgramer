package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var stringBuilder, codeMsjBuilder strings.Builder
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	rotateFactor := 2
	separator := int32(45) // "_"

	stringBuilder.WriteString(alphabet[rotateFactor:]) // Le quito las primeras tres letras
	stringBuilder.WriteString(alphabet[:rotateFactor]) // tomo las primeras tres letras
	rotateAlphabet := stringBuilder.String()           // armo el string en el orden deseado
	fmt.Println(rotateAlphabet)

	msj := "middle-Outz"
	letterPosition := []int{}
	separatorPosition := []int{}
	for k1, v := range msj {
		if v == separator {
			separatorPosition = append(separatorPosition, k1)
		}
		for k, l := range alphabet {
			if strings.EqualFold(string(v), string(l)) {
				letterPosition = append(letterPosition, k)
			}
		}

	}

	fmt.Println(letterPosition)
	fmt.Println(separatorPosition)

	for _, v := range letterPosition {
		codeMsjBuilder.WriteString(string(rotateAlphabet[v]))
	}

	codeMsj := codeMsjBuilder.String()
	finalMsj := codeMsj[:separatorPosition[0]] + string(separator) + codeMsj[separatorPosition[0]:]

	fmt.Println(fmt.Sprintf("Mensaje codificado: %s", finalMsj))

}

func IsAscii(r rune) bool {
	return unicode.IsLetter(r)
}
