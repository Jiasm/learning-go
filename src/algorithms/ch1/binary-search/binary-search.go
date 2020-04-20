package main

import "fmt"

func main() {
	arr := []int{ 1, 2, 3, 4, 5, 6, 7, 8 }

	target := 7

	fmt.Println(binarySearch(arr, target))
}

func binarySearch(arr []int, target int) int {
	arrLen := len(arr)
	
	start := 0
	end := arrLen
	
	for ; start <= end ; {
		mid := start + int((end - start) / 2)

		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}