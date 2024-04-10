package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

type Questions struct {
	AItests     []AItest
	Generated   time.Time
	Leaderboard []Score
}

type Score struct {
	Model string
	Value int
}

type AItest struct {
	PromptPath string
	Question   string
	Answers    Answers
}

type Answers []Answer

type Answer struct {
	Name       string
	Value      string
	Assertions []Assertion
}

type Assertion struct {
	Name        string // what asserted the answer
	Ok          bool   // ok or not ok
	Description string // description https://en.wikipedia.org/wiki/Test_Anything_Protocol
}

func ReadFileToString(filePath string) (string, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Convert bytes to string and return
	content := string(bytes)
	return content, nil
}

func main() {
	qs := Questions{
		Generated: time.Now(),
	}

	for i := 1; i < 13; i++ {
		prompt := fmt.Sprintf("../%d.gpt", i)
		question, err := ReadFileToString(prompt)
		if err != nil {
			panic(err)
		}
		answer, err := ReadFileToString(fmt.Sprintf("../%d.answer", i))
		if err != nil {
			panic(err)
		}
		openaiAnswer, err := ReadFileToString(fmt.Sprintf("/tmp/%d.gpt.openai", i))
		if err != nil {
			panic(err)
		}
		mistralAnswer, err := ReadFileToString(fmt.Sprintf("/tmp/%d.gpt.mistral", i))
		if err != nil {
			panic(err)
		}
		anthropicAnswer, err := ReadFileToString(fmt.Sprintf("/tmp/%d.gpt.anthropic", i))
		if err != nil {
			panic(err)
		}

		sanityAssertion, err := parseTAP("Sanity check", fmt.Sprintf("/tmp/test.%d.answer", i))
		if err != nil {
			panic(err)
		}

		openAIAssertion, err := parseTAP("assert.gpt", fmt.Sprintf("/tmp/test.%d.gpt.openai", i))
		if err != nil {
			panic(err)
		}

		mistralAssertion, err := parseTAP("assert.gpt", fmt.Sprintf("/tmp/test.%d.gpt.mistral", i))
		if err != nil {
			panic(err)
		}

		anthropicAssertion, err := parseTAP("assert.gpt", fmt.Sprintf("/tmp/test.%d.gpt.anthropic", i))
		if err != nil {
			panic(err)
		}

		qs.AItests = append(qs.AItests, AItest{
			// base name of the prompt file
			PromptPath: path.Base(prompt),
			Question:   question,
			Answers: Answers{
				Answer{
					Name:  "Sanity check with OpenAI on correct answer",
					Value: answer,
					Assertions: []Assertion{
						sanityAssertion,
					},
				},
				Answer{
					Name:  "OpenAI",
					Value: openaiAnswer,
					Assertions: []Assertion{
						openAIAssertion,
					},
				},
				Answer{
					Name:  "Mistral",
					Value: mistralAnswer,
					Assertions: []Assertion{
						mistralAssertion,
					},
				},
				Answer{
					Name:  "Anthropic",
					Value: anthropicAnswer,
					Assertions: []Assertion{
						anthropicAssertion,
					},
				},
			},
		})
	}

	// Compute the scope of Sanity check, OpenAI, Mistral, and Anthropic
	var score = map[string]int{}

	for _, aiTest := range qs.AItests {
		for _, answer := range aiTest.Answers {
			for _, assertion := range answer.Assertions {
				if assertion.Ok {
					score[answer.Name]++
				}
			}
		}
	}

	// concert the map to a sorted slice
	for name, value := range score {
		qs.Leaderboard = append(qs.Leaderboard, Score{
			Model: name,
			Value: value,
		})
	}
	sort.Slice(qs.Leaderboard, func(i, j int) bool {
		return qs.Leaderboard[i].Value > qs.Leaderboard[j].Value
	})

	err := generateReport(qs)
	if err != nil {
		panic(err)
	}
}

func parseTAP(name, filePath string) (a Assertion, err error) {
	assertionText, err := ReadFileToString(filePath)
	if err != nil {
		return
	}
	// split on -
	testLine := strings.Split(assertionText, "-")
	if len(testLine) < 1 {
		err = fmt.Errorf("%s: expected at least 1 part in %s, actual %d", filePath, assertionText, len(testLine))
		return
	}
	if strings.HasPrefix(testLine[0], "ok") {
		a.Ok = true
	} else if strings.HasPrefix(testLine[0], "not ok") {
		a.Ok = false
	} else {
		err = fmt.Errorf("%s: expected ok or not ok in %s", filePath, assertionText)
		return
	}
	a.Description = strings.TrimSpace(strings.Join(testLine[1:], "-"))
	a.Name = name
	return
}

func generateReport(qs Questions) error {
	// use index.gohtml
	var tmplFile = "index.gohtml"
	// check template file exists
	if _, err := os.Stat(tmplFile); os.IsNotExist(err) {
		return err
	}
	t := template.Must(template.New(tmplFile).ParseFiles(tmplFile))
	err := t.Execute(os.Stdout, qs)
	return err
}
