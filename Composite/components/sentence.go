package components

import (
	"fmt"
	"strings"
)

type sentence struct {
	children []TextElement
}

func NewSentence(raw string) *sentence {
	cleaned := strings.Trim(raw, " ")
	cleaned = strings.TrimRight(cleaned, ".")
	rws := strings.Split(cleaned, " ")

	children := []TextElement{}
	for _, rw := range rws {
		w := NewWord(rw)
		children = append(children, w)
	}

	return &sentence{
		children: children,
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

func (s *sentence) AddChild(t TextElement) error {
	s.children = append(s.children, t)
	return nil
}
