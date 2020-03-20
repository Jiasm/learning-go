package main

import (
	"fmt"
	"math"
	"errors"
)

func main() {
	arraySize := 5

	array := make([]int, arraySize)

	for index := 0; index < arraySize; index++ {
		array[index] = index
	}

	result, ok := binarySearch(array, arraySize, 3)

	if ok != nil {
		fmt.Println(ok.Error())
	} else {
		fmt.Printf("find index: %d", result)
	}
}

func binarySearch(array []int, arraySize int, target int) (int, error) {
	current := 0
	index := 0

	size := arraySize - 1
	current = 0

	for ;current < size; {
		index = int(math.Floor(float64((current + size) / 2)))
		
		item := array[index]
		
		if item == target {
			return index, nil
			} else {
			if item > target {
				size = index - 1
			} else {
				current = index + 1
			}
		}
	}

	return -1, errors.New("not found")
}