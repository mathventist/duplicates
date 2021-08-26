package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mathventist/duplicates"
	"github.com/schollz/progressbar/v3"
)

var removeStops bool

func init() {
	const (
		defaultRemoveStops = false
	)

	flag.BoolVar(&removeStops, "r", defaultRemoveStops, "remove stop words from text")
	flag.BoolVar(&removeStops, "removeStops", defaultRemoveStops, "remove stop words from text")
}

func main() {
	flag.Usage = func() {
		usageText := `compareEqual - a utility for finding matching sentences in two files.

It uses the duplicate package's Preprocess method to normalize the text before comparing.

USAGE
  $ compareEqual [ -h | --help ] [ -r | --removeStops ] FILE1 FILE2

OPTIONS
  -r, --removeStops  remove stop words from the text
  -h, --help         print the help message

EXAMPLES

  $ compareEqual -r myfile1 myfile2
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

	fileName1, fileName2 := args[0], args[1]

	c := populateSliceFromFile(fileName1)
	d := populateSliceFromFile(fileName2)

	a, b := <-c, <-d
	results := compare(a, b, fileName1, fileName2, removeStops)

	// Display result summary
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "File\tNumber of sentences\tPercentage of matched sentences")
	fmt.Fprintf(w, "%v\t%v\t%v\n", fileName1, len(a), 100*len(results)/len(a))
	fmt.Fprintf(w, "%v\t%v\t%v\n", fileName2, len(b), 100*len(results)/len(b))
	w.Flush()

	// Display full results
	fmt.Printf("\n\n%v matched sentences.\n\n", len(results))
	for _, r := range results {
		fmt.Fprintf(os.Stdout, "%v sentence number %v\n\n\t%v\n\nmatched to %v sentence number %v\n\n\t%v\n\n",
			fileName1, r[0].Index+1, r[0].String,
			fileName2, r[1].Index+1, r[1].String,
		)
	}
}

type indexedString struct {
	Index  int
	String string
}

func compare(a []string, b []string, fileName1 string, fileName2 string, removeStops bool) [][2]indexedString {
	ca := preprocess(a, fileName1, removeStops)
	cb := preprocess(b, fileName2, removeStops)

	la, lb := <-ca, <-cb

	var results [][2]indexedString
	bar := progressbar.Default(int64(len(a)*len(b)), "comparing files...")

	// TODO: improve performance by using goroutines to run comparisons concurrently.
	for i, aa := range la {
		for j, bb := range lb {
			bar.Add(1)
			if aa == bb {
				var match [2]indexedString
				match[0] = indexedString{i, a[i]}
				match[1] = indexedString{j, b[j]}

				results = append(results, match)
			}
		}
	}

	return results
}

func isSentenceTerminator(b byte) bool {
	return b == '.' || b == '?' || b == '!'
}

func preprocess(a []string, fileName string, removeStops bool) <-chan []string {
	c := make(chan []string)

	go func() {
		var r []string
		bar := progressbar.Default(int64(len(a)), "preprocessing "+fileName+"...")

		for _, aa := range a {
			bar.Add(1)
			normalizedText := duplicates.Preprocess(aa, removeStops)

			// trim trailing punctuation
			if isSentenceTerminator(normalizedText[len(normalizedText)-1]) {
				normalizedText = normalizedText[:len(normalizedText)-1]
			}

			r = append(r, normalizedText)
		}
		c <- r
	}()

	return c
}

func populateSliceFromFile(fileName string) <-chan []string {
	c := make(chan []string)

	go func() {
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error opening file: ", err)
			flag.Usage()

			os.Exit(1)
		}
		defer f.Close()

		fileScanner := bufio.NewScanner(f)
		fileScanner.Split(duplicates.ScanSentences)

		var lines []string
		for fileScanner.Scan() {
			lines = append(lines, fileScanner.Text())
		}

		c <- lines
	}()

	return c
}
