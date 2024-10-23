package main

func main() {

	cards := newDeck()

	//cards.toString()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	new_cards := newDeckFromFile("my")
	new_cards.print()

}

// go run main.go deck.go
