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
		line := scanner.Text()
		res := computeValue(line)
		sum += res
		fmt.Println(line, " - ", res)
	}
	fmt.Print(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func computeValue(input string) int {
	re := regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine")
	res := re.FindAllString(input, -1)
	fmt.Println(res)
	firstNum := getIntFromString(res[0])
	lastNum := getIntFromString(res[len(res)-1])

	result, err := strconv.Atoi(firstNum + lastNum)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func getIntFromString(number string) string {
	dict := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	if res, ok := dict[number]; ok {
		return res
	} else {
		return number
	}
}
