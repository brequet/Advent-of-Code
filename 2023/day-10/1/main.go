package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"log"
)

const (
	DAY             = "10"
	EXPECTED_RESULT = 4
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

type Node struct {
	row, col  int
	isVisited bool
}

func (n Node) String() string {
	if !n.isVisited {
		return fmt.Sprintf("(., .)")
	} else {
		return fmt.Sprintf("(%d, %d)", n.row, n.col)
	}
}

func solve(input []string) int {
	sRow, sCol := findStartingNode(input)
	fmt.Println("Starting at", sRow, sCol)

	mat := buildAdjacencyMatrix(input, sRow, sCol)
	printMat(mat)

	return 0
}

func printMat[T any](mat [][]T) {
	for _, row := range mat {
		for _, node := range row {
			fmt.Print(node)
		}
		fmt.Println()
	}
}

func findStartingNode(input []string) (int, int) {
	for i, row := range input {
		for j, e := range row {
			if e == 'S' {
				return i, j
			}
		}
	}
	log.Fatalln("Could not found starting node !")
	return 0, 0
}

func buildAdjacencyMatrix(input []string, sRow, sCol int) [][]Node {
	// init matrix
	mat := make([][]Node, len(input))
	for i := range mat {
		mat[i] = make([]Node, len(input[0]))
	}

	isLoopDone := false
	rowCurrent, colCurrent := sRow, sCol
	i := 0
	for !isLoopDone {
		fmt.Printf("Iteration [%d]: (%d, %d) %c\n", i, rowCurrent, colCurrent, input[rowCurrent][colCurrent])
		rowNext, colNext := findNextCoordinate(input, mat, rowCurrent, colCurrent)
		nextValue := input[rowNext][colNext]
		fmt.Printf("\tFound: (%d, %d) %c\n", rowNext, colNext, nextValue)
		mat[rowCurrent][colCurrent] = Node{row: rowNext, col: colNext, isVisited: true}

		rowCurrent, colCurrent = rowNext, colNext
		if nextValue == 'S' {
			isLoopDone = true
			break
		}
		i++
	}

	return mat
}

func findNextCoordinate(input []string, mat [][]Node, row, col int) (nextRow, nextCol int) {
	tile := input[row][col]
	if tile == '|' {
		nextCol = col
		if isVisited(mat, row+1, col) {
			fmt.Println("↑")
			nextRow = row - 1
		} else {
			fmt.Println("↓")
			nextRow = row + 1
		}
	} else if tile == '-' {
		nextRow = row
		if isVisited(mat, row, col+1) {
			fmt.Println("←")
			nextCol = col - 1
		} else {
			fmt.Println("→")
			nextCol = col + 1
		}
	} else if tile == 'L' {
		if isVisited(mat, row, col+1) {
			fmt.Println("↑")
			nextRow = row - 1
			nextCol = col
		} else {
			fmt.Println("→")
			nextRow = row
			nextCol = col + 1
		}
	} else if tile == 'J' {
		if isVisited(mat, row, col-1) {
			fmt.Println("↑")
			nextRow = row - 1
			nextCol = col
		} else {
			fmt.Println("←")
			nextRow = row
			nextCol = col - 1
		}
	} else if tile == '7' {
		if isVisited(mat, row, col-1) {
			fmt.Println("↓")
			nextRow = row + 1
			nextCol = col
		} else {
			fmt.Println("←")
			nextRow = row
			nextCol = col - 1
		}
	} else if tile == 'F' {
		if isVisited(mat, row, col+1) {
			fmt.Println("↓")
			nextRow = row + 1
			nextCol = col
		} else {
			fmt.Println("→")
			nextRow = row
			nextCol = col + 1
		}
	} else {
		fmt.Println("Is starting point, default")
		rows := []int{row - 1, row, row + 1}
		cols := []int{col - 1, col, col + 1}
		for _, cRow := range rows {
			if cRow < 0 || cRow >= len(input[0]) {
				continue
			}
			for _, cCol := range cols {
				if (cRow == row && cCol == col) || cCol < 0 || cCol >= len(input) {
					continue
				}
				cur := input[cRow][cCol]
				if cur != '.' && cur != input[row][col] {
					fmt.Printf("\tfound candidate: %d %d %c\n", cRow, cCol, cur)
					nextRow = cRow
					nextCol = cCol
				}
			}
		}
	}
	return nextRow, nextCol
}

func isVisited(mat [][]Node, row, col int) bool {
	return mat[row][col].isVisited
}
