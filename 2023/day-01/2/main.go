package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2023/day-01/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// line := scanner.Text()
		// res := computeValue(line)
		// sum += res
		// fmt.Println(line, " - ", res)
	}
	// fmt.Print(sum, " - 54980")
	fmt.Println(computeValue("4threethreegctxg3dmbm1"))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func computeValue(input string) int {
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	parsedWords := findWordsInText(input, words)
	fmt.Println(parsedWords)
	firstNum := getIntFromString(parsedWords[getMinKey(parsedWords)][0])
	lastNum := getIntFromString(parsedWords[getMaxKey(parsedWords)][0])

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

func findWordsInText(text string, words []string) map[int][]string {
	fmt.Println("findWordsInText")
	results := make(map[int][]string)
	for _, word := range words {
		index := strings.Index(text, word)
		for index >= 0 {
			fmt.Println("word", word, "index", index)
			results[index] = append(results[index], word)
			index = strings.Index(text[index+1:], word)
		}
	}
	return results
}

func getMinKey(aMap map[int][]string) int {
	min := 9999
	for k, _ := range aMap {
		if k < min {
			min = k
		}
	}
	return min
}

func getMaxKey(aMap map[int][]string) int {
	max := -1
	for k, _ := range aMap {
		if k > max {
			max = k
		}
	}
	return max
}
