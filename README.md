# segwise-backend-test
Backend Intern Test for Segwise.ai

## Resources
* [This shows how to use llm with api-endpoints](https://go.dev/blog/llmpowered)
* [How to run LLM's locally](https://medium.com/@rifai201/talkative-your-golang-gateway-to-powerful-large-language-models-7577814bb7c3
)
* [Ollama](https://github.com/ollama/ollama) (Fun Fact, Ollama is written in GO :))


## Dependencies
* [Ollama](https://github.com/ollama/ollama)
    * Installation instructions at repo's README
    * I am using OSX
    * We will use the basic model *Llama 3.2*
    * The command for it is `ollama run llama3.2`
    * To  start the Ollama server locally, just run the command `ollama server`
    * We will intereact with  it using out Go app, not the shell, so need for `ollama run llama3.2`
