package main

import (
	"errors"
	"fmt"
)

type TextElement interface {
	Print()
	Add(t TextElement) error
}

type sentence struct {
	children []TextElement
}

type word struct {
	text string
}

func NewSentence() *sentence {
	return &sentence{
		children: []TextElement{},
	}
}

func (s *sentence) Print() {
	for i, c := range s.children {
		c.Print()
		if i < len(s.children)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print(".\n")
}

func (s *sentence) Add(t TextElement) error {
	s.children = append(s.children, t)
	return nil
}

func (s *sentence) GetChild(int) (TextElement, error) {
	return nil, errors.New("word cannot have children")
}

func (w *word) Print() {
	fmt.Print(w.text)
}

func (w *word) Add(t TextElement) error {
	return errors.New("word cannot have children")
}

func (w *word) GetChild(int) (TextElement, error) {
	return nil, errors.New("word cannot have children")
}

func main() {
	sentence := NewSentence()

	w1 := &word{"Hello"}
	w2 := &word{"world"}
	w3 := &word{"composite"}
	w4 := &word{"example"}

	words := []*word{w1, w2, w3, w4}
	for _, w := range words {
		sentence.Add(w)
	}

	sentence.Print()
}
