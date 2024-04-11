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
	for _, word := range words {
		if word == "2" || word == "two" {
			fmt.Println("ok 7")
			return
		}
	}
	fmt.Println("not ok 7 - no word is 2")
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
