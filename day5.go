package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day5_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
}

func part1() {
}

func part2() {
}
