package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	l := 52
	c1 := "Ace of Spades"
	c2 := "King of Diamonds"

	if len(d) != l {
		t.Errorf("Expected deck to contain %v cards. Found %v", l, len(d))
	}

	if d[0] != c1 {
		t.Errorf("Expected first card to be %v. Found %v", c1, d[0])
	}

	if d[len(d)-1] != c2 {
		t.Errorf("Expected last card to be %v. Found %v", c2, d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// clean up test artifacts
	fn := "decktesting.csv"
	os.Remove(fn)

	// save new deck to file
	deckToSave := newDeck()
	deckToSave.saveToFile(fn)

	// load deck from file
	d := newDeckFromFile(fn)

	l := 52
	c1 := "Ace of Spades"
	c2 := "King of Diamonds"

	if len(d) != l {
		t.Errorf("Expected loaded deck to contain %v cards. Found %v", l, len(d))
	}

	if d[0] != c1 {
		t.Errorf("Expected loaded deck first card to be %v. Found %v", c1, d[0])
	}

	if d[len(d)-1] != c2 {
		t.Errorf("Expected loaded deck last card to be %v. Found %v", c2, d[len(d)-1])
	}

	// clean up
	os.Remove(fn)
}
