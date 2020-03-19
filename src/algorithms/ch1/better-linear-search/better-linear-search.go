package main

import (
	"fmt"
	"errors"
)

func main() {
	arraySize := 100000000
	array := make([]int, arraySize)

	for i := 0; i < arraySize; i++ {
		array[i] = i
	}

	target := 98233554

	index, ok := linearSearch(array, arraySize, target)

	if ok != nil {
		fmt.Println(ok.Error())
	} else {
		fmt.Printf("find index: %d\n", index)
	}
}

func linearSearch(array []int, arraySize int, target int) (int, error) {
	for index := 0; index < arraySize; index++ {
		if array[index] == target {
			return index, nil
		}
	}

	return -1, errors.New("not found")
}