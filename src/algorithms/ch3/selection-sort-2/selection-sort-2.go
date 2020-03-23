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

	fmt.Printf("origin arr: %v, sorted arr: %v", originArr, selectionSort(arr, arrSize))
}

func selectionSort(arr []int, arrSize int) []int {
	for index := 0; index < arrSize; index++ {
		smallest := arr[index]
		smallestIndex := index
		for start := index + 1; start < arrSize; start++ {
			if arr[start] > smallest {
				smallest = arr[start]
				smallestIndex = start
			}
		}

		arr[smallestIndex], arr[index] = arr[index], smallest
	}

	return arr
}