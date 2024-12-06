package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const word string = "XMAS"

var (
  RIGHT = []int{0, 1}
  LOWER_RIGHT = []int{1, 1}
  DOWN = []int{1, 0}
  LOWER_LEFT = []int{1, -1}
  LEFT = []int{0, -1}
  UPPER_LEFT = []int{-1, -1}
  UP = []int{-1, 0}
  UPPER_RIGHT = []int{-1, 1}
)

var DIRECTIONS = [][]int{
  RIGHT,
  LOWER_RIGHT,
  DOWN,
  LOWER_LEFT,
  LEFT,
  UPPER_LEFT,
  UP,
  UPPER_RIGHT,
}

func out_of_bounds(i int, j int, width int, height int) bool {
  return i < 0 || i >= height || j < 0 || j >= width
}

func puzzle_1(lines []string, i int, j int) int {
  width := len(lines[0])
  height := len(lines)
  counter := 0
  for _, direction := range DIRECTIONS {
    current := 1
    found := true
    x := i + direction[0]
    y := j + direction[1]
    for current < len(word) {
      if out_of_bounds(x, y, width, height) {
        found = false
        break
      }

      if lines[x][y] != word[current] {
        found = false
        break
      }
      current++
      x += direction[0]
      y += direction[1]
    }

    if found {
      counter++
    }
  }
  return counter
}

func puzzle_2(lines []string, i int, j int) int {
  width := len(lines[0])
  height := len(lines)

  upper_left := []int{i - 1, j - 1}
  lower_right := []int{i + 1, j + 1}
  upper_right := []int{i - 1, j + 1}
  lower_left := []int{i + 1, j - 1}

  directions := [][]int{
    upper_left,
    lower_right,
    upper_right,
    lower_left,
  }

  for _, direction := range directions {
    if out_of_bounds(direction[0], direction[1], width, height) {
      return 0
    }
  }

  diag_1 := string([]byte{lines[upper_left[0]][upper_left[1]], lines[i][j], lines[lower_right[0]][lower_right[1]]})
  diag_2 := string([]byte{lines[upper_right[0]][upper_right[1]], lines[i][j], lines[lower_left[0]][lower_left[1]]})

  if (diag_1 == "MAS" || diag_1 == "SAM") && (diag_2 == "MAS" || diag_2 == "SAM") {
    return 1
  }
  return 0
}

func main() {
  input, _ := io.ReadAll(os.Stdin)
  text := strings.TrimSpace(string(input))
  lines := strings.Split(text, "\n")
  width := len(lines[0])
  height := len(lines)
  counter := 0

  for i := 0; i < height; i++ {
    for j := 0; j < width; j++ {
      // if lines[i][j] == word[0] {
        // counter += puzzle_1(lines, i, j)
      // }
      if lines[i][j] == 'A' {
        counter += puzzle_2(lines, i, j)
      }
    }
  }

  fmt.Println(counter)
}
