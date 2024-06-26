#!/bin/bash

test_model() {
    local model="$1"
    local suffix="$2"

    for i in $(ls -1 *.gpt | sort -h); do
        gptscript --dump-state "/tmp/dump.$i.${suffix}" --quiet=true --default-model "$model" $i >"/tmp/$i.${suffix}"
        testNum=$(basename $i .gpt)
        if test -f $testNum/main.go; then
            cat "/tmp/$i.${suffix}" | go run ${testNum}/main.go | tee "/tmp/test.${i}.${suffix}".go
        else
            gptscript --quiet=true test/assert.gpt --name "Given the query $i, does the /tmp/$i.${suffix} meet it?" | tee "/tmp/test.$i.${suffix}".gpt4
        fi
    done
}

# For OpenAI
gptscript --list-models
test_model "gpt-4-turbo-preview" "openai"

# For Mistral
gptscript --list-models "https://api.mistral.ai/v1"
test_model "mistral-large-latest from https://api.mistral.ai/v1" "mistral"

# For Anthropics
gptscript --list-models "github.com/gptscript-ai/anthropic-provider"
test_model "claude-3-haiku-20240307 from github.com/gptscript-ai/anthropic-provider" "anthropic"

# https://endpoints.ai.cloud.ovh.net/
test_model "Mistral-7B-Instruct-v0.2 from https://mistral-7b-instruct-v02.endpoints.kepler.ai.cloud.ovh.net/api/openai_compat/v1" "ovh"
