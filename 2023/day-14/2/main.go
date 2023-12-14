package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strings"
)

const (
	DAY             = "14"
	EXPECTED_RESULT = 64
)

func main() {
	input := utils.GetInput(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

type Sol struct {
	currentMat   [][]byte
	nextRotation [][]byte
}

var cache = map[string]Sol{}

func solve(input [][]byte) int {
	loads := []int{}

	tilted := utils.CopyMat(input)
	// for 1000 is same answer as for 1e10
	for i := 0; i < 1000; i++ {
		tilted = tiltCycle(tilted)
		load := countLoadNorth(tilted)

		loads = append(loads, load)

	}

	return countLoadNorth(tilted)
}

func generateKeyFromMat(mat [][]byte) (key string) {
	for _, row := range mat {
		key += string(row)
	}
	return key
}

func tiltCycle(mat [][]byte) [][]byte {
	tilted := utils.CopyMat(mat)
	for i := 0; i < 4; i++ {
		tilted = tiltNorth(tilted)
		tilted = utils.MatRotate(tilted)
	}
	return tilted
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
