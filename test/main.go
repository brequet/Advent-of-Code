package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	d, s := getDirAndStep("R 6 (#70c710)")
	fmt.Println("testing getDirAndStep", d, s)
}

func getDirAndStep(line string) (dir string, steps int64) {
	entry := strings.Split(line, "#")[1]
	stepsHex := entry[:5]
	steps, _ = strconv.ParseInt(stepsHex, 16, 64)
	dir = getDir(entry[5])
	return dir, steps
}

func getDir(n byte) string {
	switch n {
	case '0':
		return "R"
	case '1':
		return "D"
	case '2':
		return "L"
	case '3':
		return "U"
	}
	log.Fatal("NOPE")
	return ""
}
