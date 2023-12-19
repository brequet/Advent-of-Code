package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strings"
)

const (
	DAY             = "19"
	EXPECTED_RESULT = 19114
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

type Rule struct {
	keyCondition   string
	valueCondition int
	output         string
}

type Workflows map[string][]Rule

type Part map[string]int

func solve(lines []string) (res int) {
	workflows, parts := getData(lines)
	fmt.Println("workflows", workflows)
	fmt.Println("parts", parts)
	return res
}

func getData(lines []string) (Workflows, []Part) {
	workflows := Workflows{}
	isSecondPart := false
	for _, line := range lines {
		if line == "" {
			isSecondPart = true
			continue
		}
		if !isSecondPart {
			key := strings.Split(line, "{")[0]
			workflows[key] = getRulesFromLine(line)
		} else {

		}
	}
	return nil, nil
}

func getRulesFromLine(line string) []Rule {
	goodPart := strings.TrimRight(strings.Split(line, "{")[1], "}")
	fmt.Println("good part", goodPart) // todo remove
	for _, e := strings.Split(goodPart, ",") {
		
	}
	return []Rule{}
}
