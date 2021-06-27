package main

import (
	"fmt"
	"log"
	"os"

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
			score := duplicates.CompareWord2Vec(aa, bb, model)
			normalizedScore := duplicates.CompareWord2Vec(duplicates.Preprocess(aa, false), duplicates.Preprocess(bb, false), model)
			fmt.Println(aa)
			fmt.Println(bb)
			fmt.Println(score)
			fmt.Println(normalizedScore)
			fmt.Println("------------------------------------------")
		}
	}
}
