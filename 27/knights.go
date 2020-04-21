package main

import "fmt"

type characeter struct {
  name string
  leftHand *item
}

type item struct {
  name string
}

func (c *characeter) pickup(i *item) {
  if c == nil || i == nil {
    return
  }
  fmt.Printf("%v picks up a %v\n", c.name, i.name)
  c.leftHand = i
}

func (c *characeter) give(to *characeter) {
  if c == nil || to == nil {
    return
  }
  if c.leftHand == nil {
    fmt.Printf("%v has nothing to give\n", c.name)
    return
  }
  if to.leftHand != nil {
    fmt.Printf("%v's hands are full\n", to.name)
		return
  }
  to.leftHand = c.leftHand
  c.leftHand = nil
  fmt.Printf("%v gives %v a %v\n", c.name, to.name, to.leftHand.name)
}

func (c characeter) String() string {
  if c.leftHand == nil {
    return fmt.Sprintf("%v is carrying nothing", c.name)
  }
  return fmt.Sprintf("%v is carrying a %v", c.name, c.leftHand.name)
}

func main() {
  arthur := &characeter{name: "Arthur"}

  shrubbery := &item{name: "shrubbery"}
  arthur.pickup(shrubbery)

  knight := &characeter{name: "Knight"}
  arthur.give(knight)

  fmt.Println(arthur)
  fmt.Println(knight)
}
