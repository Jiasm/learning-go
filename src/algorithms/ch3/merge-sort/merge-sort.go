package main

import (
	"fmt"
	"math"
)

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

	fmt.Printf("origin arr: %v, sorted arr: %v\n", originArr, mergeSort(arr, 0, arrSize - 1))
}

func mergeSort(arr []int, start int, end int) []int {
	if (start >= end) {
		return arr
	}

	sliceIndex := int(math.Ceil(float64((end + start) / 2)))

	// tempArr := []int{}

	mergeSort(arr, start, sliceIndex)

	mergeSort(arr, sliceIndex + 1, end)

	merge(arr, start, sliceIndex, end)

	return arr
}

func merge(arr []int, start int, sliceIndex int, end int) {
	tempLeft := arr[start:sliceIndex + 1]
	arrLeft := make([]int, len(tempLeft) + 1)

	for index, item := range tempLeft {
		arrLeft[index] = item
	}

	// 如果正序需要手动设置边界值
	// arrLeft[len(tempLeft)] = 999

	tempRight := arr[sliceIndex + 1: end + 1]
	arrRight := make([]int, len(tempRight) + 1)

	for index, item := range tempRight {
		arrRight[index] = item
	}

	// 如果正序需要手动设置边界值
	// arrRight[len(tempRight)] = 999

	leftIndex := 0
	rightIndex := 0

	for index := start; index <= end; index++ {
		if arrLeft[leftIndex] >= arrRight[rightIndex] {
			fmt.Printf("----\njoin left %v %v %v %d %d: \norigin: %d origin index: %d \nnew: %d new index: %d\n----\n", arr, arrLeft, arrRight, start, end, arr[index], index, arrLeft[leftIndex], leftIndex)
			arr[index] = arrLeft[leftIndex]
			leftIndex++
		} else {
			fmt.Printf("----\njoin right %v %v %v %d %d: \norigin: %d origin index: %d \nnew: %d new index: %d\n----\n", arr, arrLeft, arrRight, start, end, arr[index], index, arrRight[rightIndex], rightIndex)
			arr[index] = arrRight[rightIndex]
			rightIndex++
		}
	}
}