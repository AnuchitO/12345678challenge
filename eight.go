package main

import (
	"fmt"
	"os"
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
	words := args()
	pym := challenge(words)
	fmt.Println()
	fmt.Println(pym)
}

func args() [8]string {
	words := os.Args[1:]
	if len(words) != 8 {
		fmt.Println("Please provide 8 words")
		os.Exit(1)
	}

	return [8]string{words[0], words[1], words[2], words[3], words[4], words[5], words[6], words[7]}
}
