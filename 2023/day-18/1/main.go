package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	DAY             = "18"
	EXPECTED_RESULT = 32583852
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func solve(input []string) int {
	rowCount, colCount, startCoord := computeGridSize(input)

	grid := traceTrenches(input, rowCount, colCount, startCoord)
	fRow, fCol := 0, 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				fRow, fCol = y, x
				break
			}
		}
		break
	}
	floodFill(grid, fRow+1, fCol+1)

	fmt.Println("\\Filled\n", utils.MatToStr(grid, utils.ByteToStr))

	sum := 0
	for _, row := range grid {
		for _, char := range row {
			if char == '#' {
				sum++
			}
		}
	}

	return sum
}

func floodFill(grid [][]byte, x, y int) {
	// Base case: if the starting point is outside the grid or not a '.', return
	fmt.Println(x, y, grid[0])
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != '.' {
		return
	}

	// Mark the current point as '#'
	grid[x][y] = '#'

	// Recursively call floodFill on all adjacent points
	floodFill(grid, x-1, y)
	floodFill(grid, x+1, y)
	floodFill(grid, x, y-1)
	floodFill(grid, x, y+1)
}

type Coord struct {
	row int
	col int
}

func traceTrenches(instr []string, row, col int, startCoord Coord) [][]byte {
	grid := make([][]byte, row)
	for i := range grid {
		grid[i] = make([]byte, col)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	// fmt.Println("Init\n", utils.MatToStr(grid, utils.ByteToStr))

	currentCoord := startCoord
	grid[currentCoord.row][currentCoord.col] = '#'
	for _, line := range instr {
		fields := strings.Fields(line)
		dir := fields[0]
		steps, _ := strconv.Atoi(fields[1])
		for i := 0; i < steps; i++ {
			switch dir {
			case "R":
				currentCoord.col++
			case "L":
				currentCoord.col--
			case "D":
				currentCoord.row++
			case "U":
				currentCoord.row--
			}
			grid[currentCoord.row][currentCoord.col] = '#'
		}
	}
	// fmt.Println("\nEnd\n", utils.MatToStr(grid, utils.ByteToStr))

	return grid
}

func computeGridSize(instructions []string) (row, col int, startingCoord Coord) {
	maxCol, minCol, minRow, maxRow := 0, 0, 0, 0

	curRow, curCol := 0, 0

	for _, line := range instructions {
		fields := strings.Fields(line)
		dir := fields[0]
		steps, _ := strconv.Atoi(fields[1])
		switch dir {
		case "R":
			curCol += steps
		case "L":
			curCol -= steps
		case "D":
			curRow += steps
		case "U":
			curRow -= steps
		}

		if curRow > maxRow {
			maxRow = curRow
		}
		if curRow < minRow {
			minRow = curRow
		}

		if curCol > maxCol {
			maxCol = curCol
		}
		if curCol < minCol {
			minCol = curCol
		}

	}

	fmt.Println(maxRow, minRow, maxCol, minCol)

	startCoord := Coord{0, 0}
	if minRow < 0 {
		startCoord.row = -minRow
	}

	if minCol < 0 {
		startCoord.col = -minCol
	}

	return maxRow - minRow + 1, maxCol - minCol + 1, startCoord
}
