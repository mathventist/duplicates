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

	for _, s := range stringToNgrams(text, 3) {
		fmt.Println(s)
	}
}

func stringToNgrams(s string, n int) []string {
	// Remove all non alpha numerics.
	r := regexp.MustCompile("[^a-zA-Z0-9]+")
	ss := r.ReplaceAllString(s, " ")
	st := strings.Split(ss, " ")

	ngrams := make(map[string]struct{})
	for i := 0; i < len(st)-n; i++ {
		// Create the ngram.
		ngram := strings.Join(st[i:i+n], " ")

		// Add each unique ngram.
		if _, ok := ngrams[ngram]; !ok {
			ngrams[ngram] = struct{}{}
		}
	}

	// Return the ngrams, in no particular order.
	keys := make([]string, len(ngrams))
	i := 0
	for k := range ngrams {
		keys[i] = k
		i++
	}
	return keys
}
