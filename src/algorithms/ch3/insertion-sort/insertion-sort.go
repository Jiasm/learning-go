package main

import "fmt"

/**
 * 插入排序：
 * 遍历数组元素从 1 开始到结束，取下标 i
 * 反向遍历数组元素 从 i - 1 到数组开始，取下标 j
 * 暂存元素 i 为 t
 * 如果 元素 j 小于 t，则将下标 j + 1 的值设置为元素 j
 * 如果 元素 j 大于等于 t，则终止循环
 * 内层循环结束后将 j + 1 下标赋值为 t
 */

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
