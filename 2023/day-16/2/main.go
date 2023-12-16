package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"log"
)

const (
	DAY             = "16"
	EXPECTED_RESULT = 51
)

func main() {
	input := utils.GetInput(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

type Coord struct {
	row int
	col int
}

func printVisited(visited [][][]string) {
	fmt.Println(utils.MatToStr(visited, func(strs []string) string {
		if len(strs) > 0 {
			return "#"
		} else {
			return "."
		}
	}))
}

func solve(input [][]byte) int {
	max := -1
	for i := range input {
		max = utils.Max(max, solveForStartingPosition(input, Coord{row: i, col: 0}, "R"))
		max = utils.Max(max, solveForStartingPosition(input, Coord{row: i, col: len(input[0]) - 1}, "L"))
	}
	for j := range input[0] {
		max = utils.Max(max, solveForStartingPosition(input, Coord{row: 0, col: j}, "D"))
		max = utils.Max(max, solveForStartingPosition(input, Coord{row: len(input) - 1, col: j}, "U"))
	}
	return max
}

func solveForStartingPosition(grid [][]byte, startingPosition Coord, direction string) int {
	visited := make([][][]string, len(grid))
	for i := range grid {
		visited[i] = make([][]string, len(grid[0]))
		for j := range grid[0] {
			visited[i][j] = []string{}
		}
	}

	step(grid, visited, startingPosition, direction)

	sum := 0
	for _, row := range visited {
		for _, dirs := range row {
			if len(dirs) > 0 {
				sum++
			}
		}
	}
	// fmt.Printf("For entry %v %s -> %v\n", startingPosition, direction, sum)
	return sum
}

func getNextCoord(currentCoord Coord, direction string) Coord {
	nextCoord := Coord{row: currentCoord.row, col: currentCoord.col}
	switch direction {
	case "R":
		nextCoord.col++
	case "L":
		nextCoord.col--
	case "U":
		nextCoord.row--
	case "D":
		nextCoord.row++
	}
	return nextCoord
}

func step(grid [][]byte, visited [][][]string, currentCoord Coord, direction string) int {
	if currentCoord.row < 0 || currentCoord.row > len(grid)-1 ||
		currentCoord.col < 0 || currentCoord.col > len(grid[0])-1 ||
		utils.Contains(visited[currentCoord.row][currentCoord.col], direction) {
		return 0
	}
	visited[currentCoord.row][currentCoord.col] = append(visited[currentCoord.row][currentCoord.col], direction)
	// fmt.Println("Step")
	// printVisited(visited)
	currentTile := getTile(grid, currentCoord)

	if currentTile == '.' {
		return step(grid, visited, getNextCoord(currentCoord, direction), direction)
	} else if currentTile == '-' {
		switch direction {
		case "R", "L":
			return step(grid, visited, getNextCoord(currentCoord, direction), direction)
		case "U", "D":
			return step(grid, visited, getNextCoord(currentCoord, "L"), "L") + step(grid, visited, getNextCoord(currentCoord, "R"), "R")
		}
	} else if currentTile == '|' {
		switch direction {
		case "R", "L":
			return step(grid, visited, getNextCoord(currentCoord, "U"), "U") + step(grid, visited, getNextCoord(currentCoord, "D"), "D")
		case "U", "D":
			return step(grid, visited, getNextCoord(currentCoord, direction), direction)
		}
	} else if currentTile == '/' {
		switch direction {
		case "R":
			return step(grid, visited, getNextCoord(currentCoord, "U"), "U")
		case "L":
			return step(grid, visited, getNextCoord(currentCoord, "D"), "D")
		case "U":
			return step(grid, visited, getNextCoord(currentCoord, "R"), "R")
		case "D":
			return step(grid, visited, getNextCoord(currentCoord, "L"), "L")
		}
	} else if currentTile == '\\' {
		switch direction {
		case "R":
			return step(grid, visited, getNextCoord(currentCoord, "D"), "D")
		case "L":
			return step(grid, visited, getNextCoord(currentCoord, "U"), "U")
		case "U":
			return step(grid, visited, getNextCoord(currentCoord, "L"), "L")
		case "D":
			return step(grid, visited, getNextCoord(currentCoord, "R"), "R")
		}
	}

	log.Fatalln("ERROR")
	return 0
	// return step(grid, nextCoord, nextDirection)
}

func getTile(grid [][]byte, coord Coord) byte {
	return grid[coord.row][coord.col]
}
