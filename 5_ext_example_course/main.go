package main

import "fmt"

func main() {

	cards := newDeck()

	hand, remainingCards := deal(cards, 3)

	hand.print()

	fmt.Println("###")

	remainingCards.print()

}

// go run main.go deck.go
