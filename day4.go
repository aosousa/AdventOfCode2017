package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("day4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// part1(scanner)
	part2(scanner)
}

func part1(input *bufio.Scanner) {
	var validCount = 0

	for input.Scan() {
		var invalidCount = 0
		currentRow := strings.Fields(input.Text())
		for index, val := range currentRow {
			for _, val2 := range currentRow[index+1:] {
				if val == val2 {
					invalidCount++
				}
			}
		}

		if invalidCount == 0 {
			validCount++
		}
	}

	fmt.Println("Valid passphrases:", validCount)
}

func part2(input *bufio.Scanner) {
	var validCount = 0

	for input.Scan() {
		var anagram = false
		currentRow := strings.Fields(input.Text())

		for index, val := range currentRow {
			sortedString := sortString(val)
			for _, val2 := range currentRow[index+1:] {
				sortedString2 := sortString(val2)
				if sortedString == sortedString2 {
					anagram = true
				}
			}
		}
		if anagram == false {
			validCount++
		}
	}

	fmt.Println("Valid passphrases:", validCount)
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
