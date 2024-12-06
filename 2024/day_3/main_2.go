package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func compute(memo string) int {
  pattern := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

  result := 0
  for _, submatch := range pattern.FindAllStringSubmatch(memo, -1) {
    op1, err := strconv.Atoi(submatch[1])
    if err != nil {
      log.Fatalf("Could not convert %s to integer\n", submatch[1])
    }
    op2, err := strconv.Atoi(submatch[2])
    if err != nil {
      log.Fatalf("Could not convert %s to integer\n", submatch[2])
    }

    // fmt.Println(op1, "*", op2, "=", op1*op2)
    result += op1 * op2
  }

  return result
}

func remove_donts(memo string) string {
  dont_pattern := "don't()"
  do_pattern := "do()"
  index := strings.Index(memo, dont_pattern)
  if index < 0 {
    return memo
  }

  var sb strings.Builder
  sb.WriteString(memo[:index])
  start := index + len(dont_pattern)
  do := true
  dos := 0
  donts := 0
  for start < len(memo) {
    if do {
      index := strings.Index(memo[start:], do_pattern)
      if index >= 0 {
        dos += 1
        start += index + len(do_pattern)
      } else {
        break
      }
      do = false
    } else {
      index := strings.Index(memo[start:], dont_pattern)
      if index >= 0 {
        donts += 1
        fmt.Println(memo[start:start+index])
        fmt.Println()
        sb.WriteString(memo[start:start+index])
        start += index + len(dont_pattern)
      } else {
        sb.WriteString(memo[start:])
        break
      }
      do = true
    }
  }
  fmt.Println("donts: ", donts)
  fmt.Println("dos: ", dos)

  return sb.String()
}

func main() {
  data, err := io.ReadAll(os.Stdin)
  if err != nil {
    log.Fatalf("Could not read from stdin. err: %s\n", err)
  }

  fmt.Println("Result is: ", compute(remove_donts(string(data))))
}
