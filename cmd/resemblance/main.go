package main

import (
	"fmt"
	"os"

	"github.com/mathventist/duplicates"

	"github.com/golang-collections/collections/set"
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

	sa := set.New(a)
	sb := set.New(b)

	intersection := sa.Intersection(sb)
	union := sa.Union(sb)

	fmt.Printf("Resemblance is %v\n", float64(intersection.Len())/float64(union.Len()))
}
