package main

import "fmt"

func main() {

	cards := newDeck()

	hand, remainingCards := deal(cards, 3)

	hand.print()

	fmt.Println("###")

	remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

}

// go run main.go deck.go
