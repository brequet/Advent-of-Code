package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"log"
	"regexp"
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
	comparator     string
	output         string
}

type Workflows map[string][]Rule

type Part map[string]int

func solve(lines []string) (res int) {
	workflows, parts := getData(lines)

	for _, part := range parts {
		if isPartAccepted(part, workflows, "in") {
			res += part["x"] + part["m"] + part["a"] + part["s"]
		}
	}

	return res
}

func isPartAccepted(part Part, workflows Workflows, currentWorkflowKey string) bool {
	if currentWorkflowKey == "A" {
		return true
	} else if currentWorkflowKey == "R" {
		return false
	}

	currentWorkflowRules := workflows[currentWorkflowKey]
	for _, rule := range currentWorkflowRules {
		if isPartPassingRule(part, rule) {
			return isPartAccepted(part, workflows, rule.output)
		}
	}
	log.Fatal("Bof..")
	return false
}

func isPartPassingRule(part Part, rule Rule) bool {
	if rule.keyCondition == "none" {
		return true
	}

	if rule.comparator == ">" {
		return part[rule.keyCondition] > rule.valueCondition
	} else {
		return part[rule.keyCondition] < rule.valueCondition
	}
}

func getData(lines []string) (Workflows, []Part) {
	workflows := Workflows{}
	parts := []Part{}
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
			parts = append(parts, getPartsFromLine(line))
		}
	}
	return workflows, parts
}

func getRulesFromLine(line string) (rules []Rule) {
	goodPart := strings.TrimRight(strings.Split(line, "{")[1], "}")
	ruleRe := regexp.MustCompile(`(x|m|a|s)(<|>)(\d+):(\w+)`)
	for _, e := range strings.Split(goodPart, ",") {
		matches := ruleRe.FindAllStringSubmatch(e, -1)
		if len(matches) > 0 {
			rule := Rule{
				keyCondition:   matches[0][1],
				comparator:     matches[0][2],
				valueCondition: utils.StrToInt(matches[0][3]),
				output:         matches[0][4],
			}
			rules = append(rules, rule)
		} else {
			rule := Rule{
				keyCondition:   "none",
				comparator:     "",
				valueCondition: 0,
				output:         e,
			}
			rules = append(rules, rule)
		}
	}
	return rules
}

func getPartsFromLine(line string) Part {
	partRe := regexp.MustCompile(`x=(\d+),m=(\d+),a=(\d+),s=(\d+)`)
	matches := partRe.FindAllStringSubmatch(line, -1)
	part := Part{}
	part["x"] = utils.StrToInt(matches[0][1])
	part["m"] = utils.StrToInt(matches[0][2])
	part["a"] = utils.StrToInt(matches[0][3])
	part["s"] = utils.StrToInt(matches[0][4])
	return part
}
