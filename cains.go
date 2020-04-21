package main

import (
  "fmt"
)

func main() {
  const distance = 236000000000000000 // km
  const lightSpeed = 299792
  const secondsPerDay = 86400
  const daysPerYear = 365

  const years = distance / lightSpeed / secondsPerDay / daysPerYear

  fmt.Println("Cains Major Dwarf Galaxy is", years, "light years away.")
}
