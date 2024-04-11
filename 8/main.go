package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		sentence := scanner.Text()
		slog.Debug("testing", "sentence", sentence)
		if validSequence(sentence) {
			fmt.Println("ok 8")
			return
		}
	}
	fmt.Println("not ok 8")
}

func validSequence(sentence string) bool {
	if strings.Contains(sentence, "xxxxxxxxxx") || strings.Contains(sentence, "yyyyyyyyyy") || strings.Contains(sentence, "yyyxyyyyyy") {
		return true
	}
	return false
}
