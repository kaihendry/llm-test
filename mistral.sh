gptscript --list-models https://api.mistral.ai/v1
for i in $(ls -1 *.gpt | sort -h); do
    echo -n "mistral $i "
    gptscript --quiet=true --default-model "mistral-large-latest from https://api.mistral.ai/v1" $i >/tmp/$i.mistral
    test "$i" == 1.gpt && cat /tmp/$i.mistral | go run 1/main.go
    gptscript --quiet=true test/assert.gpt --input "Given the query $i, does the /tmp/$i.mistral meet it?"
    cat /tmp/$i.mistral
done
