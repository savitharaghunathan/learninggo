package main

import "fmt"

func main() {

	cards := newdeck()
	fmt.Println("------- All Hands on Deck ---------")
	cards.print()
	fmt.Println("---------- After deal ------------")
	deal1, remainingDeal := deal(cards, 5)
	deal1.print()
	fmt.Println("--------- Remaining deck ----------")
	remainingDeal.print()

}
