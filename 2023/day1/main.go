package main

import (
  "bufio"
  "os"
  "strconv"
)

var wordNumberMap = map[string]int{
  "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func check(e error) {
  if (e != nil) {
    panic(e)
  }
}

func first_and_last_num(line string) (first_val int, last_val int) {
  var first int
  var last int

  for startCharIdx := 0; startCharIdx < len(line); startCharIdx++ {
    parsedCharAsInt, errParsedCharAsInt := strconv.Atoi(string(line[startCharIdx]))

    if errParsedCharAsInt == nil {
      if first == 0 {
        first = parsedCharAsInt
      }

      last = parsedCharAsInt
    } else {
      for word, assocNumber := range wordNumberMap {
        if startCharIdx + len(word) <= len(line) && line[startCharIdx:startCharIdx + len(word)] == word {
          if first == 0 {
            first = assocNumber
          }

          last = assocNumber
          break
        }
      }
    }
  }

  return first, last
}

func main() {
  f, err := os.Open("./calibration-doc-day1.txt");
  check(err)

  reader := bufio.NewScanner(f)
  total_sum := 0

  for {
    reader.Scan();
    line := reader.Text();

    if len(line) == 0 {
      break
    }

    first, last := first_and_last_num(line)
    firstToString := strconv.FormatInt(int64(first), 10)
    lastToString := strconv.FormatInt(int64(last), 10)
    firstAndLastAsWholeNumber, errParse := strconv.Atoi(string(firstToString + lastToString))

    if (errParse == nil) {
      total_sum += firstAndLastAsWholeNumber
    }
  }

  println("TOTAL:", total_sum)
}
