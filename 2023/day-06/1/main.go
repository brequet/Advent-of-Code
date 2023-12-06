package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"regexp"
)

func main() {
	input := utils.GetInputAsSlice("./2023/day-06/input")

	numsRe := regexp.MustCompile(`\d+`)
	times := utils.StrSliceToIntSlice(numsRe.FindAllString(input[0], -1))
	distances := utils.StrSliceToIntSlice(numsRe.FindAllString(input[1], -1))

	sum := 1
	for i := 0; i < len(times); i++ {
		sum *= getNumberOfPossibleWinPerRace(times[i], distances[i])
	}
	fmt.Println("Result:", sum, " / 1624896")
}

func getNumberOfPossibleWinPerRace(time, distance int) int {
	winCount := 0
	if time < 1 {
		return 0
	}

	for t := 1; t < time; t++ {
		if t*(time-t) > distance {
			winCount++
		}
	}
	return winCount
}
