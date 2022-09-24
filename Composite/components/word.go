package components

import (
	"errors"
	"fmt"
)

type word struct {
	text string
}

func NewWord(rw string) *word {
	return &word{
		text: rw,
	}
}

func (w *word) Print() {
	fmt.Print(w.text)
}

func (w *word) AddChild(t TextElement) error {
	return errors.New("word cannot have children")
}
