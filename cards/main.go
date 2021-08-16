package main

import "fmt"

func main() {
	cards := deck{"Ace Diamonds", "Ace Spades", newCard()}
	cards = append(cards, "Six Clover")

	fmt.Println(cards)
	fmt.Println("------- Learning how to iterate ---------")

	cards.print()
}

func newCard() string {
	return "Ace Hearts"
}
