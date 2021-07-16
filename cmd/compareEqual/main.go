package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mathventist/duplicates"
)

func main() {
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

func compare(a []string, b []string, removeStops bool) [][2]indexedString {
	var results [][2]indexedString

	for i, aa := range a {
		for j, bb := range b {
			if duplicates.Preprocess(aa, removeStops) == duplicates.Preprocess(bb, removeStops) {
				var match [2]indexedString
				match[0] = indexedString{i, aa}
				match[1] = indexedString{j, bb}

				results = append(results, match)
			}
		}
	}

	return results
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
