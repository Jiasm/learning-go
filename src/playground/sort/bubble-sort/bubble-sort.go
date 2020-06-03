package main

import "fmt"

func main() {
	arr := make([]int, 6)

	arr[0] = 3
	arr[1] = 8
	arr[2] = 1
	arr[3] = 5
	arr[4] = 2
	arr[5] = 4

	fmt.Println(bubbleSort(arr))
}

/**
 * 从第一个元素开始，依次与后边的元素进行排序检查，如果发现当前元素 大于｜小于 后边的元素，交换两者的位置，并继续向后检查，确保最后一个元素中存储的是 最大｜最小 的元素
 * 然后开始下一次排序检查，下次排序检查的范围为 0 - length - 排序执行的次数（因为经过上次的检查，最后 n 个元素已经是确定的 最最大｜最小 的元素了）
 */
func bubbleSort(arr []int) []int {
	length := len(arr)

	for index := 0; index < length; index++ {
		for cursor := 0; cursor < length - index - 1; cursor++ {
			if arr[cursor] > arr[cursor + 1] {
				arr[cursor], arr[cursor + 1] = arr[cursor + 1], arr[cursor]
			}
		}
	}

	return arr
}