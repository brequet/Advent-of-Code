package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"regexp"
)

const (
	DAY             = "08"
	EXPECTED_RESULT = 6
)

type Node struct {
	left  string
	right string
}

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := countSteps(input[0], input[2:])
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func countSteps(instructions string, aMap []string) int {
	nodeRe := regexp.MustCompile(`[A-Z]{3}`)

	keyToNode := map[string]Node{}
	for _, v := range aMap {
		nodeData := nodeRe.FindAllString(v, -1)
		key := nodeData[0]
		node := Node{
			left:  nodeData[1],
			right: nodeData[2],
		}
		keyToNode[key] = node
	}

	stepsCount, isOnZ := 0, false
	currentNode := "AAA"
	instructionsCount := len(instructions)
	for !isOnZ {
		action := instructions[stepsCount%instructionsCount]
		if action == 'L' {
			currentNode = keyToNode[currentNode].left
		} else {
			currentNode = keyToNode[currentNode].right
		}
		stepsCount++

		if currentNode == "ZZZ" {
			isOnZ = true
		}
	}

	return stepsCount
}
