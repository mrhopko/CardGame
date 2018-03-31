package cardgame

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
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
	r.HandleFunc("/upload", ReceiveFile).Methods("POST")
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
	const viewPath = "view/"
	t := make(map[string]*template.Template)
	layout := template.Must(template.ParseFiles(viewPath + "layout.html"))
	template.Must(layout.ParseFiles(viewPath+"cardhand.html", viewPath+"uploaddeck.html"))
	t["layout"] = layout
	return t
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
	// in your case file would be fileupload
	r.ParseMultipartForm(32 << 20)

	file, header, err := r.FormFile("uploadfile")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(w, "%v", header.Header)

	reader := csv.NewReader(bufio.NewReader(file))
	var deck []Card
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		deck = append(deck, Card{
			Category: line[0],
			Title:    line[1],
			Picture:  "picture",
		})
	}
	Buf.Reset()

	t := populateTemplates()

	t["layout"].Execute(w, deck)
}
