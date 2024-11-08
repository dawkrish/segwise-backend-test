# segwise-backend-test
Backend Intern Test for Segwise.ai


## Run this repo
* use Makefile.
* It only has one command.
* run `make` or `make server`
* we start Ollama by `ollama server &> /dev/null &`. Piping the stderr and stdout to `/dev/null` and `&`to run in background 
* We have functions `ollamaHealthCheck` &  `ollamaStart`.  They don't need to be there, but its just a way to make our code safer

## Resources
* [This shows how to use llm with api-endpoints](https://go.dev/blog/llmpowered)
* [How to run LLM's locally](https://medium.com/@rifai201/talkative-your-golang-gateway-to-powerful-large-language-models-7577814bb7c3
)
* [Ollama](https://github.com/ollama/ollama) (Fun Fact, Ollama is written in GO :))
* [Running shell commands](https://www.sohamkamani.com/golang/exec-shell-command/)
* I used ChatGPT, to find about SerpAPI, it really awesome. You can scrape google-search content with it
* I have not used ChatGPT code. just discovered the API by it by the prompts "google api for play store"
* [SerpAPI](https://serpapi.com/google-play-api)


## Dependencies
* [Ollama](https://github.com/ollama/ollama)
    * Installation instructions at repo's README
    * I am using OSX
    * We will use the basic model *Llama 3.2*
    * The command for it is `ollama run llama3.2`
    * To  start the Ollama server locally, just run the command `ollama server`
    * We will intereact with  it using out Go app, not the shell, so need for `ollama run llama3.2`
