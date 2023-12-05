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
	file, err := os.Open("./day-03/input")
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

	gears := make(map[string][]int)
	for _, num := range nums {
		computeGear(inputArr, num, gears)
	}

	sum := 0
	for _, nums := range gears {
		if len(nums) > 1 {
			sum += multArray(nums)
		}
	}
	fmt.Println("result: ", sum, " / ", 84584891)
}

func multArray(ints []int) int {
	numsSum := 1
	for _, num := range ints {
		numsSum *= num
	}
	return numsSum
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

func computeGear(inputArr []string, num Number, gears map[string][]int) {
	for i := num.x; i < num.x+len(num.value); i++ {
		gearCoordinate := getGearCoordinateIfOneClose(inputArr, i, num.y, gears)
		if gearCoordinate != "" {
			if _, ok := gears[gearCoordinate]; !ok {
				gears[gearCoordinate] = []int{}
			}
			result, err := strconv.Atoi(num.value)
			if err != nil {
				log.Fatal(err)
			}
			gears[gearCoordinate] = append(gears[gearCoordinate], result)
			break
		}
	}
}

func getGearCoordinateIfOneClose(inputArr []string, x, y int, gears map[string][]int) (gearCoordinate string) {
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
			if isGear(inputArr[j][i]) {
				return fmt.Sprintf("%v,%v", i, j)
			}
		}
	}
	return ""
}

func isGear(char byte) bool {
	return char == '*'
}
