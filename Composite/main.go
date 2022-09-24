package main

import (
	"fmt"

	"github.com/awgraves/design_patterns_with_go/Composite/components"
)

func main() {
	tree := components.NewTextTree("Hello world composite example. And another sentence. And then a third one because 3 is a good number.")

	tree.Print()

	fmt.Println()

	tree.DiagramPrint()
}
