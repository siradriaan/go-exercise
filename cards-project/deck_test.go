package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected length of deck is 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("the first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Hearts" {
		t.Errorf("expected first card to be Four of Hearts, but got %v", d[len(d)-1])
	}
}

func TestSaveNewDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of deck to be 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
