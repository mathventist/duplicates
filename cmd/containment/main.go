package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/golang-collections/collections/set"
	"github.com/mathventist/duplicates"
)

func main() {
	flag.Usage = func() {
		usageText := `containment - a utility for computing the containment of one document within another.

Given two documents A and B and their respective sets of ngrams, S(A), S(B), the measure to which document B is contained in document A, C(A,B), is defined as:

  C(A,B) = |S(A) ∩ S(B)| / |S(B)|

Input files must each contain a single ngram per line, and the ngrams must all be the same size for an accurate calculation.

The output is a floating point value, greater or equal to 0.0 and less than or equal to 1.0. A value of 1.0 indicates complete containment.

USAGE
  $ containment [ -h | --help ] <file A> <file B>

OPTIONS
  -h, --help  print the help message

EXAMPLES
  $ containment fileA fileB
  0.3
  `

		fmt.Println(usageText)
	}

	flag.Parse()

	args := flag.Args()

	if flag.NArg() != 2 {
		fmt.Fprintln(os.Stderr, "invalid number of arguments; please provide two file names.")
		flag.Usage()

		os.Exit(1)
	}

	// use separate channels here because order is important!
	c := populateSetFromFile(args[0])
	d := populateSetFromFile(args[1])

	a, b := <-c, <-d

	fmt.Fprintf(os.Stdout, "%v\n", duplicates.Containment(a, b))
	os.Exit(0)
}

func populateSetFromFile(fileName string) <-chan *set.Set {
	c := make(chan *set.Set)

	go func() {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error opening file: ", err)
			flag.Usage()

			os.Exit(1)
		}
		defer f.Close()

		s := set.New()

		fmt.Fprintf(os.Stderr, "starting setting ngrams for "+fileName+"...\n")
		fs := bufio.NewScanner(f)
		for fs.Scan() {
			s.Insert(fs.Text())
		}
		fmt.Fprintf(os.Stderr, "...done setting ngrams for "+fileName+"\n")

		c <- s
	}()

	return c
}
