package main

import (
	"fmt"
)

func topHalf(words [8]string) (pyramid string) {
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

func bottomHalf(words [8]string) (pyramid string) {
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

func challenge(words [8]string) string {
	return topHalf(words) + bottomHalf(words)
}

func main() {
	words := [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}
	pym := challenge(words)
	fmt.Println(pym)
}
