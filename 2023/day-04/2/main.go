package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./2023/day-04/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	accCardToCount := map[int]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += getCardPoint(scanner.Text(), accCardToCount)
	}
	fmt.Println("result: ", sum, "30")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getCardPoint(line string, accCardToCount map[int]int) int {
	r := strings.Split(line, ":")[1]
	reCard := regexp.MustCompile("(\\d+):")
	card := reCard.FindStringSubmatch(line)[1]

	cardIndex, err := strconv.Atoi(card)
	if err != nil {
		log.Fatal(err)
	}

	numberOfCard, ok := accCardToCount[cardIndex]
	if ok {
		numberOfCard++
	} else {
		numberOfCard = 1
	}

	pipe := strings.Split(r, "|")
	re := regexp.MustCompile("(\\d+)")
	winningNumbers := re.FindAllString(pipe[0], -1)
	myNumbers := re.FindAllString(pipe[1], -1)

	matchingCount := 0
	for _, myNum := range myNumbers {
		for _, winningNum := range winningNumbers {
			if myNum == winningNum {
				matchingCount++
				break
			}
		}
	}

	for j := 0; j < numberOfCard; j++ {
		for i := cardIndex + 1; i < cardIndex+matchingCount+1; i++ {
			if _, ok := accCardToCount[i]; !ok {
				accCardToCount[i] = 0
			}
			accCardToCount[i] = accCardToCount[i] + 1
		}
	}

	return numberOfCard
}
