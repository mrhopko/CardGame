package cardgame

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Card struct {
	Category string
	Title    string
	Picture  string
}

type CardHand struct {
	Cards []*Card
}

func init() {
	r := newRouter()
	http.Handle("/", r)
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {

	t := populateTemplates()

	Card1 := Card{Category: "Card 1 Category", Title: "Card 1 Title"}
	Card2 := Card{Category: "Card 2 Category", Title: "Card 2 Title"}
	Card3 := Card{Category: "Card 3 Category", Title: "Card 3 Title"}
	PlayerCards := CardHand{Cards: []*Card{&Card1, &Card2, &Card3}}

	t["layout"].Execute(w, PlayerCards)

}

func populateTemplates() map[string]*template.Template {
	const viewPath = "CardGame/view/"
	t := make(map[string]*template.Template)
	layout := template.Must(template.ParseFiles(viewPath + "layout.html"))
	template.Must(layout.ParseFiles(viewPath + "cardhand.html"))
	t["layout"] = layout
	return t
}
