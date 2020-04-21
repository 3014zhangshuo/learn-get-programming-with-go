package main

import (
  "fmt"
)

func main() {
  plainText := "your message goes here"
  keyword := "GOLANG"

  plainText = strings.ToUpper(strings.Replace(plainText, " ", "", -1))
}
