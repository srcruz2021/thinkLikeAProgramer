package main

import (
	"fmt"
	"unicode"
)

func contarDiferencias(mensajeBase, mensajeRecibido string) int32 {
	longitudBase := len(mensajeBase)

	// Valido las  restricciones
	if len(mensajeRecibido) < 1 || len(mensajeRecibido) > 99 {
		fmt.Println("Error: La longitud del mensaje recibido debe estar entre 1 y 99.")
		return -1
	}
	if len(mensajeRecibido)%longitudBase != 0 {
		fmt.Println("Error: La longitud del mensaje recibido debe ser múltiplo de", longitudBase)
		return -1
	}
	for _, char := range mensajeRecibido {
		if !unicode.IsUpper(char) {
			fmt.Println("Error: El mensaje debe contener solo letras mayúsculas A-Z.")
			return -1
		}
	}

	// Contar diferencias por cada segmento del mensaje base esperado
	diferencias := int32(0)
	for i := 0; i < len(mensajeRecibido); i++ {
		if mensajeRecibido[i] != mensajeBase[i%longitudBase] {
			diferencias++
		}
	}

	return diferencias
}

func main() {
	mensajeBase := "SOS"
	mensajeRecibido := "SOSSPSSQSSOR" // Prueba con diferentes valores aquí

	diferencias := contarDiferencias(mensajeBase, mensajeRecibido)
	if diferencias != -1 {
		fmt.Printf("Número de letras modificadas: %d\n", diferencias)
	}
}
