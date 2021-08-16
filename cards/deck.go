package main

import "fmt"

//create a new type 'deck'
//Slice of string type

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
