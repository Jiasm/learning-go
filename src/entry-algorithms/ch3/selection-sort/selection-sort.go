package main

/**
 * 选择排序： 遍历数组元素，取下标 i
 * 内部遍历数组从 下标 i 到数组结尾，取下标 j
 * 监测到元素 j 大于等于 元素 i，则交换两者位置
 */

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

	fmt.Printf("origin arr: %v sorted arr: %v", originArr, sortArr(arr, arrSize))
}

func sortArr(arr []int, arrSize int) []int {
	for index := 0; index < arrSize; index++ {
		for start := index + 1; start < arrSize; start++ {
			item := arr[index]
			newItem := arr[start]

			if newItem >= item {
				arr[start], arr[index] = item, newItem
			}
		}
	}
	return arr
}