package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type wordAggregator []string

func newWordAggregator(size int) wordAggregator {
	return make([]string, 0, size)
}

func (w *wordAggregator) Push(s string) {
	if w.IsFull() {
		for i := 0; i < cap(*w)-1; i++ {
			(*w)[i] = (*w)[i+1]
		}
		(*w)[cap(*w)-1] = s
	} else {
		*w = append(*w, s)
	}
}

func (w *wordAggregator) GetNgram() string {
	s := ""

	for i, word := range *w {
		if i == cap(*w)-1 {
			s += word
		} else {
			s += (word + " ")
		}
	}

	return s
}

func (w *wordAggregator) IsFull() bool {
	return len(*w) == cap(*w)
}

var fileName string
var ngramSize int

func init() {
	const (
		defaultFileName  = ""
		defaultNgramSize = 3
	)

	flag.IntVar(&ngramSize, "s", defaultNgramSize, "size of the ngrams")
	flag.IntVar(&ngramSize, "size", defaultNgramSize, "size of the ngrams")

	flag.StringVar(&fileName, "f", defaultFileName, "input filename")
	flag.StringVar(&fileName, "file", defaultFileName, "input filename")
}

func main() {
	flag.Usage = func() {
		usageText := `str2ngrams - a utility for generating unique ngrams for text input.

Given text input, this utility generates the unique ngrams of the desired size, and outputs one ngram per line.

USAGE
  $ str2ngrams [ -h | --help ] [ -s <ngram size> | --size <ngram size> ] [ -f <filename> | --file <filename> ]

OPTIONS
  -f, --file  input filename. Standard input when omitted
  -s, --size  size of the desired ngrams, default value is 3
  -h, --help  print the help message

EXAMPLES
  $ echo "here is the text" | str2ngrams -s 2
  here is
  is the
  the text

  $ str2ngrams -f myfile
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
	scanner.Split(bufio.ScanWords)

	ngrams := make(map[string]struct{})
	wg := newWordAggregator(ngramSize)

	for scanner.Scan() {
		w := scanner.Text()
		w = strings.TrimRight(w, ",.!?")
		wg.Push(w)

		if wg.IsFull() {
			ngram := wg.GetNgram()

			// Track and print each unique ngram.
			if _, ok := ngrams[ngram]; !ok {
				ngrams[ngram] = struct{}{}
				fmt.Println(ngram)
			}
		}
	}
}
