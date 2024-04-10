gptscript --list-models github.com/gptscript-ai/anthropic-provider
model="claude-3-haiku-20240307 from github.com/gptscript-ai/anthropic-provider"
suffix=anthropic

for i in $(ls -1 *.gpt | sort -h); do
    gptscript --quiet=true --default-model "$model" $i >/tmp/$i.${suffix}
    test "$i" == 1.gpt && cat /tmp/$i.${suffix} | go run 1/main.go | tee /tmp/test.${i}.${suffix}
    test -f /tmp/test.${i}.${suffix} || gptscript --quiet=true test/assert.gpt --name "Given the query $i, does the /tmp/$i.${suffix} meet it?" | tee /tmp/test.$i.${suffix}
done
