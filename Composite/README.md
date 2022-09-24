# Composite Design Pattern

## Intent

"Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly." - 'Gang of Four' Design Patterns, pg. 163

## Elements

- Component (TextElement in interface.go)

  - interface for compositions
  - all other nodes should adhere to this

- Leaf (word in word.go)

  - is the edge of the tree (has no children)
  - base building blocks

- Composite (textTree in tree.go, sentence in sentence.go)

  - stores child components
  - implements child-related operations of the Component interface

- Client (main.go) manipulates objects in the composition through the Component interface
