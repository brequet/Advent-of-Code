package main

import (
	"aoc-2023-go/utils"
	"fmt"
	"strconv"
	"strings"
)

const (
	DAY             = "15"
	EXPECTED_RESULT = 145
)

func main() {
	input := utils.GetInputAsSlice(fmt.Sprintf("./2023/day-%s/input", DAY))
	fmt.Printf("Input of size: %d\n", len(input))

	result := solve(input[0])
	fmt.Printf("Result %d (expected %d)\n", result, EXPECTED_RESULT)
}

func hash(entry string) (res int) {
	for _, char := range entry {
		res += int(char)
		res *= 17
		res = res % 256
	}
	return res
}

type Lens struct {
	label string
	focal int
}

func solve(input string) int {
	sequences := strings.Split(input, ",")

	boxes := map[int][]Lens{}

	for _, sequence := range sequences {
		if strings.Contains(sequence, "=") {
			e := strings.Split(sequence, "=")
			label, focalStr := e[0], e[1]
			focal, _ := strconv.Atoi(focalStr)
			h := hash(label)
			lenses, ok := boxes[h]
			if !ok {
				lenses = []Lens{}
			}

			isFound := false
			for i, lens := range lenses {
				if lens.label == label {
					lenses[i] = Lens{label: label, focal: focal}
					isFound = true
					break
				}
			}

			if !isFound {
				lenses = append(lenses, Lens{label: label, focal: focal})
			}

			boxes[h] = lenses
		} else {
			label := strings.Split(sequence, "-")[0]
			h := hash(label)
			lenses, ok := boxes[h]
			if !ok {
				continue
			}

			newLenses := []Lens{}

			for _, lens := range lenses {
				if lens.label != label {
					newLenses = append(newLenses, lens)
				}
			}
			boxes[h] = newLenses
		}
	}

	sum := 0
	for key, lenses := range boxes {
		for i := range lenses {
			sum += (key + 1) * (i + 1) * lenses[i].focal
		}
	}

	return sum
}
