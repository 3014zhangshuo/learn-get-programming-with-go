package main

import (
  "fmt"
  "math/rand"
)

func main() {
  const secondsPerDay = 24 * 60 * 60
  const distance = 62100000 // km

  fmt.Println("太空航空公司      飞行天数   飞行类型   价格（百万美元）")

  company := ""
  trip := ""

  for count := 10; count > 0; count-- {
    switch rand.Intn(3) {
    case 0:
      company = "Space Adventures"
    case 1:
      company = "SpaceX"
    case 2:
      company = "Virgin Galactic"
    }

    speed := rand.Intn(15) + 16 // 16 ~ 30 km/s
    duration := distance / speed / secondsPerDay
    price := 20.0 + speed

    switch rand.Intn(2) {
    case 0:
      trip = "单程"
    case 1:
      trip = "往返"
      price *= 2
    }

    fmt.Printf("%-16v %6v %9v %12v\n", company, duration, trip, price)
  }
}
