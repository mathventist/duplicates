package main

import (
	"fmt"
	"os"
	"regexp"
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

	fmt.Println(preprocess(text))
}

func preprocess(s string) string {
	// Remove "St.", for example in "St. Paul"
	st := strings.ReplaceAll(s, "St.", "")

	// Replace ligatures
	st = strings.ReplaceAll(st, "æ", "ae")

	// Remove internal punctuation
	a := regexp.MustCompile(`[,'"“;:”’]`)
	st = a.ReplaceAllString(st, "")

	// Replace hyphens with single space
	st = strings.ReplaceAll(st, "-", " ")

	// Compress multiple whitespaces into a single space
	b := regexp.MustCompile(`\s+`)
	st = b.ReplaceAllString(st, " ")

	return st
}
