package main

import "fmt"

func main() {
  var nowhere *int

  fmt.Println(nowhere)
  if nowhere != nil {
    fmt.Println(*nowhere)
  }
}
