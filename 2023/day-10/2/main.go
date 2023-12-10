package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"log"
)

const (
	DAY             = "10"
	EXPECTED_RESULT = 6815
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

	return part2(input, mat)
}

func part2(input []string, mat [][]Node) (count int) {
	for i, row := range mat {
		for j, node := range row {
			if !node.isVisited && isWithin(input, i, j) {
				count++
			}
		}
	}
	return count
}

func isWithin(input []string, row, col int) bool {

}

func printMat[T any](mat [][]T) {
	for _, row := range mat {
		for _, node := range row {
			fmt.Print(node)
		}
		fmt.Println()
	}
}

func countNonNullNodes(mat [][]Node) (res int) {
	for _, row := range mat {
		for _, node := range row {
			if node.isVisited {
				res++
			}
		}
	}
	return res
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
		// fmt.Printf("Iteration [%d]: (%d, %d) %c\n", i, rowCurrent, colCurrent, input[rowCurrent][colCurrent])
		rowNext, colNext := findNextCoordinate(input, mat, rowCurrent, colCurrent)
		nextValue := input[rowNext][colNext]
		// fmt.Printf("\tFound: (%d, %d) %c\n", rowNext, colNext, nextValue)
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
	if tile == 'S' {
		nextRow, nextCol = findStartingPointNeighbour(input, row, col)
	}
	if tile == '|' {
		nextCol = col
		if isPrevious(mat, row+1, col, row, col) {
			fmt.Println("↑")
			nextRow = row - 1
		} else {
			fmt.Println("↓")
			nextRow = row + 1
		}
	} else if tile == '-' {
		nextRow = row
		if isPrevious(mat, row, col+1, row, col) {
			fmt.Println("←")
			nextCol = col - 1
		} else {
			fmt.Println("→")
			nextCol = col + 1
		}
	} else if tile == 'L' {
		if isPrevious(mat, row, col+1, row, col) {
			fmt.Println("↑")
			nextRow = row - 1
			nextCol = col
		} else {
			fmt.Println("→")
			nextRow = row
			nextCol = col + 1
		}
	} else if tile == 'J' {
		if isPrevious(mat, row, col-1, row, col) {
			fmt.Println("↑")
			nextRow = row - 1
			nextCol = col
		} else {
			fmt.Println("←")
			nextRow = row
			nextCol = col - 1
		}
	} else if tile == '7' {
		if isPrevious(mat, row, col-1, row, col) {
			fmt.Println("↓")
			nextRow = row + 1
			nextCol = col
		} else {
			fmt.Println("←")
			nextRow = row
			nextCol = col - 1
		}
	} else if tile == 'F' {
		if isPrevious(mat, row, col+1, row, col) {
			fmt.Println("↓")
			nextRow = row + 1
			nextCol = col
		} else {
			fmt.Println("→")
			nextRow = row
			nextCol = col + 1
		}
	}
	return nextRow, nextCol
}

func findStartingPointNeighbour(input []string, row, col int) (int, int) {
	fmt.Println("Is starting point")

	if row > 0 && utils.Contains([]byte{'|', '7', 'F'}, input[row-1][col]) {
		// Is above
		return row - 1, col
	} else if row < len(input)-1 && utils.Contains([]byte{'|', 'L', 'J'}, input[row+1][col]) {
		// Is below
		return row + 1, col
	} else if col > 0 && utils.Contains([]byte{'-', '7', 'J'}, input[row][col-1]) {
		// Is left
		return row, col - 1
	} else if row < len(input[0])-1 && utils.Contains([]byte{'-', 'L', 'F'}, input[row][col+1]) {
		// Is right
		return row, col + 1
	} else {
		log.Fatalln("Can't find starting point !")
		return 0, 0
	}
}

func isPrevious(mat [][]Node, row, col, cRow, cCol int) bool {
	node := mat[row][col]
	return node.row == cRow && node.col == cCol
}
