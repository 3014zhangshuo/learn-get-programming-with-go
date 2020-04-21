package main

import (
  "fmt"
  "strings"
)

type talker interface {
  talk() string
}

func shout(t talker) {
  louder := strings.ToUpper(t.talk())
  fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
  return "nack nack"
}

type laser int

func (l laser) talk() string {
  return strings.Repeat("toot ", int(l))
}

type rover string

func (r rover) talk() string {
  return string(r)
}

type starship struct {
  laser
}

func main() {
  s := starship{laser(3)}

  fmt.Println(s.talk())
  shout(s)

  r := rover("whir whir")
  shout(r)
}
