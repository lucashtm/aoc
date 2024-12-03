package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const max_difference = 3

func main() {
  reader := bufio.NewReader(os.Stdin)
  var count int
  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    line = strings.TrimSpace(line)
    if len(line) == 0 {
      break
    }
    str_values := strings.Split(line, " ")
    var values []int
    for _, str := range str_values {
      value, err := strconv.Atoi(str)
      if err != nil {
        log.Fatalf("Could not convert %s to an integer\n", str)
      }
      values = append(values, value)
    }

    // Check initial direction
    var direction int
    if values[1] > values[0] {
      direction = 1
    } else if values[1] < values[0] {
      direction = -1
    } else {
      fmt.Printf("%s is unsafe\n", line)
      continue
    }

    is_safe := true
    for i := 1; i < len(values); i++ {
      difference := (values[i] - values[i-1]) * direction
      if difference <= 0 || difference > max_difference  {
        is_safe = false
        break
      }
    }

    if !is_safe {
      fmt.Printf("%s is unsafe\n", line)
      continue
    }

    fmt.Printf("%s is safe\n", line)
    count++
  }

  fmt.Printf("There are %d safe levels\n", count)
}
