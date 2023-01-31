package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Error("Expected length of deck 52 but actual length ", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card \"Ace of Spades\" but actual card \"%v\"", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card \"King of Clubs\" but actual card \"%v\"", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Error("Expected length of deck 52 but actual length ", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
