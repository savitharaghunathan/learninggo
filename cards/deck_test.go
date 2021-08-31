package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newdeck()

	if len(cards) != 56 {
		t.Errorf("Expected deck length of 56, instead got %v", len(cards))
	}

	if cards[0] != "Ace Spades" {
		t.Errorf("Expected first card as Ace Spaces, but got %v", cards[0])
	}

	if cards[len(cards)-1] != "King Hearts" {
		t.Errorf("Expected last card as King Hearts, but got %v", cards[len(cards)-1])
	}
}

func TestWrites2FileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newdeck()
	deck.write2File("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 56 {
		t.Errorf("Expected deck length of 56, instead got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
