package main

import (
  "encoding/json"
  "fmt"
  "log"
)

func main() {
  type location struct {
    Name string `json:"name"`
    Lat float64 `json:"lat"`
    Long float64 `json:"long"`
  }

  locations := []location{
    {Name: "A", Lat: -4.5895, Long: 137.4417},
    {Name: "B", Lat: -1.9426, Long: 354.4743},
    {Name: "C", Lat: -14.5684, Long: 175.472636},
  }

  bytes, err := json.MarshalIndent(locations, "", " ")
  if err != nil {
		log.Fatal(err)
	}

  fmt.Println(string(bytes))
}
