package main

import (
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

	a, err := duplicates.FileToArray(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error processing file: ", err)
		flag.Usage()

		os.Exit(1)
	}

	b, err := duplicates.FileToArray(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error processing file: ", err)
		flag.Usage()

		os.Exit(1)
	}

	sa := set.New()
	for _, aa := range a {
		sa.Insert(aa)
	}
	sb := set.New()
	for _, bb := range b {
		sb.Insert(bb)
	}

	intersection := sa.Intersection(sb)
	fmt.Fprintf(os.Stdout, "%v\n", float64(intersection.Len())/float64(sb.Len()))
	os.Exit(0)
}
