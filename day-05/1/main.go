package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

type Map map[int]int

func main() {
	// file, err := os.Open("./day-04/input")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// var seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation Map
	// sum := 0
	// scanner := bufio.NewScanner(file)
	file, err := os.ReadFile("./day-05/input")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(file))
	groupRe := regexp.MustCompile("(\\w+)-to-\\w+ map:\r")
	matches := groupRe.FindAllStringSubmatch(string(file), -1)
	fmt.Println(matches)

	// for scanner.Scan() {
	// 	// sum += getCardPoint(scanner.Text())
	// 	processInput(scanner.Text(), seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation)
	// }
	// fmt.Println("result: ", sum, "25231")

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}

func processInput(input string, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemperature, temperatureToHumidity, humidityToLocation Map) {

}
