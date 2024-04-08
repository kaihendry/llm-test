gptscript --list-models github.com/gptscript-ai/anthropic-provider
for i in $(ls -1 *.gpt | sort -h); do
    echo -n "claude-3-haiku-20240307 $i "
    gptscript --quiet=true --default-model "claude-3-haiku-20240307 from github.com/gptscript-ai/anthropic-provider" $i >/tmp/$i.anthropic
    test "$i" == 1.gpt && cat /tmp/$i.anthropic | go run 1/main.go
    gptscript --quiet=true test/assert.gpt --input "Given the query $i, does the /tmp/$i.anthropic meet it?"
    cat /tmp/$i.anthropic
done
