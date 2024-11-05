package main

import (
	"html/template"
	"net/http"
)

type Config struct{
	mux *http.ServeMux
	templates *template.Template
}

func NewConfig() *Config{
	mux := http.NewServeMux()
	templates, err := template.ParseGlob("templates/*.html")
	if err != nil{
		panic(err)
	}

	cfg := Config{
		mux : mux,
		templates : templates,
	}
	return &cfg
}

