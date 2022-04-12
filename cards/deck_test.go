package main

import (
	"os"
	"testing"
)

func TestNewDeckLength(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected length of 16, but actual is %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected Ace of Clubs, but actual is %v", d[0])
	}

	if d[len(d)-1] != "Four of Spades" {
		t.Errorf(("Expected Four of Spades, but actual is %v"), d[len(d)-1])
	}

}

func TestSaveToDeck(t *testing.T) {
	filename := "_deckTesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length %d, actual: %v", 16, len(loadedDeck))
	}

	os.Remove(filename)
}
