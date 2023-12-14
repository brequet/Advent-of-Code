package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strings"
)

const (
	DAY             = "13"
	EXPECTED_RESULT = 43614
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
	hIndex, hCount := findHorizontalReflection(pattern)
	vIndex, vCount := findHorizontalReflection(getTranspose(pattern))
	if hCount >= vCount {
		return 100 * hIndex
	} else {
		return vIndex
	}
}

func findHorizontalReflection(pattern []string) (rowIndex int, count int) {
	for i := 0; i < len(pattern)-1; i++ {
		line, nextLine := pattern[i], pattern[i+1]
		if line == nextLine {
			firstHalf, secondHalf := utils.ReverseList(pattern[:i+1]), pattern[i+1:]
			innerCount := 0
			for j := 0; j < utils.Min(len(firstHalf), len(secondHalf)); j++ {
				if firstHalf[j] != secondHalf[j] {
					break
				}
				innerCount++
			}
			if innerCount == utils.Min(len(firstHalf), len(secondHalf)) && innerCount > count {
				count = innerCount
				rowIndex = i
			}
			if innerCount == utils.Min(len(firstHalf), len(secondHalf)) {
				break
			}
		}
	}

	return (rowIndex + 1), count
}

func getTranspose(pattern []string) []string {
	// Find the length of the longest string in the slice
	maxLen := 0
	for _, str := range pattern {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}

	// Create a new slice of strings with the same length as the longest string
	transposed := make([]string, maxLen)

	// Iterate over each string in the input slice
	for _, str := range pattern {
		// Convert the string to a slice of runes
		runes := []rune(str)
		// Iterate over each rune in the string
		for i, r := range runes {
			// Append the rune to the corresponding string in the transposed slice
			transposed[i] += string(r)
		}
	}

	return transposed
}

func patternToString(pattern []string) (res string) {
	return "\n" + strings.Join(pattern, "\n") + "\n"
}
