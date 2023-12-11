package main

import (
	"aoc-2023-go/utils"
	"fmt"
)

const (
	DAY             = "11"
	EXPECTED_RESULT = 840988812853
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
	grid, emptyRowIndexes, emptyColIndexes := createGridUniverse(input)
	printMat(grid)

	galaxies := []Galaxy{}
	for y, row := range grid {
		for x, b := range row {
			if b == '#' {
				galaxies = append(galaxies, Galaxy{row: y, col: x})
			}
		}
	}

	fmt.Println("number of galaxies to process", len(galaxies), emptyRowIndexes, emptyColIndexes)

	sumShortest := 0
	for _, g1 := range galaxies {
		for _, g2 := range galaxies {
			d := distance(g1, g2, emptyRowIndexes, emptyColIndexes)
			sumShortest += d
		}
	}

	return sumShortest / 2
}

func distance(g1, g2 Galaxy, emptyRowIndexes, emptyColIndexes []int) int {
	if g1.col == g2.col && g1.row == g2.row {
		return 0
	}

	expandedColCrossedCount, expandedRowCrossedCount := 0, 0
	for _, emptyRowIndex := range emptyRowIndexes {
		if (g2.row > emptyRowIndex && emptyRowIndex > g1.row) || (g1.row > emptyRowIndex && emptyRowIndex > g2.row) {
			expandedRowCrossedCount++
		}
	}
	for _, emptyColIndex := range emptyColIndexes {
		if (g2.col > emptyColIndex && emptyColIndex > g1.col) || (g1.col > emptyColIndex && emptyColIndex > g2.col) {
			expandedColCrossedCount++
		}
	}
	return abs(g2.row-g1.row) + abs(g2.col-g1.col) + (expandedColCrossedCount+expandedRowCrossedCount)*(1000000-1)
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func createGridUniverse(input []string) ([][]byte, []int, []int) {
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

	// expand cols
	emptyColIndexes := []int{}
	for j := range grid[0] {
		if isEmptyCol(grid, j) {
			emptyColIndexes = append(emptyColIndexes, j)
		}
	}

	return grid, emptyRowIndexes, emptyColIndexes
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
	for i, row := range mat {
		for j, b := range row {
			if b == '#' {
				fmt.Printf("(%d, %d)", i, j)
			} else {
				fmt.Print("[    ]")
			}
		}
		fmt.Println()
	}
}
