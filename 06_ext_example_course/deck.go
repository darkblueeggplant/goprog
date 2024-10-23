package main

import "fmt"

// create new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
