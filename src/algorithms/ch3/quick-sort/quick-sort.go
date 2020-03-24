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

	fmt.Printf("origin arr: %v, sorted arr: %v", originArr, quickSort(arr, 0, arrSize - 1))
}

func quickSort(arr []int, start int, end int) []int {
	if start >= end {
		return arr
	}

	sliceIndex := partition(arr, start, end)

	quickSort(arr, start, sliceIndex - 1)
	quickSort(arr, sliceIndex + 1, end)

	return arr
}

func partition(arr []int, start int, end int) int {
	originIndex := start

	for checkIndex := start; checkIndex < end; checkIndex++ {
		if arr[checkIndex] > arr[end] {
			fmt.Printf("排序范围 %d %d 监测到节点%d大于基准节点%d，交换%d与%d %v\n", start, end, arr[checkIndex], arr[end], originIndex, checkIndex, arr)
			arr[originIndex], arr[checkIndex] = arr[checkIndex], arr[originIndex]
			originIndex++
			fmt.Printf("完成后的数组%v\n", arr)
		}
	}

	fmt.Printf("排序范围 %d %d 将基准节点%d与节点移动次数对应下标%d的节点%d交换位置，保证基准节点在中间 %v\n", start, end, arr[end], originIndex, arr[originIndex], arr)
	arr[originIndex], arr[end] = arr[end], arr[originIndex]
	fmt.Printf("完成后的数组%v\n", arr)

	return originIndex
}