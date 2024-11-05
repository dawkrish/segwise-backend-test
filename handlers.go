package main

import (
	"net/http"
)


func (cfg *Config)homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct{
		Message string
	}{
		Message : "Segwise Test",
	}
	cfg.templates.ExecuteTemplate(w,"page", data)
}
