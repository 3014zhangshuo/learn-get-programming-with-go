package main

import "fmt"

type report struct {
  sol int
  temperature
  location
}

type temperature struct {
  high, low celsius
}

func (t temperature) average() celsius {
  return (t.high + t.low) / 2
}

type location struct {
  lat, long float64
}

type celsius float64

func main() {
  bradbury := location{-4.5895, 137.4417}
  t := temperature{high: -1.0, low: -78.0}
  report := report{sol: 15, temperature: t, location: bradbury}

  fmt.Printf("average %vº C\n", report.average())
}
