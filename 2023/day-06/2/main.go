package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.GetInputAsSlice("./2023/day-06/input")

	time := utils.StrToInt(strings.Replace(strings.Split(input[0], ":")[1], " ", "", -1))
	distance := utils.StrToInt(strings.Replace(strings.Split(input[1], ":")[1], " ", "", -1))

	fmt.Println("Result:", getNumberOfPossibleWinPerRace(time, distance), " / 32583852")
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
