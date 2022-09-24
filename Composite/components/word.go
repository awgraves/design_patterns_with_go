package components

import (
	"errors"
	"fmt"
)

// word is a leaf TextElement with no children of its own.
type word struct {
	text string
}

//NewWord takes a raw string of a single word and converts it into a word leaf TextElement.
func NewWord(rw string) *word {
	return &word{
		text: rw,
	}
}

//Print displays the text string to stdout
func (w *word) Print() {
	fmt.Print(w.text)
}

func (w *word) AddChild(t TextElement) error {
	return errors.New("word cannot have children")
}
