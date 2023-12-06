package main

import (
	"aoc-2023-go/utils"
	"fmt"
)

const (
	DAY = "01"
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2015/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := partA(input[0])
	fmt.Printf("Result A %d\n", result)
	result = partB(input[0])
	fmt.Printf("Result B %d\n", result)
}

func partA(input string) int {
	count := 0
	for _, rune := range input {
		if rune == '(' {
			count++
		} else if rune == ')' {
			count--
		}
	}
	return count
}

func partB(input string) int {
	count := 0
	for i, rune := range input {
		if rune == '(' {
			count++
		} else if rune == ')' {
			count--
		}
		if count == -1 {
			return i + 1
		}
		i++
	}
	return count
}
