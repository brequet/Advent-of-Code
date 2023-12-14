package main

import (
	"fmt"
)

func main() {
	slice := []int{5, 4, 2, 69, 69, 69, 65, 64, 65, 63, 68, 69, 69, 65, 64, 65, 63, 68, 69, 69}

	fmt.Println(detectLoop(slice))
}

func detectLoop(nums []int) []int {
	seen := map[int][]int{}

	for i, v := range nums {
		if indices, ok := seen[v]; ok {
			fmt.Println("Already seen", v, i, seen[v])

			if indices[0] == i {
				fmt.Printf("Cycle detected: %v\n", nums[indices[0]:i+1])
			}
		}
		seen[v] = append(seen[v], i)
	}

	fmt.Println("No cycle detected")
	return []int{}
}
