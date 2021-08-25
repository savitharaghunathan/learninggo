package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newdeck()

	if len(cards) != 56 {
		t.Errorf("Expected deck length of 52, instead got %v", len(cards))
	}

	if cards[0] != "Ace Spades" {
		t.Errorf("Expected first card as Ace Spaces, but got %v", cards[0])
	}

	if cards[len(cards)-1] != "King Hearts" {
		t.Errorf("Expected last card as King Hearts, but got %v", cards[len(cards)-1])
	}
}
