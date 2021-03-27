package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mathventist/duplicates"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(duplicates.ScanSentences)

	for scanner.Scan() {
		fmt.Println(strings.TrimSpace(scanner.Text()))
	}
}
