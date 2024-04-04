openai: list-openai-models openai-gptscript

mistral: list-mistral-models mistral-gptscript

anthropic: list-anthropic-models anthropic-gptscript

list-mistral-models:
	gptscript --list-models https://api.mistral.ai/v1

list-openai-models:
	gptscript --list-models

openai-gptscript:
	@for i in $$(ls -v *.gpt); \
	do \
		gptscript --debug --default-model "gpt-4-turbo-preview" $$i; \
		printf "\033[1mExpected answer:\033[0m\n"; \
		cat `basename $$i .gpt`.answer; \
	done

mistral-gptscript:
	@for i in $$(ls -v *.gpt); \
	do \
		gptscript --debug --default-model "mistral-large-latest from https://api.mistral.ai/v1" $$i; \
		printf "\033[1mExpected answer:\033[0m\n"; \
		cat `basename $$i .gpt`.answer; \
	done

list-anthropic-models:
	gptscript --list-models github.com/gptscript-ai/anthropic-provider

anthropic-gptscript:
	@for i in $$(ls -v *.gpt); \
	do \
		gptscript --debug --default-model='claude-3-haiku-20240307 from github.com/gptscript-ai/anthropic-provider' $$i; \
		printf "\033[1mExpected answer:\033[0m\n"; \
		cat `basename $$i .gpt`.answer; \
	done



