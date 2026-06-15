package main

import "strings"

func uniqueWords(text string) map[string]int {
	frequency := make(map[string]int)

	words := strings.Fields(text)

	for _, word := range words {
		frequency[word]++
	}

	return frequency
}