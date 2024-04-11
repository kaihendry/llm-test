package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	nopunctuation := removePunctuation(string(bytes))
	words := strings.Fields(strings.ToLower(nopunctuation))
	// create a map for each letter and the number of times it appears
	letterCount := make(map[rune]int)
	for _, word := range words {
		for _, r := range word {
			letterCount[r]++
		}
	}
	// if any letter appears more than once, we fail
	for r, count := range letterCount {
		if count > 1 {
			fmt.Printf("not ok 6 - letter %c appears more than once\n", r)
			return
		}
	}
	fmt.Printf("ok 6 - %s\n", words)
}

func removePunctuation(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			// If the rune is a punctuation, replace it with -1 (which removes it)
			return -1
		}
		// Otherwise, keep the rune
		return r
	}, s)
}
