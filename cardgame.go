package cardgame

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

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
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	// Copy the file data to my buffer
	io.Copy(&Buf, file)
	// do something with the contents...
	// I normally have a struct defined and unmarshal into a struct, but this will
	// work as an example
	contents := Buf.String()
	fmt.Println(contents)
	// I reset the buffer in case I want to use it again
	// reduces memory allocations in more intense projects
	Buf.Reset()
	// do something else
	// etc write header
	return
}
