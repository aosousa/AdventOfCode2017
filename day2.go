package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// part1(scanner)
	part2(scanner)
}

func part1(input *bufio.Scanner) {
	var min int
	var max int
	checksum := 0

	for input.Scan() {
		max = math.MinInt32
		min = math.MaxInt32
		currentRow := strings.Fields(input.Text())
		for _, col := range currentRow {
			val, _ := strconv.Atoi(col)
			if val > max {
				max = val
			}

			if val <= min {
				min = val
			}
		}
		checksum += (max - min)
	}

	fmt.Println("Checksum:", checksum)
}

func part2(input *bufio.Scanner) {
	sum := 0

	for input.Scan() {
		currentRow := strings.Fields(input.Text())
		for index, col := range currentRow {
			for _, col2 := range currentRow[index+1:] {
				val1, _ := strconv.Atoi(col)
				val2, _ := strconv.Atoi(col2)

				if val1%val2 == 0 {
					sum += val1 / val2
				}

				if val2%val1 == 0 {
					sum += val2 / val1
				}
			}
		}
	}

	fmt.Println("Sum:", sum)
}
