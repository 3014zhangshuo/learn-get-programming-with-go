package main

import (
  "fmt"
  "math/rand"
)

func main() {
  var total float64 = 0

  for {
    switch rand.Intn(3) {
    case 0:
      total += 0.05
    case 1:
      total += 0.10
    case 2:
      total += 0.25
    }

    fmt.Printf("%5.2f\n", total)

    if (total >= 20.00) {
      break
    }
  }
}
