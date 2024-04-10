for i in $(ls -v *.gpt); do
	answer=$(basename $i .gpt).answer
	gptscript test/assert.gpt --name "Given the query file $i, does the file ${answer} meet it?" | tee /tmp/test.${answer}
done
