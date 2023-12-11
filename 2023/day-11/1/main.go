package main

import (
	"aoc-2023-go/utils"
	"fmt"
)

const (
	DAY             = "11"
	EXPECTED_RESULT = 374
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func solve(input []string) int {
	grid := createGridUniverse(input)
	printMat(grid)
	return 0
}

func createGridUniverse(input []string) [][]byte {
	grid := make([][]byte, len(input))
	for i := range input {
		grid[i] = []byte(input[i])
	}
	// TODO expand if necessary
	return grid
}

func printMat(mat [][]byte) {
	for _, row := range mat {
		for _, b := range row {
			fmt.Printf("%c", b)
		}
		fmt.Println()
	}
}
