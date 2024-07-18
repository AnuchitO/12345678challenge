package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Calculate the total number of rows
	numRows := len(input)*2 - 1

	for i := 0; i < numRows; i++ {
		// Determine the starting index for each row
		startIndex := 0
		if i < len(input) {
			startIndex = i
		} else {
			startIndex = numRows - i - 1
		}

		// Print leading spaces
		for j := 0; j < startIndex; j++ {
			fmt.Print(" ")
		}

		// Print numbers
		for j := startIndex; j < len(input); j++ {
			fmt.Printf("%d ", input[j])
		}

		fmt.Println()
	}
}
