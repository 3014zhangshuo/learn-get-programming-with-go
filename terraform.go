package main

import "fmt"

type Planets []string

func (planets Planets) terraform() {
  for i, plante := range planets {
    planets[i] = "New " + plante
  }
}

func main() {
  planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
  Planets(planets[3:4]).terraform()
  Planets(planets[6:]).terraform()
  fmt.Print(planets)
}
