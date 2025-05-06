package main

import (
	"fmt"
	"unicode"
)

type PasswordValidator struct {
	HasSpecialCharacters bool
	HasLowerCase         bool
	HasUpperCase         bool
	HasNumbers           bool
}

func isSpecialCharacter(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
}

func isAlphaNumeric(s rune) bool {
	return unicode.IsDigit(s)
}

func isUpperCase(s rune) bool {
	return unicode.IsUpper(s)
}

func isLowerCase(s rune) bool {
	return unicode.IsLower(s)
}

func main() {
	n := int32(6)
	currentPassword := "Ab1"
	pv := PasswordValidator{}
	SuccessValidation := int32(0)
	validationToExecute := int32(4)

	if int32(len(currentPassword)) <= n {

		for _, char := range currentPassword {
			if !pv.HasSpecialCharacters {
				pv.HasSpecialCharacters = isSpecialCharacter(char)
			}
			if !pv.HasLowerCase {
				pv.HasLowerCase = isLowerCase(char)
			}
			if !pv.HasUpperCase {
				pv.HasUpperCase = isUpperCase(char)
			}
			if !pv.HasNumbers {
				pv.HasNumbers = isAlphaNumeric(char)
			}
		}

		if pv.HasNumbers {
			SuccessValidation++
		}
		if pv.HasSpecialCharacters {
			SuccessValidation++
		}
		if pv.HasLowerCase {
			SuccessValidation++
		}
		if pv.HasUpperCase {
			SuccessValidation++
		}
	} else {
		fmt.Println(6 - len(currentPassword))
	}

	fmt.Println(fmt.Sprintf("Number of pending validations: %d", validationToExecute-SuccessValidation))
	fmt.Println(fmt.Sprintf("Current password: %s", currentPassword))
}
