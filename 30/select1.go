package main

import (
  "fmt"
  "time"
)

func sleepyGopher(id int, c chan int) {
  time.Sleep(time.Second * 3)
  fmt.Println("...", id, " snore ...")
  c <- id
}

func main() {
  c := make(chan int)
  timeout := time.After(2 * time.Second)
  for i := 0; i < 5; i++ {
    go sleepyGopher(i, c)
  }
  for i := 0; i < 5; i++ {
    select {
      case gopherID := <-c:
        fmt.Println("gopher ", gopherID, " has finished sleeping")
      case <-timeout:
        fmt.Println("my patience ran out")
        return
    }
  }
}
