package main

import (
	"aoc-2023-go/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./day-02/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += getInputPower(scanner.Text())
	}

	fmt.Println("res: ", sum)
	utils.Hello()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getInputPower(input string) int {
	r, g, b := getMaxRgb(input)
	return r * g * b
}

func getMaxRgb(input string) (r int, g int, b int) {
	return getMaxFromColor(input, "red"), getMaxFromColor(input, "green"), getMaxFromColor(input, "blue")
}

func getMaxFromColor(input, color string) (max int) {
	re := regexp.MustCompile(fmt.Sprintf("(\\d+) %s", color))
	res := re.FindAllStringSubmatch(input, -1)

	for _, v := range res {
		conv, err := strconv.Atoi(v[1])
		if err != nil {
			log.Fatal(err)
		}

		if conv > max {
			max = conv
		}
	}
	return max
}
