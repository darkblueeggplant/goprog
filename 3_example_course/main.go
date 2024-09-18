package main

import "fmt"

// slice
func main() {

	cards := []string{newCard(), newCard(), "Ace of Diamonds"}

	fmt.Println(cards)

	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	for i, cards := range cards {
		fmt.Println(i, cards)
	}

}

func newCard() string {
	return "Ace of Spades"
}
