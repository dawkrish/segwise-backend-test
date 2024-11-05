package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

/*
A safety function.
We explicity call `ollama serve` in our Makefile, still we check it
*/
func ollamaHealthCheck() bool {
	resp, err := http.Get(ollamaURL)
	if err != nil {
		log.Println(err)
		return false
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK HTTP status:", resp.StatusCode)
		return false
	}
	log.Println("Ollama is runing!")
	return true
}

/*
This function should be called if the function "ollamaHealthCheack" fails
*/
func ollamaStart() {
	log.Println("Let's start ollama")
	cmd := exec.Command("ollama", "serve")
	err := cmd.Start()
	if err != nil {
		log.Fatal("err: ", err.Error())
		log.Fatal("High chance that 'ollama' is not installed")
		os.Exit(1)
	}
	log.Println("Ollama started!")
}

func ollamaQuery(query string) string {
	endpoint := ollamaURL + "/api/generate"

	reqParams := struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
		Stream bool   `json:"stream"`
		Format string `json:"format"`
	}{
		Model:  "llama3",
		Prompt: query,
		Stream: false,
		Format: "json",
	}

	jsonBody, _ := json.Marshal(reqParams)
	// if err != nil {
	// 	log.Println(err)
	// 	return ""
	// }
	resp, err := http.Post(endpoint, "", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println(err)
		return ""
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	log.Println("Resp Status: ", resp.Status)

	responseParam := struct {
		Model     string `json:"model"`
		CreatedAt string `json:"created_at"`
		Response  string `json:"response"`
		Done      bool   `json:"done"`
	}{}

	err = json.Unmarshal(respBody, &responseParam)
	if err != nil {
		log.Println(err)
		return ""
	}
	resp.Body.Close()

	return responseParam.Response
}
