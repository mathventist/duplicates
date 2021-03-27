package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"code.sajari.com/word2vec"
	"github.com/mathventist/duplicates"
)

func main() {
	args := os.Args

	if len(args) < 4 {
		fmt.Println("missing operand")
		return
	}

	if len(args) > 4 {
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

	// Load word2vec model
	// Ex: curl -O https://s3.amazonaws.com/dl4j-distribution/GoogleNews-vectors-negative300.bin.gz
	r, err := os.Open(args[3])
	if err != nil {
		fmt.Printf("error opening %s: %v\n", args[3], err)
	}
	defer r.Close()

	model, err := word2vec.FromReader(r)
	if err != nil {
		log.Fatalf("error loading model: %v", err)
	}

	for _, aa := range a {
		for _, bb := range b {
			score := compare(aa, bb, model)
			fmt.Println(aa)
			fmt.Println(bb)
			fmt.Println(score)
			fmt.Println("------------------------------------------")
		}
	}
}

func compare(a string, b string, model *word2vec.Model) float32 {
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
