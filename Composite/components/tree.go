package components

import "strings"

type textTree struct {
	children []TextElement
}

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
		t.AddChild(s)
	}

	return &t
}

func (t *textTree) Print() error {
	for _, c := range t.children {
		c.Print()
	}
	return nil
}

func (t *textTree) AddChild(c TextElement) {
	t.children = append(t.children, c)
}
