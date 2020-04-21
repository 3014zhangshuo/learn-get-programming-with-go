package main

import (
  "fmt"
  "math/rand"
)

func main() {
  total := 0
  target := 20 * 100

  for total < target {
    switch rand.Intn(3) {
    case 0:
      total += 5
    case 1:
      total += 10
    case 2:
      total += 25
    }

    dollars := total / 100
    cents := total % 100

    fmt.Printf("$ %d.%02d\n", dollars, cents)
  }
}
