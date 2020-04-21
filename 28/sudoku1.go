package main

import (
  "fmt"
  "errors"
  "os"
)

const (
  rows = 9
  columns = 9
)

type Grid [rows][columns]int8

func inBounds(row, column int) bool {
  if row < 0 || row > rows {
    return false
  }
  if column < 0 || column > columns {
    return false
  }
  return true
}

func (g *Grid) Set(row, column int, digit int8) error {
  if !inBounds(row, column) {
    return errors.New("out of bounds")
  }
  g[row][column] = digit
  return nil
}

func main() {
  var g Grid
  err := g.Set(10, 0, 5)

  if err != nil {
    fmt.Printf("An error occurred: %v.\n", err)
    os.Exit(1)
  }
}
