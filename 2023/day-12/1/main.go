package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strings"
)

const (
	DAY             = "12"
	EXPECTED_RESULT = 1
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

func countArrangementInLine(line string) int {
	s := strings.Split(line, " ")
	springs := []byte(s[0])
	countiguousGroups := utils.Map(strings.Split(s[1], ","), utils.StrToInt)

	res := countArrangement(springs, countiguousGroups)
	fmt.Println("Row:", string(springs), "Configuration:", printTest(countiguousGroups), "=>", res)
	return res
}

func printTest(ns []int) string {
	return "[" + strings.Join(utils.Map(ns, func(n int) string { return fmt.Sprintf("'%d'", n) }), ", ") + "]"
}

func countArrangement(springs []byte, contiguousGroups []int) (arrangementCount int) {
	// fmt.Println("count in ", string(springs), contiguousGroups, "      (size", utils.SumAll(contiguousGroups), "versus", len(springs), ")")
	if len(contiguousGroups) == 0 || utils.SumAll(contiguousGroups) > len(springs) {
		// fmt.Println("\t->0")
		return 0
	} else if len(contiguousGroups) == 1 {
		v := findPossibleMatchCount(springs, contiguousGroups[0])
		// fmt.Println("\t-> ", v)
		return v
	}

	h, t := contiguousGroups[0], contiguousGroups[1:]
	arrangements := getPossibleArrangements(springs, h)
	// fmt.Println("  Possible arrangement:", arrangements)

	for arrangementSprings := range arrangements {
		arrangementCount += countArrangement([]byte(arrangementSprings), t)
	}

	return arrangementCount
}

func isUnusedHashes(springs []byte, group int) bool {
	return strings.Contains(string(springs[group:]), "#")
}

func findPossibleMatchCount(springs []byte, groupSize int) (count int) {
	for i := 0; i < len(springs)-groupSize+1; i++ {
		subSprings := springs[i : i+groupSize]
		if isFullSprings(subSprings) && !strings.Contains(string(springs[i+groupSize:]), "#") && !strings.Contains(string(springs[:i]), "#") {
			// fmt.Println("\t\tMATCHING", string(subSprings))
			count++
		}
	}

	return count
}

func getPossibleArrangements(springs []byte, groupSize int) map[string]struct{} {
	newSprings := map[string]struct{}{}
	encounteredHashes := 0
	for i := 0; i < len(springs)-groupSize+1; i++ {
		subSprings := springs[i : i+groupSize]
		if isFullSprings(subSprings) {
			arrangement := replaceChars(springs, i, i+groupSize)
			if i+groupSize+1 < len(springs) && springs[i+groupSize] != '#' {
				// fmt.Println("Possible", string(arrangement[i+groupSize+1:]), "(from)", i, "sub", string(subSprings), "springs[i+groupSize]", string(springs[i+groupSize]))
				newSprings[string(arrangement[i+groupSize+1:])] = struct{}{}
			}
		}
		if springs[i] == '#' {
			encounteredHashes++
		}
		if countBackwardHashes(springs, i+groupSize) >= groupSize || encounteredHashes > 0 {
			break
		}
	}
	return newSprings
}

func countBackwardHashes(springs []byte, index int) (count int) {
	return strings.Count(string(springs[:index]), "#")
}

func isFullSprings(subSprings []byte) bool {
	if len(subSprings) == 0 {
		return false
	}
	for _, b := range subSprings {
		if b == '.' {
			return false
		}
	}
	return true
}

func replaceChars(str []byte, from, to int) []byte {
	newStr := []byte(strings.Clone(string(str)))
	for i := from; i < to; i++ {
		newStr[i] = '#'
	}
	return newStr
}
