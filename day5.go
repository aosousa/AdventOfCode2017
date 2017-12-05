package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var numbersOne, numbersTwo []int

	file, err := os.Open("day5_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbersOne = append(numbersOne, num)
		numbersTwo = append(numbersTwo, num)
	}

	part1(numbersOne)
	part2(numbersTwo)
}

func part1(nums []int) {
	var steps = 0
	var index = 0
	var prevInstructionValue = 0

	for {
		index += prevInstructionValue
		if index >= len(nums) {
			break
		} else {
			steps++
			nums[index]++
			prevInstructionValue = nums[index] - 1
		}
	}

	fmt.Println("Steps:", steps)
}

func part2(nums []int) {
	var steps = 0
	var index = 0
	var prevInstructionValue = 0

	for {
		index += prevInstructionValue
		if index >= len(nums) {
			break
		} else {
			steps++
			if nums[index] >= 3 {
				prevInstructionValue = nums[index]
				nums[index]--
			} else {
				nums[index]++
				prevInstructionValue = nums[index] - 1
			}
		}
	}

	fmt.Println("Steps:", steps)
}
