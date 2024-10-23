package main

func main() {

	cards := newDeck()

	//cards.toString()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}

// go run main.go deck.go
