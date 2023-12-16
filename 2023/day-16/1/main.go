package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"log"
)

const (
	DAY             = "16"
	EXPECTED_RESULT = 46
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
	visited := make([][][]string, len(input))
	for i := range input {
		visited[i] = make([][]string, len(input[0]))
		for j := range input[0] {
			visited[i][j] = []string{}
		}
	}

	// printVisited(visited)
	step(input, visited, Coord{row: 0, col: 0}, "R")
	// printVisited(visited)

	sum := 0
	for _, row := range visited {
		for _, dirs := range row {
			if len(dirs) > 0 {
				sum++
			}
		}
	}
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
