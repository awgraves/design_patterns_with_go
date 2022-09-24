package components

type TextElement interface {
	//addChild appends a new child to this node
	AddChild(t TextElement) error
	//Print displays its contents as a string to stdout
	Print()
	//DiagramPrint displays formatted content to stdout for tree visualization
	DiagramPrint()
}
