package main

import "fmt"

func main() {
	//var card string = "Ace of Spades" //variable name type
	cards := []string{"Ace of Diamonds", newCard()} //This is the same as above, but ur making go infer its a string
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
