package utils

import (
	"fmt"
	"log"
	"strconv"
)

func Hello() {
	fmt.Println("Hello, World!")
}

func Map[T, O any](things []T, mapper func(thing T) O) []O {
	result := make([]O, 0, len(things))
	for _, thing := range things {
		result = append(result, mapper(thing))
	}
	return result
}

func StrToInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
