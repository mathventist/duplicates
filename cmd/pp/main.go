package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/mathventist/duplicates"
)

func main() {
	flag.Usage = func() {
		usageText := `pp - a utility for normalizing text for further comparison.

It removes titles, numerics, hyphens, and internal sentence punctuation, expands ligatures, and compresses multiple whitespace characters into a single whitespace character.

Optionally, this utility will also strip out English stop words (see https://gist.github.com/sebleier/554280).

USAGE
  $ pp

OPTIONS
  -r remove stop words from the text

EXAMPLES
  $ echo "testing 1 i 2 me 3 normalizing" | pp -r
  testing normalizing
  `

		fmt.Println(usageText)
	}

	removeStops := flag.Bool("r", false, "remove stop words from text")
	flag.Parse()

	file := os.Stdin
	fi, err := file.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "stdin error", err)
		flag.Usage()

		os.Exit(1)
	}

	// Exit if stdin is empty.
	size := fi.Size()
	if size == 0 {
		fmt.Fprintln(os.Stderr, "error: standard input is empty.")
		flag.Usage()

		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, duplicates.Preprocess(scanner.Text(), *removeStops))
	}

	os.Exit(0)
}
