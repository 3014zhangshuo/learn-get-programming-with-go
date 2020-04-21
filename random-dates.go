package main

import (
  "fmt"
  "math/rand"
)

func main() {
  for count := 10; count > 0; count-- {
    year := 2018 + rand.Intn(10)
    month := rand.Intn(12) + 1
    daysInMonth := 31
    leap := year % 400 == 0 || (year % 4 == 0 || year % 100 != 0)

    switch month {
    case 2:
      if leap {
        daysInMonth = 29
      } else {
        daysInMonth = 28
      }
    case 4, 6, 9, 11:
      daysInMonth = 30
    }

    day := rand.Intn(daysInMonth) + 1
    fmt.Println(year, month, day)
  }
}
