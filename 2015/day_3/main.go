package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  var directions = map[rune][2]int{
    '^': {0, -1},
    '>': {1, 0},
    'v': {0, 1},
    '<': {-1, 0},
  }

  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  visited := map[[2]int]bool{}
  santa_current := [2]int{0, 0}
  robo_current := [2]int{0, 0}
  total := 0
  flag := true
  for _, c := range text {
    if flag {
      if !visited[santa_current] {
        visited[santa_current] = true
        total++
      }
      santa_current[0] += directions[c][0]
      santa_current[1] += directions[c][1]
    } else {
      if !visited[robo_current] {
        visited[robo_current] = true
        total++
      }
      robo_current[0] += directions[c][0]
      robo_current[1] += directions[c][1]
    }
    flag = !flag
  }
  fmt.Printf("Total homes: %d\n", total)
}
