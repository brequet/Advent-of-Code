package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strings"
)

const (
	DAY             = "15"
	EXPECTED_RESULT = 1320
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input[0])
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func hash(entry string) (res int) {
	for _, char := range entry {
		res += int(char)
		res *= 17
		res = res % 256
	}
	return res
}

func solve(input string) int {
	sum := 0
	for _, seq := range strings.Split(input, ",") {
		sum += hash(seq)
	}
	return sum
}
