package main

import "fmt"

func main() {
	arrSize := 6
	arr := make([]int, arrSize)

	arr[0] = 12
	arr[1] = 9
	arr[2] = 3
	arr[3] = 7
	arr[4] = 14
	arr[5] = 11

	originArr := make([]int, arrSize)

	for index, item := range arr {
		originArr[index] = item
	}

	fmt.Printf("origin arr: %v, sorted arr: %v", originArr, insertionSort(arr, arrSize))
}

func insertionSort(arr []int, arrSize int) []int {
	for index := 1; index < arrSize; index++ {
		var item int
		var checkIndex int
		
		item = arr[index]

		for checkIndex = index - 1; checkIndex >= 0; checkIndex-- {
			checkItem := arr[checkIndex]
			if checkItem < item {
				arr[checkIndex + 1] = checkItem
			} else {
				break
			}
		}

		arr[checkIndex + 1] = item
	}

	return arr
}
