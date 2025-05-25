package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumNumber' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING password
 */

func minimumNumber(n int32, password string) int32 {
	result := minimumNumberOfAPassword(n, password)
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	password := readLine(reader)

	answer := minimumNumber(n, password)

	fmt.Fprintf(writer, "%d\n", answer)

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

func minimumNumberOfAPassword(n int32, password string) int32 {
	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"

	hasDigit := false
	hasLower := false
	hasUpper := false
	hasSpecial := false

	// Check for required character types
	for _, char := range password {
		if strings.ContainsRune(numbers, char) {
			hasDigit = true
		}
		if strings.ContainsRune(lower_case, char) {
			hasLower = true
		}
		if strings.ContainsRune(upper_case, char) {
			hasUpper = true
		}
		if strings.ContainsRune(special_characters, char) {
			hasSpecial = true
		}
	}

	// Count missing character types
	missingTypes := int32(0)
	if !hasDigit {
		missingTypes++
	}
	if !hasLower {
		missingTypes++
	}
	if !hasUpper {
		missingTypes++
	}
	if !hasSpecial {
		missingTypes++
	}

	// Ensure password length is at least 6
	return max(missingTypes, 6-n)
}

// Helper function to find max of two numbers
func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
