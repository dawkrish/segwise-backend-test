package main

import (
	"log"
	"net/http"
)

const (
	ollamaURL = "http://localhost:11434"
	serverHost = "localhost"
	serverPort = "8080"
)

func main() {
	if ollamaHealthCheck() == false{
		ollamaStart()
	}
	cfg := NewConfig()

	cfg.mux.HandleFunc("/", cfg.homeHandler)	

	log.Printf("Server started at http://%v:%v\n", serverHost, serverPort)
	log.Fatal(http.ListenAndServe(serverHost + ":" + serverPort, cfg.mux))
}



