package duplicates

import (
	"bufio"
	"os"
	"strings"
)

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
