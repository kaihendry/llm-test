gptscript --list-models
for i in $(ls -1 *.gpt | sort -h); do
    echo -n "openai $i "
    gptscript --quiet=true --default-model "gpt-4-turbo-preview" $i >/tmp/$i.openai
    gptscript --quiet=true test/assert.gpt --name "Given the query $i, does the /tmp/$i.openai meet it?"
done
