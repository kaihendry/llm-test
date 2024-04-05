for i in $(ls -v ../*.gpt); do
	echo gptscript assert.gpt --name "Given the query $i, does the ../$(basename $i .gpt).answer meet it?"
done
