package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"regexp"
	"strings"
)

const (
	DAY             = "19"
	EXPECTED_RESULT = 167409079868000
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

type Interval struct {
	from, to int
}

func (i Interval) String() string {
	return fmt.Sprintf("[%d;%d]", i.from, i.to)
}

type Rule struct {
	keyCondition   string
	valueCondition int
	comparator     string
	output         string
}

type Workflows map[string][]Rule

type Part map[string]int

func (part Part) rating() int {
	return part["x"] + part["m"] + part["a"] + part["s"]
}

func solve(lines []string) (res int) {
	workflows := getData(lines)
	fmt.Println("workflows", workflows)

	// mergedIntervals := map[string][]Interval{}
	for i, rules := range workflows {
		intervals := getIntervalsForWorkflow(rules)

		// for _, key := range []string{"x", "m", "a", "s"} {
		// 	mergedIntervals[key] = mergeInterval(mergedIntervals[key], intervals[key])
		// }
		// res += getIntervalsForWorkflow(rules)
		fmt.Printf("Workflow [%3s]: %v ----> %v \n", i, rules, intervals)
	}

	return res
}

// func mergeInterval(intervalList []Interval, interval2 Interval) []Interval {

// }

func getIntervalsForWorkflow(rules []Rule) map[string]Interval {
	intervals := map[string]Interval{}
	for _, rule := range rules {
		if inter, ok := getIntervalForRule(rule); ok {
			intervals[rule.keyCondition] = inter
		}
	}

	for _, key := range []string{"x", "m", "a", "s"} {
		if _, ok := intervals[key]; !ok {
			intervals[key] = Interval{1, 4000}
		}
	}

	return intervals
}

func getIntervalForRule(rule Rule) (Interval, bool) {
	if rule.keyCondition == "none" {
		return Interval{}, false
	}

	if rule.comparator == ">" {
		return Interval{rule.valueCondition + 1, 4000}, true
	} else {
		return Interval{1, rule.valueCondition - 1}, true
	}
}

func getData(lines []string) Workflows {
	workflows := Workflows{}
	for _, line := range lines {
		if line == "" {
			break
		}
		key := strings.Split(line, "{")[0]
		workflows[key] = getRulesFromLine(line)
	}
	return workflows
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
