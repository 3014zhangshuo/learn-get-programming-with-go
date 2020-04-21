package main

import "fmt"

func main() {
  var curiosity struct {
    lat float64
    long float64
    altitude float64
  }

  curiosity.lat = -4.5895
  curiosity.long = 137.4417
  curiosity.altitude = -4400

  fmt.Println(curiosity.lat, curiosity.long, curiosity.altitude)
  fmt.Println(curiosity)
}
