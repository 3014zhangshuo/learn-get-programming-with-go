package main

import (
  "fmt"
)

func main() {
  cipherText := "CSOITEUIWUIZNSROCNKFD"
  keyword := "GOLANG"
  message := ""
	keyIndex := 0

  for i := 0; i < len(cipherText); i++ {
    // 这里减去 'A' 的目的是初始化 A 的值为 0 A-Z = 0-25
    c := cipherText[i] - 'A'
    k := keyword[keyIndex] - 'A'

    // c - k 可能为复数，这时候要回绕
    fmt.Println(c-k+26)
    c = (c-k+26)%26 + 'A'

    message += string(c)

    keyIndex++
		keyIndex %= len(keyword)
  }

  fmt.Println(message)
}
