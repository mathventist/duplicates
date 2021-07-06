package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/golang-collections/collections/set"
)

func main() {
	flag.Usage = func() {
		usageText := `resemblance - a utility for computing the resemblance of one document within another.

Given two documents A and B, and the sets of ngrams (for equal n) for each, S(A) and S(B), the resemblance R(A,B) of documents A and B is defined as:

  R(A,B) = |S(A) ∩ S(B)| / |S(A) ∪ S(B)|

Input files must each contain a single ngram per line, and the ngrams must all be the same size for an accurate calculation.

The output is a floating point value, greater or equal to 0.0 and less than or equal to 1.0. A value of 1.0 indicates complete resemblance.

USAGE
  $ resemblance [ -h | --help ] <file A> <file B>

OPTIONS
  -h, --help  print the help message

EXAMPLES
  $ resemblance fileA fileB
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

	a, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening file: ", err)
		flag.Usage()

		os.Exit(1)
	}
	defer a.Close()

	b, err := os.Open(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening file: ", err)
		flag.Usage()

		os.Exit(1)
	}
	defer b.Close()

	sa := set.New()
	fs := bufio.NewScanner(a)
	for fs.Scan() {
		sa.Insert(fs.Text())
	}

	sb := set.New()
	fs = bufio.NewScanner(b)
	for fs.Scan() {
		sb.Insert(fs.Text())
	}

	intersection := sa.Intersection(sb)
	union := sa.Union(sb)

	fmt.Fprintf(os.Stdout, "%v\n", float64(intersection.Len())/float64(union.Len()))
	os.Exit(0)
}
