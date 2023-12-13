package main

import (
	"aoc-2023-go/utils"
	"fmt"
)

const (
	DAY             = "13"
	EXPECTED_RESULT = 405
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func solve(input []string) int {
	patterns := getPatterns(input)

	res := 0
	for _, pattern := range patterns {
		res += countPatternPoint(pattern)
	}

	return res
}

func getPatterns(input []string) [][]string {
	res := [][]string{}

	pattern := []string{}
	for _, line := range input {
		if len(line) > 0 {
			pattern = append(pattern, line)
		} else {
			res = append(res, pattern)
			pattern = []string{}
		}
	}
	if len(pattern) > 0 {
		res = append(res, pattern)
	}

	return res
}

func countPatternPoint(pattern []string) int {
	return 100*findHorizontalReflection(pattern) + findHorizontalReflection(getTranspose(pattern))
}

func findHorizontalReflection(pattern []string) int {
	// TODO
	return 0
}

func getTranspose(pattern []string) []string {
	return pattern // TODO
}
