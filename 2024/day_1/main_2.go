package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  var left_list []int
  var right_list []int

  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    values := strings.Split(strings.TrimSpace(line), "   ")
    left, err := strconv.Atoi(values[0])
    if err != nil {
      log.Fatalf("Could not parse %s to integer\n", values[0])
    }
    right, err := strconv.Atoi(values[1])
    if err != nil {
      log.Fatalf("Could not parse %s to integer\n", values[1])
    }

    left_list = append(left_list, left)
    right_list = append(right_list, right)
  }

  var result int
  for i := 0; i < len(left_list); i++ {
    
    count_current := 0
    for j := 0; j < len(right_list); j++ {
      if right_list[j] == left_list[i] {
        count_current++
      }
    }
    result += left_list[i] * count_current
  }
  fmt.Println(result)
}
