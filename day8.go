package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	Register    string
	RegisterVal int
	Change      string
	ChangeVal   int
	ChangeComp  string
	ChangeOp    string
	CompVal     int
}

func main() {
	file, err := ioutil.ReadFile("day8_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1And2(string(file))
}

func part1And2(input string) {
	var maxValue = 0
	var highestGeneralValue = 0
	instructions := createInstructionMap(input)
	lines := strings.Split(input, "\n")

	for _, val := range lines {
		instruction := createInstruction(val)
		if compareInstruction(instruction, instructions) {
			runInstruction(instruction, instructions)
			for _, regVal := range instructions {
				if regVal > highestGeneralValue {
					highestGeneralValue = regVal
				}
			}
		}
	}

	// find max value
	for _, regVal := range instructions {
		if regVal > maxValue {
			maxValue = regVal
		}
	}

	fmt.Println("Largest value is:", maxValue)
	fmt.Println("Highest value held in any register during the process is:", highestGeneralValue)
}

func createInstruction(line string) Instruction {
	var instruction Instruction

	s := strings.Fields(line)
	intChangeVal, _ := strconv.Atoi(s[2])
	intCompVal, _ := strconv.Atoi(s[6])

	instruction = Instruction{
		Register:    s[0],
		RegisterVal: 0,
		Change:      s[1],
		ChangeVal:   intChangeVal,
		ChangeComp:  s[4],
		ChangeOp:    s[5],
		CompVal:     intCompVal,
	}

	return instruction
}

func compareInstruction(inst Instruction, instructions map[string]int) bool {
	comp := instructions[inst.ChangeComp]
	compVal := inst.CompVal

	switch inst.ChangeOp {
	case "==":
		return comp == compVal
	case "!=":
		return comp != compVal
	case ">":
		return comp > compVal
	case ">=":
		return comp >= compVal
	case "<":
		return comp < compVal
	case "<=":
		return comp <= compVal
	}

	return false
}

func runInstruction(inst Instruction, instructions map[string]int) {
	if inst.Change == "inc" {
		instructions[inst.Register] += inst.ChangeVal
	} else {
		instructions[inst.Register] -= inst.ChangeVal
	}
}

func createInstructionMap(input string) map[string]int {
	instructions := make(map[string]int)
	lines := strings.Split(input, "\n")
	for _, val := range lines {
		instructionLine := strings.Split(val, " ")
		if _, ok := instructions[instructionLine[0]]; !ok {
			instructions[instructionLine[0]] = 0
		}
	}

	return instructions
}
