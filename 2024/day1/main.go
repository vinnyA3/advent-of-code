package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func diff(x int, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}

func get_sorted_lists(reader *bufio.Scanner) ([]int, []int) {
	l1 := []int{}
	l2 := []int{}

	for {
		reader.Scan()
		line := reader.Text()

		if len(line) == 0 {
			break
		}

		numbers := strings.Fields(line)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])

		l1 = append(l1, num1)
		l2 = append(l2, num2)
	}

	sort.Sort(sort.IntSlice(l1))
	sort.Sort(sort.IntSlice(l2))

	return l1, l2
}

func get_total_distance(l1 []int, l2 []int) int {
	total := 0

	for i := 0; i < len(l1); i++ {
		distance := diff(l1[i], l2[i])
		total += distance
	}

	return total
}

func get_total_similarity(l1 []int, l2 []int) int {
	total := 0

	for i := 0; i < len(l1); i++ {
		n1 := l1[i]
		n1_occurrence_in_l2 := 0

		for j := 0; j < len(l2); j++ {
			n2 := l2[j]

			if n1 == n2 {
				n1_occurrence_in_l2 += 1
			}
		}

		if n1_occurrence_in_l2 > 0 {
			total += n1_occurrence_in_l2 * n1
			n1_occurrence_in_l2 = 0
		}
	}

	return total
}

func main() {
	f, err := os.Open("input-list.txt")
	check(err)

	reader := bufio.NewScanner(f)
	l1, l2 := get_sorted_lists(reader)
	// total_distance:= get_total_distance(l1, l2) // Part 1 Solution
	total_similarity := get_total_similarity(l1, l2) // Part 2 Solution
	// fmt.Println(total_distance) // Part 1 solution
	fmt.Println(total_similarity)
}
