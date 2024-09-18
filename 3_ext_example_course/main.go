package main

import "fmt"

func main() {

	cards := deck{newCard(), newCard(), "Ace of Diamonds"}

	fmt.Println(cards)

	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	cards.print()

}

func newCard() string {
	return "Ace of Spades"
}

// go run main.go deck.go
