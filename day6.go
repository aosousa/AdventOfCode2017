package main

import (
	"fmt"
	"reflect"
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

func part1(nums []int) ([]int, []int) {
	var compNums = make([][]int, 0)
	var cycles = 0

	for {
		var max = 0
		var maxIndex = 0
		var prevNums = make([]int, len(nums))

		// Find highest value in slice
		for index, val := range nums {
			if val > max {
				max = val
				maxIndex = index
			}
		}

		nums[maxIndex] = 0

		// Redistribute blocks
		for i := 0; i < max; i++ {
			maxIndex++
			if maxIndex > len(nums)-1 {
				maxIndex = 0
			}
			nums[maxIndex]++
		}

		// Check for a possible repeat
		for _, value := range compNums {
			if reflect.DeepEqual(nums, value) {
				fmt.Println("Redistribution Cycles:", cycles+1)
				return nums, value
			}
		}

		copy(prevNums, nums)
		compNums = append(compNums, prevNums)

		cycles++
	}
}

func part2(nums []int) {
	nums, compNums := part1(nums)

	var cycles = 0

	for {
		var max = 0
		var maxIndex = 0

		// Find highest value in slice
		for index, val := range nums {
			if val > max {
				max = val
				maxIndex = index
			}
		}

		nums[maxIndex] = 0

		// Redistribute values
		for i := 0; i < max; i++ {
			maxIndex++
			if maxIndex > len(nums)-1 {
				maxIndex = 0
			}
			nums[maxIndex]++
		}

		// Check for a possible repeat
		if reflect.DeepEqual(nums, compNums) {
			fmt.Println("Redistribution Cycles:", cycles+1)
			return
		}

		cycles++
	}
}
