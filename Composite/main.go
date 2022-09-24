package main

import "github.com/awgraves/design_patterns_with_go/Composite/components"

func main() {
	tree := components.NewTextTree("Hello world composite example. And another sentence.")
	tree.Print()
}
