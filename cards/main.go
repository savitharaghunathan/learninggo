package main

import "fmt"

func main() {

	cards := newdeck()

	fmt.Println(cards)
	fmt.Println("------- Learning how to iterate ---------")

	cards.print()
}
