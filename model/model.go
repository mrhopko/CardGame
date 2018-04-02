package model

import (
	"context"
	"encoding/csv"
	"io"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

// Card Struct for holding Card
type Card struct {
	Category string
	Title    string
	Picture  string
}

// CardHand struct representing a players hand
type CardHand struct {
	Cards []*Card
}

// Deck struct representing a complete Deck
type Deck struct {
	Cards []*Card
}

// CSVToDeck Convert a CSV to a Deck
func CSVToDeck(reader csv.Reader) Deck {

	var cards []*Card
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			print(error)
		}
		cards = append(cards, &Card{
			Category: line[0],
			Title:    line[1],
			Picture:  "picture",
		})
	}

	result := Deck{Cards: cards}
	return result
}

// SaveDeck Save a Deck to datastore
func SaveDeck(ctx context.Context, d *Deck) {

	key := datastore.NewIncompleteKey(ctx, "Deck", nil)

	if _, err := datastore.Put(ctx, key, &d); err != nil {
		log.Errorf(ctx, "datastore.Put: %v", err)
		return
	}

	return
}
