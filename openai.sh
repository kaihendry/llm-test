gptscript --list-models
for i in $(ls -1 *.gpt | sort -h); do
    gptscript --quiet=true --default-model "gpt-4-turbo-preview" $i >/tmp/$i.openai
    test "$i" == 1.gpt && cat /tmp/$i.openai | go run 1/main.go | tee /tmp/test.${i}.openai
    test -f /tmp/test.${i}.openai || gptscript --quiet=true test/assert.gpt --name "Given the query $i, does the /tmp/$i.openai meet it?" | tee /tmp/test.$i.openai
done
