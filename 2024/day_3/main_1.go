package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
  data, err := io.ReadAll(os.Stdin)
  if err != nil {
    log.Fatalf("Could not read from stdin. err: %s\n", err)
  }
  pattern := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

  result := 0
  for _, submatch := range pattern.FindAllStringSubmatch(string(data), -1) {
    op1, err := strconv.Atoi(submatch[1])
    if err != nil {
      log.Fatalf("Could not convert %s to integer\n", submatch[1])
    }
    op2, err := strconv.Atoi(submatch[2])
    if err != nil {
      log.Fatalf("Could not convert %s to integer\n", submatch[2])
    }

    fmt.Println(op1, "*", op2, "=", op1*op2)
    result += op1 * op2
  }

  fmt.Println("Result is: ", result)
}
