package duplicates

import (
	"bufio"
	"bytes"
	"os"
	"regexp"
	"strings"
)

// Taken from https://gist.github.com/sebleier/554280
var stopWords = []string{
	"i",
	"me",
	"my",
	"myself",
	"we",
	"our",
	"ours",
	"ourselves",
	"you",
	"your",
	"yours",
	"yourself",
	"yourselves",
	"he",
	"him",
	"his",
	"himself",
	"she",
	"her",
	"hers",
	"herself",
	"it",
	"its",
	"itself",
	"they",
	"them",
	"their",
	"theirs",
	"themselves",
	"what",
	"which",
	"who",
	"whom",
	"this",
	"that",
	"these",
	"those",
	"am",
	"is",
	"are",
	"was",
	"were",
	"be",
	"been",
	"being",
	"have",
	"has",
	"had",
	"having",
	"do",
	"does",
	"did",
	"doing",
	"a",
	"an",
	"the",
	"and",
	"but",
	"if",
	"or",
	"because",
	"as",
	"until",
	"while",
	"of",
	"at",
	"by",
	"for",
	"with",
	"about",
	"against",
	"between",
	"into",
	"through",
	"during",
	"before",
	"after",
	"above",
	"below",
	"to",
	"from",
	"up",
	"down",
	"in",
	"out",
	"on",
	"off",
	"over",
	"under",
	"again",
	"further",
	"then",
	"once",
	"here",
	"there",
	"when",
	"where",
	"why",
	"how",
	"all",
	"any",
	"both",
	"each",
	"few",
	"more",
	"most",
	"other",
	"some",
	"such",
	"no",
	"nor",
	"not",
	"only",
	"own",
	"same",
	"so",
	"than",
	"too",
	"very",
	"s",
	"t",
	"can",
	"will",
	"just",
	"don",
	"should",
	"now",
}

// FileToString reads a file and returns it as a single string.
func FileToString(filename string) (string, error) {
	lines, err := FileToArray(filename)

	if err != nil {
		return "", err
	}

	// Recombine all text into single string
	st := strings.Join(lines, "")

	return st, nil
}

// FileToArray reads a file into an array, storing each line as an element of the array.
func FileToArray(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines, nil
}

// Preprocess prepares a string for further processing.
func Preprocess(s string, removeStops bool) string {
	// Remove "St.", for example in "St. Paul"
	st := strings.ReplaceAll(s, "St.", "")

	// Replace ligatures
	st = strings.ReplaceAll(st, "æ", "ae")

	// Replace hyphens with single space
	st = strings.ReplaceAll(st, "-", " ")

	// Remove internal punctuation that doesn't terminate a sentence.
	a := regexp.MustCompile(`[,'"“;:”’]`)
	st = a.ReplaceAllString(st, "")

	// Remove numerics
	n := regexp.MustCompile(`\d+`)
	st = n.ReplaceAllString(st, "")

	if removeStops {
		for _, w := range stopWords {
			re := regexp.MustCompile("\\b" + w + "\\b")
			st = re.ReplaceAllString(st, "")
		}
	}

	// Compress multiple whitespaces into a single space
	b := regexp.MustCompile(`[\s]{2,}`)
	st = b.ReplaceAllString(st, " ")

	// lower case
	return strings.ToLower(st)
}

func dropLast(data []byte) []byte {
	if len(data) > 0 && isSentenceTerminator(data[len(data)-1]) {
		return data[0 : len(data)-1]
	}

	return data
}

// ScanSentences scans a file sentence by sentence.
func ScanSentences(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexAny(data, ".?!"); i >= 0 {
		return i + 1, dropLast(data[0:i]), nil
	}

	if atEOF {
		return len(data), dropLast(data), nil
	}

	return 0, nil, nil
}

func isSentenceTerminator(b byte) bool {
	return b == '.' || b == '?' || b == '!'
}
