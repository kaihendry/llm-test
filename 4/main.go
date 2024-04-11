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
	for _, word := range words {
		// fail if any common english words found
		switch word {
		case "any", "the", "as", "gentle", "be", "to", "of", "a", "in", "that", "have", "i", "it", "for", "not":
			fmt.Printf("not ok 4 - common word %s found\n", word)
			return
		}
	}
	fmt.Println("ok 4 - no common words found")
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
