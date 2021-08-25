package main

func main() {

	cards := newdeck()
	//cards := newdeck()
	//cards.write2File("mycards")
	cards.print()

	cards.shuffle()
	cards.print()
}
