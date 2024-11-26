package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  current := 0
  for i := 0; i < len(text); i++ {
    if text[i] == '(' {
      current += 1
    }
    if text[i] == ')' {
      current -= 1
    }
    if current == -1 {
      fmt.Printf("index: %d\n", i)
    }
  }
  fmt.Printf("%d\n", current)
}
