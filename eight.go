package main

import (
	"fmt"
)

func challenge(input []int) (pyramid string) {
	numRows := len(input)*2 - 1
	for i := 0; i < numRows; i++ {
		startIndex := 0
		if i < len(input) {
			startIndex = i
		} else {
			startIndex = numRows - i - 1
		}

		for j := 0; j < startIndex; j++ {
			pyramid += " "
		}

		for j := startIndex; j < len(input); j++ {
			pyramid += fmt.Sprintf("%d ", input[j])
		}

		pyramid += "\n"
	}

	return pyramid
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	pym := challenge(input)
	fmt.Println(pym)
}
