package main

import "fmt"

type report struct {
  sol int
  temperature temperature
}

func (r report) average() celsius {
  return r.temperature.average()
}

type temperature struct {
  high, low celsius
}

func (t temperature) average() celsius {
  return (t.high + t.low) / 2
}

type celsius float64

func main() {
  t := temperature{high: -1.0, low: -78.0}
  report := report{sol: 15, temperature: t}

  fmt.Printf("average %vº C\n", report.average())
}
