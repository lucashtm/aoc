package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  count_nice := 0
  for {
    line, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    line = strings.TrimSpace(line)
    if len(line) == 0 {
      break
    }

    if is_nice(line) {
      count_nice++
    }
  }
  fmt.Printf("Nice strings: %d\n", count_nice)
}

// func is_vowel(char rune) bool {
  // return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'
// }

// func has_enough_vowels(line string) bool {
  // var count_vowels uint32 = 0

  // for _, c := range line {
    // if is_vowel(c) {
      // count_vowels++
    // }
  // }

  // return count_vowels >= 3
// }

// func is_forbidden(sub string) bool {
  // return sub == "ab" || sub == "cd" || sub == "pq" || sub == "xy"
// }

// func no_forbidden_substrings(line string) bool {
  // for i := range line {
    // if i > 0 && is_forbidden(line[i-1:i+1]) {
      // return false
    // }
  // }
  // return true
// }

// func has_double_char(line string) bool {
  // for i, c := range line {
    // if i > 0 && c == []rune(line)[i-1] {
      // return true
    // }
  // }
  // return false
// }

func has_same_pair_twice(line string) bool {
  return false
}

func has_repeating_with_single_letter_between(line string) bool {
  if len(line) < 3 {
    return false
  }

  start := 0
  end := 2
  for end < len(line) {
    if line[start] == line[end] {
      return true
    }
    start++
    end++
  }
  return false
}


func is_nice(line string) bool {
  return has_same_pair_twice(line) &&
  has_repeating_with_single_letter_between(line)
}
