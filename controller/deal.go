package controller

import (
	"CardGame/model"
	"html/template"
	"net/http"
)

type deal struct {
	dealTemplate *template.Template
}

func (d deal) registerRoutes() {
	http.HandleFunc("/deal", d.handleDeal)
}

func (d deal) handleDeal(w http.ResponseWriter, r *http.Request) {

	Card1 := model.Card{
		Category: "cat1",
		Title:    "title1",
		Picture:  "pic1",
	}

	Card2 := model.Card{
		Category: "cat2",
		Title:    "title2",
		Picture:  "pic2",
	}

	Card3 := model.Card{
		Category: "cat3",
		Title:    "title3",
		Picture:  "pic3",
	}

	PlayerDeal := model.CardHand{Cards: []*model.Card{&Card1, &Card2, &Card3}}

	d.dealTemplate.Execute(w, PlayerDeal)
}
