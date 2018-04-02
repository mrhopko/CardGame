package controller

import (
	"html/template"
	"net/http"
)

type home struct {
	homeTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	h.homeTemplate.Execute(w, nil)
}
