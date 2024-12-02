package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
  "strconv"
  "strings"
	"fmt"
	"os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
  text = strings.TrimSpace(text)
  number := 0
  for {
    number_as_str := strconv.Itoa(number)
    var sb strings.Builder
    sb.WriteString(text)
    sb.WriteString(number_as_str)
    hash := md5.Sum([]byte(sb.String()))
    result := hex.EncodeToString(hash[:])
    if strings.HasPrefix(result, "00000") {
      fmt.Printf("Suffix: %d\n", number)
      return
    }
    number++
  }
}
