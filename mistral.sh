gptscript --list-models https://api.mistral.ai/v1
model="mistral-large-latest from https://api.mistral.ai/v1"
suffix=mistral

for i in $(ls -1 *.gpt | sort -h); do
    gptscript --quiet=true --default-model "$model" $i >/tmp/$i.${suffix}
    test "$i" == 1.gpt && cat /tmp/$i.${suffix} | go run 1/main.go | tee /tmp/test.${i}.${suffix}
    test -f /tmp/test.${i}.${suffix} || gptscript --quiet=true test/assert.gpt --name "Given the query $i, does the /tmp/$i.${suffix} meet it?" | tee /tmp/test.$i.${suffix}
done
