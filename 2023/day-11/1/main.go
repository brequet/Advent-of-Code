package main

import (
	"aoc-2023-go/utils"
	"fmt"
)

const (
	DAY             = "11"
	EXPECTED_RESULT = 374
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

type Galaxy struct {
	row int
	col int
}

func solve(input []string) int {
	grid := createGridUniverse(input)
	printMat(grid)

	galaxies := []Galaxy{}
	for y, row := range grid {
		for x, b := range row {
			if b == '#' {
				galaxies = append(galaxies, Galaxy{row: y, col: x})
			}
		}
	}

	sumShortest := 0
	for i, g1 := range galaxies {
		// closest := findClosest(galaxies, galaxy)
		for j, g2 := range galaxies {
			sumShortest += distance(g1, g2)
			fmt.Println("between", i+1, j+1, distance(g1, g2))
		}
		// fmt.Println("Closest to", galaxy, "is", closest)
	}

	return sumShortest / 2
}

func findClosest(galaxies []Galaxy, galaxy Galaxy) int {
	minDistance := -1
	for _, g := range galaxies {
		distance := distance(galaxy, g)
		if distance > 0 {
			if minDistance < 0 || distance < minDistance {
				minDistance = distance
			}
		}
	}
	return minDistance
}

func distance(g1, g2 Galaxy) int {
	return abs(g2.row-g1.row) + abs(g2.col-g1.col)
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func createGridUniverse(input []string) [][]byte {
	// init
	grid := [][]byte{}
	for i := range input {
		row := []byte(input[i])
		grid = append(grid, row)
	}

	// expand rows
	emptyRowIndexes := []int{}
	for i := range grid {
		if isEmptyRow(grid, i) {
			emptyRowIndexes = append(emptyRowIndexes, i)
		}
	}
	for _, i := range utils.ReverseList(emptyRowIndexes) {
		emptyRow := grid[i]
		if len(grid) == i {
			grid = append(grid, emptyRow)
		} else {
			tmpGrid := append(grid[:i], emptyRow)
			grid = append(tmpGrid, grid[i:]...)
		}
	}

	// expand cols
	emptyColIndexes := []int{}
	for j := range grid[0] {
		if isEmptyCol(grid, j) {
			emptyColIndexes = append(emptyColIndexes, j)
		}
	}
	newGrid := [][]byte{}
	for _, row := range grid {
		newRow := []byte{}
		for j, b := range row {
			newRow = append(newRow, b)
			if utils.Contains(emptyColIndexes, j) {
				newRow = append(newRow, b)
			}
		}
		newGrid = append(newGrid, newRow)
	}

	return newGrid
}

func isEmptyRow(mat [][]byte, rowIndex int) bool {
	for _, b := range mat[rowIndex] {
		if b != '.' {
			return false
		}
	}
	return true
}

func isEmptyCol(mat [][]byte, colIndex int) bool {
	for _, row := range mat {
		if row[colIndex] != '.' {
			return false
		}
	}
	return true
}

func printMat(mat [][]byte) {
	for _, row := range mat {
		for _, b := range row {
			fmt.Printf("%c", b)
		}
		fmt.Println()
	}
}
