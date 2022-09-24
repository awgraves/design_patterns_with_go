package components

import (
	"fmt"
	"strings"
)

//sentence is a composite TextElement that contains children leaf nodes
type sentence struct {
	children []TextElement
}

//NewSentence converts a raw sentence string into an initialized sentence
//instance that contains child word leaf nodes.
func NewSentence(raw string) *sentence {
	cleaned := strings.Trim(raw, " ")
	cleaned = strings.TrimRight(cleaned, ".")
	rws := strings.Split(cleaned, " ")

	s := sentence{
		children: []TextElement{},
	}
	for _, rw := range rws {
		w := NewWord(rw)
		err := s.AddChild(w)
		if err != nil {
			panic(err)
		}
	}

	return &s
}

//Print recursively calls Print on all the sentence's children (words).
//It also adds spaces in between each element and terminates with a period for punctuation (.)
func (s *sentence) Print() {
	for i, c := range s.children {
		c.Print()
		if i < len(s.children)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print(".")
}

//DiagramPrint is similar to Print except the output is formatted to display as a tree
func (s *sentence) DiagramPrint() {
	for i, c := range s.children {
		fmt.Printf("\tW%d: ", i+1)
		c.DiagramPrint()
		fmt.Printf("\n")
	}
}

func (s *sentence) AddChild(t TextElement) error {
	s.children = append(s.children, t)
	return nil
}
