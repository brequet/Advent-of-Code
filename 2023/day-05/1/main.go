package main

import (
	"aoc-2023-go/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	sourceStart int
	destStart   int
	rangeLength int
}

func main() {
	file, err := os.Open("./2023/day-05/input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation []Range
	var rangesPointer *[]Range
	// sum := 0
	// fmt.Println(string(file))
	// fmt.Println(matches)
	reNumRow := regexp.MustCompile(`(\d+)`)
	seeds := []int{}
	mod := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println("Reading line:", line)
		if strings.HasPrefix(line, "seeds:") {
			re := regexp.MustCompile(`(\d+)`)
			matches := re.FindAllString(line, -1)
			for _, i := range matches {
				j, err := strconv.Atoi(i)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, j)
			}
			continue
		} else if strings.HasPrefix(line, "seed-") {
			mod = "seed"
		} else if strings.HasPrefix(line, "soil-") {
			mod = "soil"
		} else if strings.HasPrefix(line, "fertilizer-") {
			mod = "fertilizer"
		} else if strings.HasPrefix(line, "water-") {
			mod = "water"
		} else if strings.HasPrefix(line, "light-") {
			mod = "light"
		} else if strings.HasPrefix(line, "temperature-") {
			mod = "temperature"
		} else if strings.HasPrefix(line, "humidity-") {
			mod = "humidity"
		} else if line == "" {
			mod = ""
		}

		matchesNumRow := utils.Map[string, int](reNumRow.FindAllString(line, -1), func(str string) int {
			n, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			return n
		})
		if !strings.HasPrefix(line, "seeds:") && len(matchesNumRow) > 0 {
			if mod == "seed" {
				rangesPointer = &seedToSoil
			} else if mod == "soil" {
				rangesPointer = &soilToFertilizer
			} else if mod == "fertilizer" {
				rangesPointer = &fertilizerToWater
			} else if mod == "water" {
				rangesPointer = &waterToLight
			} else if mod == "light" {
				rangesPointer = &lightToTemperature
			} else if mod == "temperature" {
				rangesPointer = &temperatureToHumidity
			} else if mod == "humidity" {
				rangesPointer = &humidityToLocation
			} else {
				fmt.Println("PROBLEM SOULD NOT BE HERE")
				rangesPointer = nil
			}

			processInput(matchesNumRow, rangesPointer)
		}
	}

	lowestSeed, lowestLocation := int(^uint(0)>>1), int(^uint(0)>>1)
	for _, seed := range seeds {
		location := getLocationForSeed(seed, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation)
		if location < lowestLocation {
			lowestLocation = location
			lowestSeed = seed
		}
	}
	fmt.Println()
	fmt.Println("Solution: seed", lowestSeed, "for lowest location:", lowestLocation, "/88151870")
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processInput(nums []int, ranges *[]Range) {
	if (*ranges) == nil {
		(*ranges) = []Range{}
	}
	destStart, sourceStart, rangeLength := nums[0], nums[1], nums[2]
	newRange := Range{
		sourceStart: sourceStart,
		destStart:   destStart,
		rangeLength: rangeLength,
	}
	(*ranges) = append((*ranges), newRange)
}

func getLocationForSeed(seed int, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation []Range) int {
	soil := getValueForKey(seed, seedToSoil)
	fertilizer := getValueForKey(soil, soilToFertilizer)
	water := getValueForKey(fertilizer, fertilizerToWater)
	light := getValueForKey(water, waterToLight)
	temperature := getValueForKey(light, lightToTemperature)
	humidity := getValueForKey(temperature, temperatureToHumidity)
	location := getValueForKey(humidity, humidityToLocation)

	return location
}

func getValueForKey(key int, ranges []Range) int {
	var val int
	for _, aRange := range ranges {
		diff := key - aRange.sourceStart
		if diff >= 0 && diff < aRange.rangeLength {
			val = aRange.destStart + diff
		}
	}

	if val == 0 {
		val = key
	}
	return val
}
