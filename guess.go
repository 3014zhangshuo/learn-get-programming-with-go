package main

import (
  "fmt"
  "math/rand"
)

func main() {
  num := 4

  min := 1
  max := 100

  for {
    var randNum = rand.Intn(max-min) + min
    fmt.Printf("guess is %v \n", randNum)

    if (num == randNum) {
      break
    } else if (randNum > num) {
      fmt.Println("guess is big")
    } else {
      fmt.Println("guess is small")
    }
  }
}
