package duplicates

import (
	"bufio"
	"os"
	"strings"
)

//func main() {
//args := os.Args[1:]

//if len(args) < 2 {
//	fmt.Println("missing operand")
//}

//if len(args) > 2 {
//	fmt.Println("too many operands")
//}

//fileone, err := FileToString(args[0])
//if err != nil {
//	fmt.Printf("error processing %s\n", args[0])
//}

//filetwo, err := FileToString(args[1])
//if err != nil {
//	fmt.Printf("error processing %s\n", args[1])
//}

//fone := stringToSentences(fileone)
//ftwo := stringToSentences(filetwo)

//// Check for duplicate sentences, ignoring case
//var duplicates []string
//for _, s := range fone {
//	for _, t := range ftwo {
//		if strings.ToLower(s) == strings.ToLower(t) {
//			duplicates = append(duplicates, s)
//		}
//	}
//}

//// Display results
//w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
//fmt.Fprintln(w, "File\tNumber of sentences\t% of duplicate sentences")
//fmt.Fprintf(w, "%v\t%v\t%v\n", args[0], len(fone), 100*len(duplicates)/len(fone))
//fmt.Fprintf(w, "%v\t%v\t%v\n", args[1], len(ftwo), 100*len(duplicates)/len(ftwo))
//w.Flush()

//fmt.Printf("\n\nDuplicate sentences:\n\n")
//for i, d := range duplicates {
//	fmt.Printf("%v: %v\n\n", i+1, d)
//}
//}

// FileToString reads a file and returns it as a single string.
func FileToString(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", err
	}

	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// Recombine all text into single string
	st := strings.Join(lines, "")

	return st, nil
}
