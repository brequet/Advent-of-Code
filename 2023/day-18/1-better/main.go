package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	DAY             = "18"
	EXPECTED_RESULT = 62
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	directions := map[string][]int{
		"U": {0, -1},
		"D": {0, 1},
		"L": {-1, 0},
		"R": {1, 0},
	}

	digPlan := [][]int{{0, 0}}
	boundaryPoints := 0

	for _, line := range input {
		parts := strings.Split(line, " ")
		direction := parts[0]
		distance, _ := strconv.Atoi(parts[1])

		dr, dc := directions[direction][0], directions[direction][1]

		row, col := digPlan[len(digPlan)-1][0], digPlan[len(digPlan)-1][1]

		boundaryPoints += distance
		digPlan = append(digPlan, []int{row + dr*distance, col + dc*distance})
	}

	area := 0
	for i := 0; i < len(digPlan); i++ {
		if i == 0 {
			area += digPlan[0][0] * (digPlan[len(digPlan)-1][1] - digPlan[(1)%len(digPlan)][1])
		} else {
			area += digPlan[i][0] * (digPlan[i-1][1] - digPlan[(i+1)%len(digPlan)][1])
		}
	}

	area = int(math.Abs(float64(area))) / 2

	interiorPoints := area - boundaryPoints/2 + 1

	// fmt.Println(interiorPoints + boundaryPoints)
	result := interiorPoints + boundaryPoints
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}
