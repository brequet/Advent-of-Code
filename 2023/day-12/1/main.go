package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strings"
)

const (
	DAY             = "12"
	EXPECTED_RESULT = 21
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func solve(input []string) (arrangementCount int) {
	for _, line := range input {
		arrangementCount += countArrangementInLine(line)
	}

	return arrangementCount
}

func countArrangementInLine(line string) (arrangementCount int) {
	s := strings.Split(line, " ")
	springs := []byte(s[0])
	countiguousGroups := utils.Map(strings.Split(s[1], ","), utils.StrToInt)

	return countArrangement(springs, countiguousGroups)
}

func countArrangement(springs []byte, contiguousGroups []int) (arrangementCount int) {
	fmt.Println(string(springs), contiguousGroups)

	return 0
}
