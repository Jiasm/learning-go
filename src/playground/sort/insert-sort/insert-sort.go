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

	fmt.Println(insertSort(arr))
}

/**
 * 插入排序： 选择一个基准元素，从第二个开始，从头开始检查，如果发现前边有元素比当前元素 大｜小，
 * 将前边的元素位置后挪一位，等到检测到集合头部 or 监测到比当前元素 小｜大，则停止循环，然后将当前遍历到的下标 + 1 赋值为用于检查的当前元素
 * 然后开始用第三个元素为基准元素，开始检查，一直循环到列表尾部
 */
func insertSort(arr []int) []int {
	length := len(arr)
	for index := 1; index < length; index++ {
		cursorItem := arr[index]
		cursor := index - 1
		for ; cursor >= 0; cursor-- {
			if arr[cursor] > cursorItem {
				arr[cursor + 1] = arr[cursor]
			} else {
				break
			}
		}

		arr[cursor + 1] = cursorItem
	}

	return arr
}