package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strings"
)

const (
	DAY             = "14"
	EXPECTED_RESULT = 136
)

func main() {
	input := utils.GetInput(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func solve(input [][]byte) int {
	fmt.Printf("Puzzle entry:\n\n%s\n", utils.MatToStr(input, utils.ByteToStr))

	tiltedNorth := tiltNorth(input)
	fmt.Printf("Tilted north:\n\n%s\n", utils.MatToStr(tiltedNorth, utils.ByteToStr))

	return countLoadNorth(tiltedNorth)
}

func tiltNorth(mat [][]byte) (tiltedMat [][]byte) {
	tiltedMat = utils.CopyMat(mat)
	isProcessOver := false
	for !isProcessOver {
		operationCount := 0
		for i := range tiltedMat {
			if i == 0 {
				continue
			}
			for j := range tiltedMat[i] {
				if tiltedMat[i][j] == 'O' && tiltedMat[i-1][j] == '.' {
					tiltedMat[i-1][j] = 'O'
					tiltedMat[i][j] = '.'
					operationCount++
				}
			}
		}
		if operationCount == 0 {
			isProcessOver = true
		}
	}
	return tiltedMat
}

func countLoadNorth(mat [][]byte) (res int) {
	matLen := len(mat)
	for i, row := range mat {
		res += strings.Count(string(row), "O") * (matLen - i)
	}
	return res
}
