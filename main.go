package main

import (
	"log"
	"net/http"
)

const (
	ollamaURL = "http://localhost:11434/"
	serverHost = "localhost"
	serverPort = "8080"
)

func main() {
	if ollamaHealthCheck() == false{
		ollamaStart()
	}
	log.Fatal(http.ListenAndServe(serverHost + ":" + serverPort, nil))
}



