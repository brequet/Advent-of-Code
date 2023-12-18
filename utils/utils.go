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

func GetInput(filePath string) (input [][]byte) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}

func GetInputInt(filePath string) (input [][]int) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []int{}
		for _, char := range scanner.Text() {
			n, _ := strconv.Atoi(string(char))
			row = append(row, n)
		}
		input = append(input, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
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

func MatToStr[T any](mat [][]T, toStr func(t T) string) (res string) {
	for _, row := range mat {
		for _, node := range row {
			res += toStr(node)
		}
		res += "\n"
	}
	return res
}

func SumAll(nums []int) (res int) {
	for _, n := range nums {
		res += n
	}
	return res
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func ByteToStr(b byte) string {
	return string(b)
}

func CopyMat[T any](matrix [][]T) [][]T {
	n := len(matrix)
	m := len(matrix[0])
	duplicate := make([][]T, n)
	data := make([]T, n*m)
	for i := range matrix {
		start := i * m
		end := start + m
		duplicate[i] = data[start:end:end]
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func MatTranspose[T any](matrix [][]T) [][]T {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]T, cols)
	for i := range transposed {
		transposed[i] = make([]T, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func MatRotate[T any](matrix [][]T) [][]T {
	// Transpose the matrix
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}

	// Reverse each row
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			matrix[i][j], matrix[i][len(matrix)-j-1] = matrix[i][len(matrix)-j-1], matrix[i][j]
		}
	}

	return matrix
}
