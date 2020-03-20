package main

import (
	"errors"
	"math"
	"fmt"
)

func main() {
	arraySize := 5
	array := make([]int, arraySize)

	for index := 0; index < arraySize; index++ {
		array[index] = index
	}

	result, ok := recursiveBinarySearch(array, 0, arraySize - 1, 3)

	if ok != nil {
		fmt.Println(ok.Error())
	} else {
		fmt.Printf("find index: %d", result)
	}
}

func recursiveBinarySearch(array []int, start int, end int, target int) (int, error) {
	if (start > end) {
		return -1, errors.New("not found")
	} else {
		current := int(math.Floor(float64(start + end) / 2))

		if array[current] == target {
			return current, nil
		} else {
			if current > target {
				return recursiveBinarySearch(array, start, current - 1, target)
			} else {
				return recursiveBinarySearch(array, current + 1, end, target)
			}
		}
	}
}