package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(preprocess(scanner.Text()))
	}
}

func preprocess(s string) string {
	// Remove "St.", for example in "St. Paul"
	st := strings.ReplaceAll(s, "St.", "")

	// Replace ligatures
	st = strings.ReplaceAll(st, "æ", "ae")

	// Replace hyphens with single space
	st = strings.ReplaceAll(st, "-", " ")

	// Remove internal punctuation that doesn't terminate a sentence.
	a := regexp.MustCompile(`[,'"“;:”’]`)
	st = a.ReplaceAllString(st, "")

	// Remove numerics
	n := regexp.MustCompile(`\d+`)
	st = n.ReplaceAllString(st, "")

	// Compress multiple whitespaces into a single space
	b := regexp.MustCompile(`[\s]{2,}`)
	st = b.ReplaceAllString(st, " ")

	// lower case
	return strings.ToLower(st)
}
