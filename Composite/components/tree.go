package components

import (
	"fmt"
	"strings"
)

//textTree is the root container for subesquent child TextElements
type textTree struct {
	children []TextElement
}

//NewTextTree takes a raw string of text, breaks it into child sentence composite TextElement nodes, then returns an initialized textTree to contain them.
func NewTextTree(raw string) *textTree {
	rawSentences := strings.Split(raw, ".")

	t := textTree{
		children: []TextElement{},
	}

	for _, rs := range rawSentences {
		if len(rs) <= 1 {
			continue
		}
		s := NewSentence(rs)
		err := t.AddChild(s)
		if err != nil {
			panic(err)
		}
	}

	return &t
}

// Print recursively calls Print on all the tree's children (sentences) with added spaces in between child prints for readability.
func (t *textTree) Print() {
	for i, c := range t.children {
		if i > 0 {
			fmt.Print(" ")
		}
		c.Print()
	}
}

//DiagramPrint is similar to Print except the output is formatted to display as a tree
func (t *textTree) DiagramPrint() {
	for i, c := range t.children {
		fmt.Printf("S%d:\n", i+1)
		c.DiagramPrint()
	}
}

func (t *textTree) AddChild(c TextElement) error {
	t.children = append(t.children, c)
	return nil
}
