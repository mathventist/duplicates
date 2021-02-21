package main

import (
	"fmt"
	"os"

	"github.com/golang-collections/collections/set"
	"github.com/mathventist/duplicates"
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

	sa := set.New()
	for _, aa := range a {
		sa.Insert(aa)
	}
	sb := set.New()
	for _, bb := range b {
		sb.Insert(bb)
	}

	intersection := sa.Intersection(sb)

	fmt.Printf("Containment is %v\n", float64(intersection.Len())/float64(sb.Len()))
}
