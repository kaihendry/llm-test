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
	lastWord := words[len(words)-1]
	if lastWord != "cat" {
		fmt.Println("not ok 3 - last word is not cat")
		return
	}
	// check no word is four letters long
	for _, word := range words {
		if len(word) == 4 {
			fmt.Printf("not ok 3 - %s is four letters long\n", word)
			return
		}
	}
	fmt.Println("ok 3")
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
