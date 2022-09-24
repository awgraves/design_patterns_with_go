package components

import "strings"

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

// Print recursively calls Print on all the tree's children (sentences).
func (t *textTree) Print() error {
	for _, c := range t.children {
		c.Print()
	}
	return nil
}

func (t *textTree) AddChild(c TextElement) error {
	t.children = append(t.children, c)
	return nil
}
