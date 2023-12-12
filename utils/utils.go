package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetInputAsSlice(filePath string) (res []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
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

func StrSliceToIntSlice(strSlice []string) []int {
	return Map(strSlice, StrToInt)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(numbers ...int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = result * numbers[i] / GCD(result, numbers[i])
	}
	return result
}

func GetLastElement[T any](list []T) T {
	return list[len(list)-1]
}

func ReverseList[T any](list []T) (res []T) {
	for i := len(list) - 1; i >= 0; i-- {
		res = append(res, list[i])
	}
	return res
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func PrintMat[T any](mat [][]T) {
	for _, row := range mat {
		for _, node := range row {
			fmt.Print(node)
		}
		fmt.Println()
	}
}

func SumAll(nums []int) (res int) {
	for _, n := range nums {
		res += n
	}
	return res
}
