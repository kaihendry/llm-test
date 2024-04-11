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
	words := strings.Fields(nopunctuation)
	firstWordLen := len(words[0])
	lastWordLen := len(words[len(words)-1])

	// check first word is the longest
	for _, word := range words {
		if len(word) > firstWordLen {
			fmt.Printf("not ok 2 - %s is longer than %s\n", word, words[0])
			return
		}
	}

	// check last word is the shortest
	for _, word := range words {
		if len(word) < lastWordLen {
			fmt.Printf("not ok 2 - %s is shorter than %s\n", word, words[len(words)-1])
			return
		}
	}

	fmt.Println("ok 2")
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
