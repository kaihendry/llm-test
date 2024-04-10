package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Run isEndStart line be line on stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sentence := scanner.Text()
		slog.Debug("testing", "sentence", sentence)
		err := isEndStart(sentence)
		if err == nil {
			fmt.Println("ok")
		} else {
			fmt.Printf("not ok - %s\n", err.Error())
		}
	}
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

func isEndStart(sentence string) error {
	if len(sentence) == 0 {
		return fmt.Errorf("sentence is empty")
	}
	slog.Debug("input", "sentence", sentence)
	// split into lower case words without punctuation
	words := strings.Fields(strings.ToLower(removePunctuation(sentence)))
	// check if last letter of first work is the same as first letter of the next word
	for i := 0; i < len(words)-1; i++ {
		slog.Debug("testing", "word", words[i], "last letter", string(words[i][len(words[i])-1]), "next word", words[i+1], "first letter", string(words[i+1][0]))
		if words[i][len(words[i])-1] != words[i+1][0] {
			return fmt.Errorf("last letter of %q is not the same as the first letter of %q", words[i], words[i+1])
		}
	}
	return nil
}
