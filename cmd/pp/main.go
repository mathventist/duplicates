package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/mathventist/duplicates"
)

var removeStops bool
var fileName string

func init() {
	const (
		defaultRemoveStops = false
		defaultFileName    = ""
	)

	flag.BoolVar(&removeStops, "r", defaultRemoveStops, "remove stop words from text")
	flag.BoolVar(&removeStops, "removeStops", defaultRemoveStops, "remove stop words from text")

	flag.StringVar(&fileName, "f", defaultFileName, "input filename")
	flag.StringVar(&fileName, "file", defaultFileName, "input filename")
}

func main() {
	flag.Usage = func() {
		usageText := `pp - a utility for normalizing text for further comparison.

It removes titles, numerics, hyphens, and internal sentence punctuation, expands ligatures, and compresses multiple whitespace characters into a single whitespace character.

Optionally, this utility will also strip out English stop words (see https://gist.github.com/sebleier/554280).

USAGE
  $ pp [ -h | --help ] [ -f <filename> | --file <filename> ] [ -r | --removeStops ]

OPTIONS
  -f, --file         input filename. Standard input when omitted
  -r, --removeStops  remove stop words from the text
  -h, --help         print the help message

EXAMPLES
  $ echo "testing 1 i 2 me 3 normalizing" | pp -r
  testing normalizing

  $ pp -f myfile
  `

		fmt.Println(usageText)
	}

	flag.Parse()

	var inputErr error
	file := os.Stdin

	if len(fileName) > 0 {
		file, inputErr = os.Open(fileName)

		if inputErr != nil {
			fmt.Fprintln(os.Stderr, "error opening input file: ", inputErr)
			flag.Usage()

			os.Exit(1)
		}

		defer file.Close()
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "stdin error: ", err)
		flag.Usage()

		os.Exit(1)
	}

	// Exit if stdin is empty.
	size := fi.Size()
	if size == 0 {
		fmt.Fprintln(os.Stderr, "input is empty")
		flag.Usage()

		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintln(os.Stdout, duplicates.Preprocess(scanner.Text(), removeStops))
	}

	os.Exit(0)
}
