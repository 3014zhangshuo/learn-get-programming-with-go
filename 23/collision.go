package main

import "fmt"

type sol int

func (s sol) days(s2 sol) int {
  days := int(s2 - s)
  if days < 0 {
    days = -days
  }
  return days
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

func (l location) days(l2 location) int {
  return 5
}

type celsius float64

type report struct {
  sol
  location
  temperature
}

func (r report) days(s2 sol) int {
  return r.sol.days(s2)
}

func main() {
  report := report{sol: 15}

  fmt.Println(report.days(1446))
}
