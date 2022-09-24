package components

type TextElement interface {
	Print()
	AddChild(t TextElement) error
}
