package main

import (
	"aoc-2023-go/utils"
	"fmt"
)

const (
	DAY             = "01"
	EXPECTED_RESULT = 32583852
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2015/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := partA(input[0])
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}
