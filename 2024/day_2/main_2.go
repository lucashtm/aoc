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

    if report_is_safe(values) {
      count++
      continue
    }

    skip := 0
    found_safe := false
    for skip < len(values) && !found_safe {
      new_values := make([]int, len(values) - 1)
      i := 0
      j := 0
      for i < len(new_values) {
        if j == skip {
          j++
          continue
        }
        new_values[i] = values[j]
        i++
        j++
      }
      found_safe = report_is_safe(new_values)
      skip++
    }
    if found_safe {
      count++
    }
  }

  fmt.Printf("There are %d safe levels\n", count)
}

func report_is_safe(report []int) bool {
  fmt.Print("Report: ", report)
  // Check initial direction
  var direction int
  if report[1] > report[0] {
    direction = 1
  } else if report[1] < report[0] {
    direction = -1
  } else {
    fmt.Println(" is unsafe")
    return false
  }

  current := 1
  last := 0
  for current < len(report) {
    if levels_are_safe(report[current], report[last], direction) {
      current++
      last++
      continue
    }

    fmt.Println(" is unsafe")
    return false
  }
  fmt.Println(" is safe")
  return true
}

func levels_are_safe(level_1 int, level_2 int, direction int) bool {
  difference := (level_1 - level_2) * direction
  same_direction := difference > 0
  within_limit := difference <= max_difference

  return same_direction && within_limit
}
