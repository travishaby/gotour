package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counter := map[string]int{}
	words := strings.Fields(s)
	for _, word := range words {
		value := counter[word]
		counter[word] = value + 1
	}

	return counter
}

func main() {
	wc.Test(WordCount)
}
