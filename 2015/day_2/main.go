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
  total_wrap := 0
  total_ribbon := 0
  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    line = strings.TrimSpace(line)
    if len(line) == 0 {
      break
    }
    var dimensions [3]int
    var smallest_dim int
    var second_smallest_dim int
    for i, dim := range strings.Split(line, "x") {
      dimensions[i], err = strconv.Atoi(dim)
      if err != nil {
        log.Fatalf("Cant convert %s to integer\n", dim)
      }
    }
    area_1 := dimensions[0]*dimensions[1]
    area_2 := dimensions[1]*dimensions[2]
    area_3 := dimensions[2]*dimensions[0]
    total_wrap += area_1*2 + area_2*2 + area_3*2
    total_wrap += min(area_1, area_2, area_3)

    total_ribbon += 2*(smallest_dim+second_smallest_dim)
    total_ribbon += dimensions[0] * dimensions[1] * dimensions[2]
    fmt.Print(line)
    fmt.Printf(" -> %dx%d\n", smallest_dim, second_smallest_dim)
  }
  fmt.Printf("Wrap: %d\n", total_wrap)
  fmt.Printf("Ribbon: %d\n", total_ribbon)
}
