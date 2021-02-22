package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(scanSentences)

	for scanner.Scan() {
		fmt.Println(strings.TrimSpace(scanner.Text()))
	}
}

func dropLast(data []byte) []byte {
	if len(data) > 0 && isSentenceTerminator(data[len(data)-1]) {
		return data[0 : len(data)-1]
	}

	return data
}

func scanSentences(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
