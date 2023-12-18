package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	DAY             = "18"
	EXPECTED_RESULT = 32583852
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func solve(input []string) int {
	rowCount, colCount := computeGridSize(input)
	fmt.Println("Grid size (expect 7 x 10)", rowCount, colCount)
	return 0
}

func computeGridSize(instructions []string) (row, col int) {
	maxR, maxL, maxU, maxD := 0, 0, 0, 0

	for _, line := range instructions {
		fields := strings.Fields(line)
		dir := fields[0]
		steps, _ := strconv.Atoi(fields[1])
		switch dir {
		case "R":
			maxR += steps
		case "L":
			maxL += steps
		case "U":
			maxU += steps
		case "D":
			maxD += steps
		}
	}

	if maxR > maxL {
		col = maxR + 1
	} else {
		col = maxL + 1
	}
	if maxU > maxD {
		row = maxU + 1
	} else {
		row = maxD + 1
	}

	return row, col
}
