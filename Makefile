all: list-openai-models list-mistral-models openai mistral

list-mistral-models:
	gptscript --list-models https://api.mistral.ai/v1

list-openai-models:
	gptscript --list-models

openai:
	@for i in $$(ls -v *.gpt); \
	do \
		echo gptscript --model "gpt-4-turbo-preview" $$i; \
		gptscript --model "gpt-4-turbo-preview" $$i; \
	done

mistral:
	@for i in $$(ls -v *.gpt); \
	do \
		echo gptscript --default-model "mistral-large-latest from https://api.mistral.ai/v1" $$i; \
		gptscript --default-model "mistral-large-latest from https://api.mistral.ai/v1" $$i; \
	done
