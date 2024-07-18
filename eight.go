package main

import (
	"fmt"
	"os"
)

func topHalf(words []string) (pyramid string) {
	for cursor := 0; cursor < len(words); cursor++ {
		for space := 0; space < cursor; space++ {
			pyramid += " "
		}

		for j := cursor; j < len(words); j++ {
			pyramid += fmt.Sprintf(" %s", words[j])
		}

		pyramid += "\n"
	}

	return pyramid
}

func bottomHalf(words []string) (pyramid string) {
	for i := len(words) - 1; i >= 0; i-- {
		for space := 0; space < i; space++ {
			pyramid += " "
		}

		for j := i; j < len(words); j++ {
			pyramid += fmt.Sprintf(" %s", words[j])
		}

		pyramid += "\n"
	}

	return pyramid
}

func challenge(words []string) string {
	return topHalf(words) + bottomHalf(words)
}

func main() {
	words := args()
	pym := challenge(words)
	fmt.Println()
	fmt.Println(pym)
}

func args() []string {
	words := os.Args[1:]
	if len(words) == 0 {
		words = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	}

	return words
}
