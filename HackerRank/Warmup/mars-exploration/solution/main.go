package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

/*
 * Complete the 'marsExploration' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func marsExploration(s string) int32 {
	mensajeBase := "SOS"
	mensajeRecibido := s
	longitudBase := len(mensajeBase)

	// Validar restricciones
	if len(mensajeRecibido) < 1 || len(mensajeRecibido) > 99 {
		return -1
	}
	if len(mensajeRecibido)%longitudBase != 0 {
		return -1
	}
	for _, char := range mensajeRecibido {
		if !unicode.IsUpper(char) {
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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := marsExploration(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
