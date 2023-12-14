package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strings"
)

const (
	DAY             = "13"
	EXPECTED_RESULT = 36771
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
	for i, pattern := range patterns {
		v := countPatternPoint(pattern)
		res += v
		fmt.Println(i, v)
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

func countDiff(line1, line2 string) (res int) {
	for i := range line1 {
		if line1[i] != line2[i] {
			res++
		}
	}
	return res
}

func findHorizontalReflection(pattern []string) (rowIndex int, count int) {
	// TODO
	for i := 0; i < len(pattern)-1; i++ {
		line, nextLine := pattern[i], pattern[i+1]

		cDiff := countDiff(line, nextLine)
		if cDiff <= 1 {
			// fmt.Println("Matching line at index", i, line)
			firstHalf, secondHalf := utils.ReverseList(pattern[:i+1]), pattern[i+1:]
			innerCount := 0
			cDiffCount := 0
			for j := 0; j < utils.Min(len(firstHalf), len(secondHalf)); j++ {
				if firstHalf[j] != secondHalf[j] {
					if cDiffCount > 0 && countDiff(firstHalf[j], secondHalf[j]) == 1 {
						fmt.Printf("Difference in lines ! '%s' != '%s\n", line, nextLine)
						cDiffCount++
					} else {
						break
					}
				}
				innerCount++
			}
			if innerCount > count {
				count = innerCount
				rowIndex = i
			}
			if innerCount == utils.Min(len(firstHalf), len(secondHalf)) {
				break
			}
		}
	}

	fmt.Println("For pattern", patternToString(pattern), "score", rowIndex, "count", count)
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

	// fmt.Println("Pattern", pattern)
	// fmt.Println("Pattern transposed", transposed)

	return transposed
}

func patternToString(pattern []string) (res string) {
	return "\n" + strings.Join(pattern, "\n") + "\n"
}
