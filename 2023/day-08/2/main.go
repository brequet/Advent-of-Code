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
	nodeRe := regexp.MustCompile(`[1-9A-Z]{3}`)

	startingNodeToStepCount := map[string]int{}
	keyToNode := map[string]Node{}
	for _, v := range aMap {
		nodeData := nodeRe.FindAllString(v, -1)
		key := nodeData[0]
		node := Node{
			left:  nodeData[1],
			right: nodeData[2],
		}
		keyToNode[key] = node

		if key[2] == 'A' {
			startingNodeToStepCount[key] = 0
		}
	}
	fmt.Println("Starting", startingNodeToStepCount)

	steps := []int{}
	for startingNode := range startingNodeToStepCount {
		// startingNodeToStepCount[startingNode] = getNbrStepUntilNextZ(keyToNode, instructions, startingNode, 0)
		steps = append(steps, getNbrStepUntilNextZ(keyToNode, instructions, startingNode, 0))
	}

	return utils.LCM(steps...)
}

func getNbrStepUntilNextZ(nodes map[string]Node, pattern string, startNode string, startStep int) int {
	currentNode := getNextNode(nodes, pattern, startNode, 0)
	step := startStep + 1
	for currentNode[2] != 'Z' {
		currentNode = getNextNode(nodes, pattern, currentNode, step)
		step++
	}
	return step - startStep
}

func getNextNode(nodes map[string]Node, pattern string, currentNode string, step int) string {
	action := pattern[step%len(pattern)]
	if action == 'L' {
		return nodes[currentNode].left
	} else {
		return nodes[currentNode].right
	}
}

func getLcm(numbers []int) int {
	return 0
}

func getIsAllOnZ(nodeMap map[string]string) bool {
	res := true
	for _, node := range nodeMap {
		if node[2] != 'Z' {
			res = false
			break
		}
	}
	return res
}
