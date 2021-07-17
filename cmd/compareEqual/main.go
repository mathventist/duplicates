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

func main() {
	// TODO: complete this, add flag for removing stop words.
	flag.Usage = func() {
	}

	flag.Parse()

	args := flag.Args()

	if flag.NArg() != 2 {
		fmt.Fprintln(os.Stderr, "invalid number of arguments; please provide two file names.")
		flag.Usage()

		os.Exit(1)
	}

	c := populateSliceFromFile(args[0])
	d := populateSliceFromFile(args[1])

	a, b := <-c, <-d
	results := compare(a, b, true)

	// Display results
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "File\tNumber of sentences\tPercentage of matched sentences")
	fmt.Fprintf(w, "%v\t%v\t%v\n", args[0], len(a), 100*len(results)/len(a))
	fmt.Fprintf(w, "%v\t%v\t%v\n", args[1], len(b), 100*len(results)/len(b))
	w.Flush()

	fmt.Printf("\n\n%v sentences matches in total:\n\n", len(results))
	for _, r := range results {
		fmt.Fprintf(os.Stdout, "%v\n", r)
	}
}

type indexedString struct {
	Index  int
	String string
}

// TODO: improve this by using goroutines to perform some comparisons concurrently.
func compare(a []string, b []string, removeStops bool) [][2]indexedString {
	la := preprocess(a, removeStops)
	lb := preprocess(b, removeStops)

	var results [][2]indexedString
	bar := progressbar.Default(int64(len(a) * len(b)))

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

func preprocess(a []string, removeStops bool) []string {
	var r []string
	bar := progressbar.Default(int64(len(a)))

	for _, aa := range a {
		bar.Add(1)
		r = append(r, duplicates.Preprocess(aa, removeStops))
	}

	return r
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
