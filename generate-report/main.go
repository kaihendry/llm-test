package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

type Questions struct {
	AItests   []AItest
	Generated time.Time
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

	for i := 1; i <= 12; i++ {
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

		openAIAssertion, err := parseTAP("Sanity check", fmt.Sprintf("/tmp/test.%d.gpt.openai", i))
		if err != nil {
			panic(err)
		}

		qs.AItests = append(qs.AItests, AItest{
			// base name of the prompt file
			PromptPath: path.Base(prompt),
			Question:   question,
			Answers: Answers{
				Answer{
					Name:  "Correct",
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
				},
				Answer{
					Name:  "Anthropic",
					Value: anthropicAnswer,
				},
			},
		})
	}

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
	if len(testLine) != 2 {
		err = fmt.Errorf("%s: expected 2 parts in %s, actual %d", filePath, assertionText, len(testLine))
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
	a.Description = strings.TrimSpace(testLine[1])
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
