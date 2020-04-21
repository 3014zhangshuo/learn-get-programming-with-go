package main

import (
  "fmt"
  "time"
  "image"
)

func main() {
  go worker()
  time.Sleep(5 * time.Second)
}

func worker() {
  pos := image.Point{X: 10, Y: 10}
  direction := image.Point{X: 1, Y: 0}
  delay := time.Second
  next := time.After(time.Second)
  for {
    select {
    case <-next:
      pos = pos.Add(direction)
      fmt.Println("current position is ", pos)
      delay += time.Second / 2
      next = time.After(delay)
    }
  }
}
