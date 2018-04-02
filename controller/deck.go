package controller

import (
	"html/template"
	"net/http"
)

type deck struct {
	deckTemplate *template.Template
}

func (d deck) registerRoutes() {
	http.HandleFunc("/deck", d.handleDeck)
}

func (d deck) handleDeck(w http.ResponseWriter, r *http.Request) {
	d.deckTemplate.Execute(w, nil)
}
