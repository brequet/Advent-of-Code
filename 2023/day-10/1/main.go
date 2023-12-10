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
	value          byte
	row, col       int
	previous, next *Node
}

func (n Node) String() string {
	var previous, next byte
	if n.previous != nil {
		previous = n.previous.value
	}
	if n.next != nil {
		next = n.next.value
	}
	return fmt.Sprintf("Node{value: %c, row: %d, col: %d, previous: %c, next: %c}", n.value, n.row, n.col, previous, next)
}

func solve(input []string) int {
	startNode := findStartingNode(input)
	fmt.Println(startNode)

	findLoop(input, startNode)

	return 0
}

func findStartingNode(input []string) Node {
	for i, row := range input {
		for j, e := range row {
			if e == 'S' {
				return Node{
					value: 'S',
					row:   i,
					col:   j,
				}
			}
		}
	}
	log.Fatalln("Could not found starting node !")
	return Node{}
}

func findLoop(input []string, startNode Node) {
	nextNode1 := findNextNeighbours(input, startNode)
	startNode.next = &nextNode1
	var currentNodeAddr *Node = &nextNode1
	for currentNodeAddr.value != 'S' {
		fmt.Printf("Iteration on: %s\n", *currentNodeAddr)
		nextNode := findNextNeighbours(input, *currentNodeAddr)
		currentNodeAddr.next = &nextNode

		fmt.Printf("\t--> next %s\n", nextNode)
		fmt.Printf("\t\t\t--> startNode %s\n", startNode)
		if nextNode.next.value == 'S' {
			fmt.Printf("\t\t setting previous to %s\n", currentNodeAddr)
			startNode.previous = currentNodeAddr
		}

		currentNodeAddr = &nextNode
	}
	fmt.Println("start", startNode)
}

func findNextNeighbours(input []string, node Node) (next Node) {
	tile := node.value
	if tile == '|' {
		next.col = node.col
		if node.row+1 == node.previous.row {
			fmt.Println("↑")
			next.row = node.row - 1
		} else {
			fmt.Println("↓")
			next.row = node.row + 1
		}
	} else if tile == '-' {
		next.row = node.row
		if node.col+1 == node.previous.col {
			fmt.Println("←")
			next.col = node.col - 1
		} else {
			fmt.Println("→")
			next.col = node.col + 1
		}
	} else if tile == 'L' {
		if node.row-1 != node.previous.row {
			fmt.Println("↑")
			next.row = node.row - 1
			next.col = node.col
		} else {
			fmt.Println("→")
			next.row = node.row
			next.col = node.col + 1
		}
	} else if tile == 'J' {
		if node.row-1 != node.previous.row {
			fmt.Println("↑")
			next.row = node.row - 1
			next.col = node.col
		} else {
			fmt.Println("←")
			next.row = node.row
			next.col = node.col - 1
		}
	} else if tile == '7' {
		if node.row+1 != node.previous.row {
			fmt.Println("↓")
			next.row = node.row + 1
			next.col = node.col
		} else {
			fmt.Println("←")
			next.row = node.row
			next.col = node.col - 1
		}
	} else if tile == 'F' {
		if node.row+1 != node.previous.col {
			fmt.Println("↓")
			next.row = node.row + 1
			next.col = node.col
		} else {
			fmt.Println("→")
			next.row = node.row
			next.col = node.col + 1
		}
	} else {
		fmt.Println("Is starting point, default")
		rows := []int{node.row - 1, node.row, node.row + 1}
		cols := []int{node.col - 1, node.col, node.col + 1}
		for _, row := range rows {
			if row < 0 || row >= len(input[0]) {
				continue
			}
			for _, col := range cols {
				if (row == node.row && col == node.col) || col < 0 || col >= len(input) {
					continue
				}
				cur := input[row][col]
				if cur != '.' && cur != node.value {
					fmt.Printf("isNeighbour: %d %d %c\n", row, col, cur)
					next.row = row
					next.col = col
				}
			}
		}
	}
	next.value = input[next.row][next.col]
	next.previous = &node
	return next
}
