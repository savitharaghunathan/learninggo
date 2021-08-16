package main

import "fmt"

func main() {
	cards := deck{"Ace Diamonds", "Ace Spades", newCard()}
	cards = append(cards, "Six Clover")

	fmt.Println(cards)
	fmt.Println("------- Learning how to iterate ---------")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace Hearts"
}
