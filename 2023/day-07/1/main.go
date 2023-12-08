package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"sort"
	"strings"
)

const (
	DAY             = "07"
	EXPECTED_RESULT = 253313241
)

func handType(hand string) (int, int) {
	// Define the types of hands and their strengths
	types := []string{"5", "4", "F", "3", "2", "1"}
	strengths := []int{10, 9, 8, 7, 6, 5}

	// Determine the type of hand
	for i, t := range types {
		if strings.Contains(hand, t) {
			return i, strengths[i]
		}
	}

	// If no type is found, it's a high card
	return 5, 5
}

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	sort.Slice(input, func(i, j int) bool {
		hand1, hand2 := strings.Split(input[i], " ")[0], strings.Split(input[j], " ")[0]
		type1, strength1 := handType(hand1)
		type2, strength2 := handType(hand2)

		// Compare the type of hand
		if type1 != type2 {
			return type1 < type2
		}

		// If the hands are of the same type, compare the strength
		if strength1 != strength2 {
			return strength1 > strength2
		}

		// If the hands are of the same type and strength, compare the individual cards
		for i := 0; i < 5; i++ {
			card1 := hand1[i]
			card2 := hand2[i]

			// Compare the cards
			if card1 != card2 {
				return card1 < card2
			}
		}

		// If the hands are identical, they are equal
		return false
	})

	result := 0
	for i, hand := range input {
		cardsAndBid := strings.Split(hand, " ")
		// fmt.Printf("Processing: %s\n", cardsAndBid[0])
		// fmt.Printf("Processing %d: %s (%s)\n", i, cardsAndBid[1], cardsAndBid[0])
		result += (i + 1) * utils.StrToInt(cardsAndBid[1])
	}

	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func countCards(hand string) map[rune]int {
	counts := make(map[rune]int)
	for _, card := range hand {
		counts[card]++
	}
	return counts
}

func checkHand(hand string) int {
	counts := countCards(hand)
	// fmt.Println("\t", hand, counts)
	for _, count := range counts {
		if count == 5 {
			return 6 // 5 of a kind
		} else if count == 4 {
			return 5 // 4 of a kind
		} else if count == 3 {
			if len(counts) == 2 {
				return 4 // full house
			} else {
				return 3 // 3 of a kind
			}
		} else if count == 2 {
			if len(counts) == 2 {
				return 3 // 3 of a kind
			} else if len(counts) == 3 {
				return 2 // two pair
			} else {
				return 1 // one pair
			}
		}
	}
	return 0 // high card (nothing)
}

func getCardPosition(char byte, cards []byte) int {
	for i, v := range cards {
		if v == char {
			return i
		}
	}
	fmt.Println("ERROR: SHOULD NOT BE HERE ->", char)
	return -1
}
