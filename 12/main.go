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
	nopunctuation := removePunctuation(strings.ToLower(string(bytes)))
	words := strings.Fields(nopunctuation)
	// check for "3" "zaks" in sequence
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "3" && words[i+1] == "zaks" {
			fmt.Printf("ok 12 - %s\n", words)
			return
		}
	}
	fmt.Printf("not ok 12 - %s\n", words)
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
