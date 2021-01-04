package main

import (
	"fmt"
	"os"
	"strings"
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

	// Check for duplicate sentences, ignoring case
	var duplicates []string
	for _, aa := range a {
		for _, bb := range b {
			if strings.ToLower(aa) == strings.ToLower(bb) {
				duplicates = append(duplicates, aa)
			}
		}
	}

	// Display results
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "File\tNumber of sentences\t% of duplicate sentences")
	fmt.Fprintf(w, "%v\t%v\t%v\n", args[1], len(a), 100*len(duplicates)/len(a))
	fmt.Fprintf(w, "%v\t%v\t%v\n", args[2], len(b), 100*len(duplicates)/len(b))
	w.Flush()

	fmt.Printf("\n\n%v duplicate sentences in total:\n\n", len(duplicates))
	for i, d := range duplicates {
		fmt.Printf("%v: %v\n\n", i+1, d)
	}

}
