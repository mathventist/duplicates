package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mathventist/duplicates"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("missing operand")
		return
	}

	if len(args) > 2 {
		fmt.Println("too many operands")
		return
	}

	text, err := duplicates.FileToString(args[1])
	if err != nil {
		fmt.Printf("error processing %s: %v\n", args[1], err)
		return
	}

	for _, s := range stringToSentences(text) {
		fmt.Println(s)
	}
}

func stringToSentences(s string) []string {
	// Split into sentences
	temp := strings.FieldsFunc(s, func(c rune) bool {
		return c == '.' || c == '?' || c == '!'
	})

	// Trim leading and trailing whitespace. Ignore any sentences that are empty.
	var sentences []string
	for _, s := range temp {
		st := strings.TrimSpace(s)
		if len(st) > 0 {
			sentences = append(sentences, st)
		}
	}

	return sentences
}
