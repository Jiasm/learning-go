package main

import "fmt"

func main() {
	arr := make([]int, 5)
	arr[0] = 3
	arr[1] = 8
	arr[2] = 4
	arr[3] = 2
	arr[4] = 1

	fmt.Println(quickSort(arr, 0, 4))
}

func quickSort(arr []int, start, end int) []int {
	if start >= end {
		return arr
	}

	distance := getDistance(arr, start, end)

	quickSort(arr, 0, distance)
	quickSort(arr, distance + 1, end)

	return arr
}

func getDistance(arr []int, start, end int) int {
	target := arr[start]

	for start < end {
		for ; start < end && target < arr[end]; {
			end--
		}

		arr[end], arr[start] = arr[start], arr[end]

		for ; start < end && target > arr[start]; {
			start++
		}

		arr[start], arr[end] = arr[end], arr[start]
	}

	arr[start] = target

	return start
}