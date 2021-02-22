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
		fmt.Print(preprocess(scanner.Text()))
	}
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
