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
	// if any word is 18, then we pass
	for _, word := range words {
		if word == "18" || word == "eighteen" {
			fmt.Println("ok 10")
			return
		}
	}
	fmt.Println("not ok 10 - no word is 18")
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
