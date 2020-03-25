package main

import "fmt"

func main() {
	arrSize := 6
	arr := make([]int, arrSize)

	arr[0] = 1
	arr[1] = 2
	arr[2] = 1
	arr[3] = 2
	arr[4] = 1
	arr[5] = 2

	originArr := make([]int, arrSize)

	for index, item := range arr {
		originArr[index] = item
	}

	fmt.Printf("origin arr: %v, sorted arr: %v", originArr, reallySimpleSort(arr))
}

func reallySimpleSort(arr []int) []int {
	foundCount := 0

	for _, item := range arr {
		if item == 1 {
			foundCount++
		}
	}

	for index := range arr {
		if (index < foundCount) {
			arr[index] = 1
		} else {
			arr[index] = 2
		}
	}

	return arr
}