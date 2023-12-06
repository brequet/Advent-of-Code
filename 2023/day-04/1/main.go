package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("./2023/day-04/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += getCardPoint(scanner.Text())
	}
	fmt.Println("result: ", sum, "25231")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getCardPoint(line string) int {
	r := strings.Split(line, ":")[1]
	pipe := strings.Split(r, "|")

	re := regexp.MustCompile("(\\d+)")

	winningNumbers := re.FindAllString(pipe[0], -1)
	myNumbers := re.FindAllString(pipe[1], -1)

	fmt.Println("winning", winningNumbers)
	fmt.Println("my", myNumbers)

	matchingCount := 0
	for _, myNum := range myNumbers {
		for _, winningNum := range winningNumbers {
			if myNum == winningNum {
				matchingCount++
				break
			}
		}
	}

	if matchingCount > 0 {
		return int(math.Pow(2, float64(matchingCount)-1))
	}

	return 0
}
