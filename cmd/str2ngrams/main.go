package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type wordAggregator struct {
	words []string
	size  int
}

func NewWordAggregator(size int) wordAggregator {
	w := make([]string, size, size)
	return wordAggregator{
		size:  size,
		words: w,
	}
}

func (w wordAggregator) Push(s string) {
	if len(w.words) == w.size {
		for i := 0; i < w.size-1; i++ {
			w.words[i] = w.words[i+1]
		}
		w.words[w.size-1] = s
	} else {
		w.words = append(w.words, s)
	}
}

func (w wordAggregator) GetNgram() string {
	s := ""
	for i, word := range w.words {
		if i == w.size-1 {
			s += word
		} else {
			s += (word + " ")
		}
	}
	return s
}

func (w wordAggregator) IsFull() bool {
	return w.words[w.size-1] != ""
}

func main() {
	var ngramSize = flag.Int("size", 3, "size of the ngram")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	//ngrams := make(map[string]struct{})
	wg := NewWordAggregator(*ngramSize)

	for scanner.Scan() {
		w := scanner.Text()
		wg.Push(w)

		if wg.IsFull() {
			fmt.Println(wg.GetNgram())
		}
		// Track and print each unique ngram.
		//if _, ok := ngrams[ngram]; !ok {
		//	ngrams[ngram] = struct{}{}
		//	fmt.Println(ngram)
		//}
	}
}
