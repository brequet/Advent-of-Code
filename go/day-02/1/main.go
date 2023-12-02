package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const RED_CUBE, GREEN_CUBE, BLUE_CUBE = 12, 13, 14

func main() {
	file, err := os.Open("./go/day-02/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum, i := 0, 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if isInputValid(scanner.Text()) {
			sum += i
		}
		i++
	}

	fmt.Print("res: ", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isInputValid(input string) bool {
	r, g, b := getMaxRgb(input)
	if r > RED_CUBE || g > GREEN_CUBE || b > BLUE_CUBE {
		return false
	} else {
		return true
	}
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
