package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func ollamaHealthCheck() bool {
	resp, err := http.Get(ollamaURL)
	if err != nil {
		log.Println("HTTP call failed:", err)
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
func ollamaStart(){
		cmd := exec.Command("ollama", "serve")
		err := cmd.Run()
		if err != nil{
			log.Fatal("err: ", err.Error())
			log.Fatal("High chance that 'ollama' is not installed")
			os.Exit(1)
		}
}
