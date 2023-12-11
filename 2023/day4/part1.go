package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var winningRE = regexp.MustCompile(`:(.*?)\|`)
var userNumRE = regexp.MustCompile(`\|(.*)`)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func findNumbersInLine(line string) [][2]int {
  numberIndices:= make([][2]int, 0)

  for idx := 0; idx < len(line); idx++ {
    _, errParse := strconv.Atoi(string(line[idx]))

    if errParse == nil {
      for startIndex, subStrChar := range line[idx:] {
        endIndex := 0
        _, errSubStrParse := strconv.Atoi(string(subStrChar))

        if errSubStrParse != nil {
          endIndex = startIndex + idx
          indices := [2]int{idx, endIndex}
          numberIndices = append(numberIndices, indices)
          idx = endIndex

          break
        }

        _, errParseFullSubstring := strconv.Atoi(line[idx:])

        if errParseFullSubstring == nil {
          indices := [2]int{idx, len(line)}
          numberIndices = append(numberIndices, indices)
          return numberIndices
        }
      }
    }
  }

  return numberIndices
}

func parseNumberIndices(line string, indices [][2]int) []int {
  result := []int{}

  for _, i := range indices {
    start := i[0]
    end := i[1]

    num, err := strconv.Atoi(line[start:end])

    if err == nil {
      result = append(result, num)
    }
  }

  return result
}

func getWinningUserNumbers(winningNumbers []int, userNumbers []int) []int {
  result := []int{}

  for _, userNum := range userNumbers {
    for _, winningNum := range winningNumbers {
      if userNum == winningNum {
        result = append(result, userNum)
      }
    }
  }

  return result
}

func main() {
  f, err := os.Open("./cards-day2.txt")
  check(err)

  total := 0
  scanner := bufio.NewScanner(f)

  for {
    scanner.Scan()
    line := scanner.Text()

    if len(line) == 0 {
      break
    }

    winningMatchSubstring := winningRE.FindStringSubmatch(line)
    userNumbersMatchSubstring := userNumRE.FindStringSubmatch(line)
    winningNumbers := parseNumberIndices(
      winningMatchSubstring[1],
      findNumbersInLine(winningMatchSubstring[1]))

    userNumbers := parseNumberIndices(
      userNumbersMatchSubstring[1],
      findNumbersInLine(userNumbersMatchSubstring[1]))

    userWinningNumbers := getWinningUserNumbers(winningNumbers, userNumbers)

    if len(userWinningNumbers) > 0 {
      points := 1

      for i := 1; i < len(userWinningNumbers); i++ {
        if points == 1 {
          points = 2
          continue
        }

        points = points * 2
      }

      total += points
    }
  }

  println(total)
}
