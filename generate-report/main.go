package main

import (
	"html/template"
	"log/slog"
	"os"
)

type AItest struct {
	PromptPath string
	Question   string
	Answers    Answers
}

type Answers []Answer

type Answer struct {
	Name  string
	Value string
}

func main() {

	const tmpl = `
<h1>Question</h1>
<p>{{.Question}}</p>
<h1>Answers</h1>
{{- range .Answers}}
<p>{{.Name}}: {{.Value}}</p>
{{- end}}
`

	t := template.Must(template.New("").Parse(tmpl))
	err := t.Execute(os.Stdout, AItest{
		PromptPath: "path/to/prompt",
		Question:   "What is the capital of France?",
		Answers: Answers{
			Answer{
				Name:  "A",
				Value: "Paris",
			},
		},
	})
	if err != nil {
		slog.Error("executing template", "err", err)
	}
}
