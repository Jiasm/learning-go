package main

import (
	"fmt"
	"errors"
)

func main() {
	arraySize := 100000
	array := make([]int, arraySize)

	for i := 0; i < arraySize; i++ {
		array[i] = i
	}

	target := 98554

	index, ok := linearSearch(array, arraySize, target)

	if ok != nil {
		fmt.Println(ok.Error())
	} else {
		fmt.Printf("found index: %d\n", index)
	}
}

func linearSearch(array []int, arraySize int, target int) (int, error) {
	answer := -1
	for index := 0; index < arraySize; index++ {
		if array[index] == target {
			answer = index
		}
	}

	if answer != -1 {
		return answer, nil
	} else {
		return 0, errors.New("not found")
	}
}