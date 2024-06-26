package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"log/slog"
)

type JSONData struct {
	Program struct {
		ToolSet map[string]struct {
			ModelName string `json:"modelName"`
		} `json:"toolSet"`
	} `json:"program"`
	Output string `json:"output"`
}

func parseJSONFromFile(filename string) (output, modelName string, err error) {
	// Read the file
	input, err := os.ReadFile(filename)
	if err != nil {
		return "", "", err
	}

	var data JSONData
	if err := json.Unmarshal(input, &data); err != nil {
		return "", "", err
	}

	// Assuming there's only one model in the toolSet.
	for _, tool := range data.Program.ToolSet {
		modelName = tool.ModelName
		break // Break after the first one since we just need one model name.
	}

	output = data.Output
	return output, modelName, nil
}

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
	slog.Info("Generating report")
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
		qs.AItests = append(qs.AItests, AItest{
			PromptPath: path.Base(prompt),
			Question:   question,
			Answers: Answers{
				Answer{
					Name:  "Correct answer",
					Value: answer,
				},
			},
		})
		modelTestPaths, err := filepath.Glob(fmt.Sprintf("/tmp/test.%d.gpt", i) + "*")
		if err != nil {
			panic(err)
		}
		for _, modelTestPath := range modelTestPaths {
			suffix := strings.Split(modelTestPath, ".")[3]
			slog.Info("fetching answer", "modelTestPath", modelTestPath, "suffix", suffix)
			qs.AItests[i-1].Answers = append(qs.AItests[i-1].Answers, fetchAnswer(modelTestPath))
		}
	}

	slog.Info("Computing leaderboard")

	var score = map[string]int{}

	for _, aiTest := range qs.AItests {
		for _, answer := range aiTest.Answers {
			for _, assertion := range answer.Assertions {
				if assertion.Ok {
					score[answer.Name]++
				} else {
					if _, ok := score[answer.Name]; !ok {
						score[answer.Name] = 0
					}
				}
			}
		}
	}

	for name, value := range score {
		if name == "Correct answer" {
			continue
		}
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

func fetchAnswer(modelTestPath string) Answer {
	parts := strings.Split(modelTestPath, ".")
	dumpPath := fmt.Sprintf("%s.%s.%s.%s", "/tmp/dump", parts[1], parts[2], parts[3])
	slog.Info("fetching answer", "modelTestPath", modelTestPath, "dumpPath", dumpPath)
	answer, modelName, err := parseJSONFromFile(dumpPath)
	if err != nil {
		panic(err)
	}

	assertion, err := parseTAP(modelTestPath)
	if err != nil {
		panic(err)
	}

	return Answer{
		Name:       modelName,
		Value:      answer,
		Assertions: []Assertion{assertion},
	}
}

func parseTAP(filePath string) (a Assertion, err error) {
	a.Name = filepath.Base(filePath)
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
	return
}

func generateReport(qs Questions) error {
	// use index.gohtml
	var tmplFile = "index.gohtml"
	// check template file exists
	if _, err := os.Stat(tmplFile); os.IsNotExist(err) {
		return err
	}
	t := template.Must(template.New(tmplFile).Funcs(template.FuncMap{
		"ext": filepath.Ext,
	}).ParseFiles(tmplFile))
	err := t.Execute(os.Stdout, qs)
	return err
}
