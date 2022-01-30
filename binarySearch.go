package main

import (
	"fmt"
)

func binary_search(numbers []int64, valueToFind int64, min int64, max int64) int64 {
	mid := (min + max) / 2
	numberInMid := numbers[mid]
	if numberInMid == valueToFind {
		fmt.Printf("Number found at position : %d", mid)
		return mid
	} else {
		if numberInMid > valueToFind {
			max = mid - 1
			return binary_search(numbers, valueToFind, min, max)
		} else {
			min = mid + 1
			return binary_search(numbers, valueToFind, min, max)
		}
	}
}

func main() {
	var numbers []int64
	numbers = []int64{2, 3, 5, 7, 11, 13}
	binary_search(numbers, 13, 0, int64(len(numbers)-1))
}