package main

import (
	"aoc-2023-go/utils"
	"fmt"
)

const (
	DAY             = "20"
	EXPECTED_RESULT = 32000000
)

type Module struct {
	Name     string
	Type     string
	Pulses   []int
	State    int
	Children []*Module
}

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func solve(input []string) int {

	return 0
}

func parse(input []string) {
	for _, line := range input {

	}
}
