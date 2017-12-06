package main

import (
	"strconv"
	"strings"
)

func main() {
	var numbersOne, numbersTwo []int

	input := "11	11	13	7	0	15	5	5	4	4	1	1	7	1	15	11"

	currentRow := strings.Fields(input)
	for _, val := range currentRow {
		intVal, _ := strconv.Atoi(val)

		numbersOne = append(numbersOne, intVal)
		numbersTwo = append(numbersTwo, intVal)
	}

	part1(numbersOne)
	part2(numbersTwo)
}

func part1(nums []int) {

}

func part2(nums []int) {

}
