package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newdeck()

	if len(cards) != 56 {
		t.Errorf("Expected deck length of 52, instead got %v", len(cards))
	}
}
