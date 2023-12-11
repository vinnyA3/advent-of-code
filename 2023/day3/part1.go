package main

import (
	"os"
	"strconv"
	"strings"
)

var LINE_BUFFER_COUNT = 5

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func checkIsValidSpecialCharacter(char string) bool {
  _, errParse := strconv.Atoi(char)

  if errParse != nil {
    if char != "." {
      return true
    } else {
      return false
    }
  } else {
    return false
  }
}

func byteToString(val byte) string {
  return string(val)
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
          endIndex = idx + startIndex
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

type ValidateNumberParams struct {
  lines []string
  readOffsetIndex int
}

func validateNumbersAdjacentToSpecialChars(params ValidateNumberParams) int {
  total := 0
  lines := params.lines
  lineOffset := params.readOffsetIndex
  validNumberIndices:= make([][2]int, 0)

  for lineIndex := lineOffset; lineIndex < len(lines); lineIndex++ {
    line := lines[lineIndex]
    numberIndices := findNumbersInLine(line)
    shouldCheckPrevLine := false
    shouldCheckNextLine := false

    if lineIndex - 1 >= 0 {
      shouldCheckPrevLine = true
    } 

    if lineIndex + 1 < len(lines) {
      shouldCheckNextLine = true
    }

    for _, indicesPair := range numberIndices {
      startOfNum := indicesPair[0]
      endOfNum := indicesPair[1] - 1
      startCheckIdx := startOfNum
      endCheckIdx := endOfNum

      if startOfNum - 1 >= 0 {
        startCheckIdx = startOfNum - 1
      }

      if endOfNum + 1 < len(line) {
        endCheckIdx = endOfNum + 1
      }

      for i := startCheckIdx; i <= endCheckIdx; i++ {
        if shouldCheckPrevLine && checkIsValidSpecialCharacter(byteToString(lines[lineIndex - 1][i])) {
          validNumberIndices = append(validNumberIndices, indicesPair)
          break;
        }

        if checkIsValidSpecialCharacter(byteToString(lines[lineIndex][i])) {
          validNumberIndices = append(validNumberIndices, indicesPair)
          break;
        }

        if shouldCheckNextLine && checkIsValidSpecialCharacter(byteToString(lines[lineIndex + 1][i])) {
          validNumberIndices = append(validNumberIndices, indicesPair)
          break;
        }
      }
    }

    for _, validIndices := range validNumberIndices {
      startIdx := validIndices[0]
      endIdx := validIndices[1]
      number, errParse := strconv.Atoi(line[startIdx:endIdx])

      if errParse == nil {
        total += number
      }
    }

    validNumberIndices = nil
  }

  return total
}

func main() {
  // Reading into memory, for input file is pretty small.  If were dealing
  // with a crazy large file, would implementing some buffering scheme
  file, err := os.ReadFile("./schematic-day3.txt")
  check(err)

  total := 0
  lines := strings.Split(string(file), "\n")

  total = validateNumbersAdjacentToSpecialChars(ValidateNumberParams{
    lines: lines[:len(lines) - 1], // omit last newline of slice
    readOffsetIndex: 0,
  })
 
  println("total is: ", total)
}
