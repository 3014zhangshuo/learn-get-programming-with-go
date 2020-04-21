package main

import "fmt"
import "strings"

type laser int

func (l *laser) talk() string {
  return strings.Repeat("pew ", int(*l))
}

type talker interface {
  talk() string
}

func shout(t talker) {
  louder := strings.ToUpper(t.talk())
  fmt.Println(louder)
}

func main() {
  pew := laser(2)
  shout(&pew)
}
