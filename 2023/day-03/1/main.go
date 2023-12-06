package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func getInputArr() []string {
	file, err := os.Open("./2023/day-03/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	inputArr := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputArr = append(inputArr, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputArr
}

type Number struct {
	value string
	x, y  int
}

func main() {
	inputArr := getInputArr()

	nums := []Number{}
	for i, line := range inputArr {
		nums = append(nums, extractFromLine(line, i)...)
	}
	fmt.Println(nums)

	sum := 0
	for _, num := range nums {
		sum += computeNumValue(inputArr, num)
	}
	fmt.Println("result: ", sum)
}

func extractFromLine(line string, y int) (nums []Number) {
	re := regexp.MustCompile("(\\d+)")
	res := re.FindAllStringIndex(line, -1)

	for _, numIndexes := range res {
		x0, x1 := numIndexes[0], numIndexes[1]
		nums = append(nums, Number{
			value: line[x0:x1],
			x:     x0,
			y:     y,
		})
	}

	return nums
}

func computeNumValue(inputArr []string, num Number) int {
	if isPartNumber(inputArr, num) {
		result, err := strconv.Atoi(num.value)
		if err != nil {
			log.Fatal(err)
		}
		return result
	} else {
		return 0
	}
}

// ASCII code for '.': 46
func isPartNumber(inputArr []string, num Number) (res bool) {
	fmt.Println("Computing ", num.value, num.x, num.x+len(num.value))
	for i := num.x; i < num.x+len(num.value); i++ {
		isSymbolClose := isCloseToSymbol(inputArr, i, num.y)
		if isSymbolClose {
			res = true
			break
		}
	}
	fmt.Println("-->", num.value, res)
	return res
}

func isCloseToSymbol(inputArr []string, x, y int) (res bool) {
	xs := []int{x - 1, x, x + 1}
	ys := []int{y - 1, y, y + 1}
	for _, i := range xs {
		if i < 0 || i >= len(inputArr[0]) {
			continue
		}
		for _, j := range ys {
			if (i == x && j == y) || j < 0 || j >= len(inputArr) {
				continue
			}
			fmt.Println("isSymbol", i, j, inputArr[j][i], isSymbol(inputArr[j][i]))
			res = res || isSymbol(inputArr[j][i])
		}
	}
	return res
}

func isSymbol(char byte) bool {
	return char != '.' && !(char >= 48 && char <= 57)
}
