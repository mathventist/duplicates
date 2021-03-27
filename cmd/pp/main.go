package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/mathventist/duplicates"
)

func main() {
	removeStops := flag.Bool("r", false, "remove stop words from text")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(duplicates.Preprocess(scanner.Text(), *removeStops))
	}
}
