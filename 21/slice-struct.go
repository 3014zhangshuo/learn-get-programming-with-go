package main

import "fmt"

func main() {
  type location struct {
    name string
    lat float64
    long float64
  }

  locations := []location{
    {name: "A", lat: -4.5895, long: 137.4417},
    {name: "B", lat: -1.9426, long: 354.4743},
    {name: "C", lat: -14.5684, long: 175.472636},
  }

  fmt.Println(locations)
}
