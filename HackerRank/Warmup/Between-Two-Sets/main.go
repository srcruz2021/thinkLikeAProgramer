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
There will be two arrays of integers. Determine all integers that satisfy the following two conditions:

1. The elements of the first array are all factors of the integer being considered
2. The integer being considered is a factor of all elements of the second array

These numbers are referred to as being between the two arrays. Determine how many such numbers exist.

Example

a = [2,6]
b = [24,36]

There are two numbers between the arrays: 6 and 12.

6 % 2 = 0 , 6 % 6 = 0, 24 % 6 = 0 and
36 % 6 = 0 for the first value.

12 % 2 = 0,  12 % 6 = 0  and

24 % 12 = 0, 36 % 12 = 0 for the second value. return 2

Function Description

Complete the getTotalX function in the editor below. It should return the number of integers that are betwen the sets.

getTotalX has the following parameter(s):

int a[n]: an array of integers
int b[m]: an array of integers
Returns

int: the number of integers that are between the sets
Input Format

The first line contains two space-separated integers, n  and m, the number of elements in arrays a and b.

The second line contains n distinct space-separated integers a[i]  where 0 <= i < n.

The third line contains m, distinct space-separated integer b[j] where 0 <= j < m.

Constraints

- 1 <= n,m <= 10
- 1 <= a[i] <= 100
- 1 <= b[j] <= 100

*/

// Function to find GCD
func gcd(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find LCM
func lcm(a, b int32) int32 {
	return a * b / gcd(a, b)
}

// Function to count numbers between sets
func getTotal(a, b []int32) int32 {
	// Find LCM of array `a`
	lcmA := a[0]
	for _, num := range a {
		lcmA = lcm(lcmA, num)
	}

	// Find GCD of array `b`
	gcdB := b[0]
	for _, num := range b {
		gcdB = gcd(gcdB, num)
	}

	// Count numbers between LCM(a) and GCD(b) that are factors of gcdB
	count := int32(0)
	for i := lcmA; i <= gcdB; i += lcmA {
		if gcdB%i == 0 {
			count++
		}
	}

	return count
}

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	result := getTotal(a, b)
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

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
