package main

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"strconv"
	"strings"
)

func get_int_lines(reader *bufio.Reader, sep string) [][]int {
  lines := make([][]int, 0)
  for {
    data, err := reader.ReadString('\n')
    if err != nil {
      break
    }
    data = strings.TrimSpace(data)
    if data == "" {
      break
    }

    pages_str := strings.Split(data, sep)
    var ints []int
    for _, page_str := range pages_str {
      numeric, err := strconv.Atoi(page_str)
      if err != nil {
        log.Fatalf("Could not convert %s to integer\n", page_str)
      }
      ints = append(ints, numeric)
    }
    lines = append(lines, ints)
  }
  return lines
}

type rule_map_t = map[int][2]map[int]bool
const BEFORE int = 0
const AFTER int = 1

func generate_rule_map(rules [][]int) rule_map_t {
  rule_map := rule_map_t{}
  for _, rule := range rules {
    left := rule[0]
    right := rule[1]

    if _, ok := rule_map[left]; !ok {
      rule_map[left] = [2]map[int]bool{{}, {}}
    }
    rule_map[left][AFTER][right] = true

    if _, ok := rule_map[right]; !ok {
      rule_map[right] = [2]map[int]bool{{}, {}}
    }
    rule_map[right][BEFORE][left] = true
  }
  return rule_map
}

func print_rule_map(rule_map rule_map_t) {
  for key := range maps.Keys(rule_map) {
    fmt.Printf("%d:\n", key)
    fmt.Print("  BEFORE: [")
    for b_key := range maps.Keys(rule_map[key][BEFORE]) {
      if rule_map[key][BEFORE][b_key] {
        fmt.Print(b_key, " ")
      }
    }
    fmt.Println("]")
    fmt.Print("  AFTER: [")
    for a_key := range maps.Keys(rule_map[key][AFTER]) {
      if rule_map[key][AFTER][a_key] {
        fmt.Print(a_key, " ")
      }
    }
    fmt.Println("]")
  }
}

func update_is_valid(rule_map rule_map_t, update []int) bool {
  for index, page := range update {
    for b := index - 1; b >= 0; b-- {
      if rule_map[page][AFTER][update[b]] {
        return false
      }
    }

    for a := index + 1; a < len(update); a++ {
      if rule_map[page][BEFORE][update[a]] {
        return false
      }
    }
  }
  return true
}

func fix_update_1(rule_map rule_map_t, update []int) {
  for index, page := range update {
    for b := index - 1; b >= 0; b-- {
      if rule_map[page][AFTER][update[b]] {
        update[b], update[index] = update[index], update[b]
      }
    }

    for a := index + 1; a < len(update); a++ {
      if rule_map[page][BEFORE][update[a]] {
        update[a], update[index] = update[index], update[a]
      }
    }
  }
}

func should_swap(rule_map rule_map_t, first int, last int) bool {
  return rule_map[first][BEFORE][last] || rule_map[last][AFTER][first]
}

func fix_update(rule_map rule_map_t, update []int) {
  outer_start := 0
  outer_end := len(update) - 1
  for outer_start < outer_end {
    start := outer_start
    end := outer_end
    for start < end {
      if should_swap(rule_map, update[start], update[end]) {
        update[start], update[end] = update[end], update[start]
      }
      start++
    }

    start = outer_start
    end = outer_end
    for start < end {
      if should_swap(rule_map, update[start], update[end]) {
        update[start], update[end] = update[end], update[start]
      }
      end--
    }

    outer_start++
    outer_end--
  }
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  rules := get_int_lines(reader, "|")
  rule_map := generate_rule_map(rules)
  updates := get_int_lines(reader, ",")
  result_1 := 0
  result_2 := 0
  for _, update := range updates {
    if update_is_valid(rule_map, update) {
      fmt.Println("Valid:  ", update)
      middle := update[len(update)/2]
      result_1 += middle
    } else {
      fmt.Print("Invalid:", update)
      fix_update(rule_map, update)
      middle := update[len(update)/2]
      result_2 += middle
      fmt.Println(" ->", update)
    }
  }

  fmt.Println("Puzzle 1:", result_1)
  fmt.Println("Puzzle 2:", result_2)
}
