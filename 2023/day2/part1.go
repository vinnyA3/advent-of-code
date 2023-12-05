package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
  "strconv"
)

var cubeMap = map[string]int{
  "red": 12,
  "green": 13,
  "blue": 14,
}

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
  sumOfValidGameIDs := 0
  gameCount := 1;

  for {
    scanner.Scan()
    line := scanner.Text()
    currentGameIsValid := true

    if len(line) == 0 {
      break
    }

    lineSplitByGameHand := strings.Split(re.ReplaceAllString(line, ""), ";")

    for _, hand := range lineSplitByGameHand {
      splitByCube := strings.Split(hand, ",")

      for _, val := range splitByCube {
        countAndCubeColor := strings.Split(strings.Trim(val, " "), " ")
        cubeCount, countParseError := strconv.Atoi(countAndCubeColor[0])
        countInCubeMap, countInCubeMapExists := cubeMap[countAndCubeColor[1]]

        if countParseError == nil && countInCubeMapExists && cubeCount > countInCubeMap {
          currentGameIsValid = false
          break
        }
      }

      if !currentGameIsValid {
        break
      }
    }


    if currentGameIsValid {
      sumOfValidGameIDs += gameCount
    }

    gameCount++
  }

  println("sum of game valid IDs: ", sumOfValidGameIDs)
}
