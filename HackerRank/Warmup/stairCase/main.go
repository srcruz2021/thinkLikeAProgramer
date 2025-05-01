package main

import (
	"fmt"
	"strings"
)

func stairCase(n int32) {
	if n >= 0 && n <= 100 {
		symbol := "#"
		for i := int32(1); i <= n; i++ {
			blankSpaces := n - i
			fmt.Print(strings.Repeat(" ", int(blankSpaces)))
			fmt.Print(strings.Repeat(symbol, int(i)))
			fmt.Print("\n")
		}
	}
}

func main() {
	stairCase(5)
}
