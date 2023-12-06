package main

import (
	"aoc-2023-go/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func inputToSlice(filePath string) (res []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func main() {
	input := inputToSlice("./day-06/input")

	time := utils.StrToInt(strings.Replace(strings.Split(input[0], ":")[1], " ", "", -1))
	distance := utils.StrToInt(strings.Replace(strings.Split(input[1], ":")[1], " ", "", -1))
	fmt.Println(time, distance)

	fmt.Println("Result:", getNumberOfPossibleWinPerRace(time, distance), " / 71503")
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
	fmt.Println("For", time, distance, "win", winCount)
	return winCount
}
