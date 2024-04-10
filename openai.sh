gptscript --list-models
model="gpt-4-turbo-preview"
suffix=openai

for i in $(ls -1 *.gpt | sort -h); do
    gptscript --dump-state /tmp/dump.$i.${suffix} --quiet=true --default-model "$model" $i >/tmp/$i.${suffix}
    test "$i" == 1.gpt && cat /tmp/$i.${suffix} | go run 1/main.go | tee /tmp/test.${i}.${suffix}
    test -f /tmp/test.${i}.${suffix} || gptscript --quiet=true test/assert.gpt --name "Given the query $i, does the /tmp/$i.${suffix} meet it?" | tee /tmp/test.$i.${suffix}
done
