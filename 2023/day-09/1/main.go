package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strings"
)

const (
	DAY             = "09"
	EXPECTED_RESULT = 1934898178
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input)
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func solve(input []string) (res int) {

	for _, line := range input {
		nums := utils.StrSliceToIntSlice(strings.Split(line, " "))
		res += extrapolate(nums)
	}

	return res
}

func extrapolate(nums []int) int {
	extrapolationList := [][]int{nums}

	// up to bottom
	for !isExtrapolationListRead(extrapolationList) {
		lastList := utils.GetLastElement(extrapolationList)
		newList := []int{}
		for i := 0; i < len(lastList)-1; i++ {
			a, b := lastList[i], lastList[i+1]
			newList = append(newList, b-a)
		}
		extrapolationList = append(extrapolationList, newList)
	}

	// bottom up
	for i := len(extrapolationList) - 1; i >= 0; i-- {
		newNum := 0
		if i != len(extrapolationList)-1 {
			upperList := extrapolationList[i+1]
			currentList := extrapolationList[i]

			leftVal, bottomVal := utils.GetLastElement(currentList), utils.GetLastElement(upperList)

			newNum = leftVal + bottomVal
		}

		extrapolationList[i] = append(extrapolationList[i], newNum)
	}

	nextValue := utils.GetLastElement(extrapolationList[0])

	return nextValue
}

func isExtrapolationListRead(eList [][]int) bool {
	lastList := utils.GetLastElement(eList)
	for _, num := range lastList {
		if num != 0 {
			return false
		}
	}
	return true
}
