# Composite Design Pattern

## Intent

"Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly." - 'Gang of Four' Design Patterns, pg. 163

## Elements

- Component (`TextElement` in interface.go)

  - interface for compositions
  - all other nodes should adhere to this

- Leaf (`word` in word.go)

  - is the edge of the tree (has no children)
  - base building blocks

- Composite (`textTree` in tree.go, `sentence` in sentence.go)

  - stores child components
  - implements child-related operations of the Component interface

- Client (main.go) manipulates objects in the composition through the Component interface

## Exploratory Implementation

Rudimentary text parsing into a hierarchical tree.

`textTree` is initialized by feeding it a raw string of multiple sentences separated by period for punctuation.

Two methods can be called on the tree, which will recursively call the same interface method the whole way down to the leaves.

1. `tree.Print()` will print the text to stdout like normal sentences.
2. `tree.DiagramPrint()` will output the text broken down by sentence and by word within that sentence.

![Screen Shot 2022-09-24 at 3 11 39 PM](https://user-images.githubusercontent.com/34633964/192114738-1e476a4e-8628-42c8-9f13-50860b245b79.png)
