package main

import "fmt"

func main() {
	var input = 312051

	part1(input)
	num := 5
	len := 2
	for i, draw := range part2(num) {
		fmt.Printf("%*d ", len, draw)
		if i%num == num-1 {
			fmt.Println("")
		}
	}
}

func part1(input int) {
	var start = 1
	var startCol = 0
	var inputRow = 0
	var add = 8
	var steps = 0

	for ok := true; ok; ok = (start < input) {
		start += add
		add += 8
		startCol++
		inputRow += 2
	}

	var inputCol = inputRow

	if start > input {
		inputCol = (start - input)
	}

	steps = (inputRow - startCol) + (inputCol - startCol)

	fmt.Println("Steps:", steps)
}

func part2(n int) []int {
	left, top, right, bottom := 0, 0, n-1, n-1
	width := 1
	sz := n * n
	s := make([]int, sz)
	i := 0
	for left < right {
		width += 2
		right++ // move to the right

		// work right, along top
		for c := left; c <= right; c++ {
			s[top*n+c] = i
			i++
		}
		top++

		// work up left side
		for r := bottom; r >= top; r-- {
			s[r*n+left] = i
			i++
		}
		left++

		// work down right side
		for r := top; r <= bottom; r++ {
			s[r*n+right] = i
			i++
		}
		right--
		if top == bottom {
			break
		}
		// work left, along bottom
		for c := right; c >= left; c-- {
			s[bottom*n+c] = i
			i++
		}
		bottom--
	}
	// center (last) element
	s[top*n+left] = i

	return s
}

func calculateSurroundingSum(grid []int, x int, y int) {
	var sum = 0
}
