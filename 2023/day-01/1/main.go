package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./go/day-01/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += computeValue(scanner.Text())
	}
	fmt.Print(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func computeValue(input string) int {
	re := regexp.MustCompile("\\d")
	res := re.FindAllString(input, -1)

	firstNum := res[0]
	lastNum := res[len(res)-1]

	result, err := strconv.Atoi(firstNum + lastNum)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
