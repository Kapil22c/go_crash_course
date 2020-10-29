package main

func main() {
	// var card string = "Ace of Spades"
	// for new variable... declaration
	// card := "Ace of Spades"
	// card := newCard()
	// card = "Five of diamonds"
	// fmt.Println(card)
	// cards := deck{"Two of spades", newCard()}
	// cards := newDeck()
	// cards = append(cards, "eight of diamond")
	// fmt.Println(cards)
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// greeting := "My name is Kapil"
	// fmt.Prinln([]byte(greeting))
	cards := newDeck()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
