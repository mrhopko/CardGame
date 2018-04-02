package controller

import (
	"html/template"
)

var (
	homeController home
	dealController deal
	deckController deck
)

func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	dealController.dealTemplate = templates["deal.html"]
	deckController.deckTemplate = templates["deck.html"]

	homeController.registerRoutes()
	dealController.registerRoutes()
	deckController.registerRoutes()

}
