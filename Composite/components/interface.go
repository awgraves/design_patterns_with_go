package components

type TextElement interface {
	//Print displays its contents as a string to stdout
	Print()
	//addChild appends a new child to this node
	AddChild(t TextElement) error
}
