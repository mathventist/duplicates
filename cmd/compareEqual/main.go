package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mathventist/duplicates"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("missing operand")
		return
	}

	if len(args) > 3 {
		fmt.Println("too many operands")
		return
	}

	a, err := duplicates.FileToArray(args[1])
	if err != nil {
		fmt.Printf("error processing %s: %v\n", args[1], err)
		return
	}

	b, err := duplicates.FileToArray(args[2])
	if err != nil {
		fmt.Printf("error processing %s: %v\n", args[2], err)
		return
	}

	results := compare(a, b)

	// Display results
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "File\tNumber of sentences\t% of duplicate sentences")
	fmt.Fprintf(w, "%v\t%v\t%v\n", args[1], len(a), 100*len(results)/len(a))
	fmt.Fprintf(w, "%v\t%v\t%v\n", args[2], len(b), 100*len(results)/len(b))
	w.Flush()

	fmt.Printf("\n\n%v duplicate sentences in total:\n\n", len(results))
	for i, d := range results {
		fmt.Printf("%v: %v\n\n", i+1, d)
	}

}

func compare(a []string, b []string) []string {
	var results []string

	for _, aa := range a {
		for _, bb := range b {
			if duplicates.Preprocess(aa, false) == duplicates.Preprocess(bb, false) {
				results = append(results, aa)
			}
		}
	}

	return results
}
