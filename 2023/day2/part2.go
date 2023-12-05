package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
  if (e != nil) {
    panic(e)
  }
}

func main() {
  re := regexp.MustCompile(`^Game.+:`)
  f, fileErr := os.Open("./game-doc-day2.txt")
  check(fileErr)

  scanner := bufio.NewScanner(f)
  sumOfPowers := 0

  for {
    scanner.Scan()
    line := scanner.Text()

    if len(line) == 0 {
      break
    }

    lineSplitByGameHand := strings.Split(re.ReplaceAllString(line, ""), ";")
    fewestPossibleCubeCountByColor := map[string]int{}
    power := 1

    for _, hand := range lineSplitByGameHand {
      splitByCube := strings.Split(hand, ",")

      for _, val := range splitByCube {
        countAndCubeColor := strings.Split(strings.Trim(val, " "), " ")
        cubeCount, countParseError := strconv.Atoi(countAndCubeColor[0])
        _, ok := fewestPossibleCubeCountByColor[countAndCubeColor[1]]

        if countParseError != nil {
          break
        }

        if !ok || ok && cubeCount >= fewestPossibleCubeCountByColor[countAndCubeColor[1]] {
          fewestPossibleCubeCountByColor[countAndCubeColor[1]] = cubeCount
        }
      }
    }

    for _, count := range fewestPossibleCubeCountByColor {
      power *= count
    }

    sumOfPowers += power
  }

  println("POWER!!! ", sumOfPowers)
}
