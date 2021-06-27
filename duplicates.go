package duplicates

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"regexp"
	"strings"

	"code.sajari.com/word2vec"
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

// FileToArray reads a file into an array, storing each sentence as an element of the array.
func FileToArray(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(ScanSentences)

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

// ScanSentences scans a file sentence by sentence.
func ScanSentences(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexAny(data, ".?!"); i >= 0 {

		for abbreviationIsAtPosition(string(data[0 : i+1])) {
			k := bytes.IndexAny(data[i+1:], ".?!")

			if k == -1 {

				if atEOF {
					return len(data), data, nil
				}

				return 0, nil, nil
			}

			i += (k + 1)
		}

		// handle special cases for sentences that end with a single quote, '."'
		if i < len(data)-1 {
			if isEndQuote(data[i+1]) {
				i++
			}
		}

		// skip over any trailing whitespace
		j := i + 1
		for j < len(data) && isWhiteSpace(data[j]) {
			j++
		}

		return j, data[0 : i+1], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}

func isSentenceTerminator(b byte) bool {
	return b == '.' || b == '?' || b == '!'
}

func isWhiteSpace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\v' || b == '\f' || b == '\r' || b == '\n'
}

func isEndQuote(b byte) bool {
	return b == '"' || b == '\''
}

func abbreviationIsAtPosition(s string) bool {
	return strings.HasSuffix(s, "St.")
}

func CompareWord2Vec(a, b string, model *word2vec.Model) float32 {
	similarCounter := 0
	similarTotal := float32(0)

	aWords := strings.Fields(a)
	for _, aWord := range aWords {
		aExpr := word2vec.Expr{}
		aExpr.Add(1, aWord)

		maxSim := float32(0)

		bWords := strings.Fields(b)
		for _, bWord := range bWords {
			bExpr := word2vec.Expr{}
			bExpr.Add(1, bWord)

			sim, err := model.Cos(aExpr, bExpr)
			if err != nil {
				log.Println("error calculating cos")
				log.Fatal(err)
			}
			if sim > maxSim {
				maxSim = sim
			}
		}
		if maxSim > 0 {
			similarCounter++
			similarTotal = similarTotal + maxSim
		}
	}

	if similarCounter == 0 {
		return float32(0)
	}

	return similarTotal / float32(similarCounter)
}
