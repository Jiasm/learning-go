package main

import "fmt"

type Item struct {
	score int
	name string
}

func main() {
	arrSize := 9
	arrRange := 12

	arr := make([]Item, arrSize)

	arr[0] = Item{ 3, "jarvis3" }
	arr[1] = Item{ 11, "jarvis6" }
	arr[2] = Item{ 2, "jarvis2" }
	arr[3] = Item{ 6, "jarvis4" }
	arr[4] = Item{ 6, "jarvis5" }
	arr[5] = Item{ 11, "jarvis7" }
	arr[6] = Item{ 0, "jarvis0" }
	arr[7] = Item{ 1, "jarvis1" }
	arr[8] = Item{ 11, "jarvis8" }

	originArr := make([]Item, arrSize)

	for index, item := range arr {
		originArr[index] = item
	}

	fmt.Printf("origin arr: %v, \nsorted arr: %v", originArr, countingSort(arr, arrSize, arrRange))
}

/**
 * 遍历数组，统计元素重复次数，如数组 [1, 1, 2, 3, 6] 统计结果为： [0, 2, 1, 1, 0, 0, 1]
 */
func countKeysEqual(arr []Item, arrSize int, arrRange int) []int {
	tempArr := make([]int, arrRange)

	for _, item := range arr {
		tempArr[item.score]++
	}

	return tempArr
}

/**
 * 统计小于当前元素的元素数量，如数组 [1, 1, 2, 3, 6] 重复次数数组为 [0, 2, 1, 1, 0, 0, 1]
 * 则小于该数的数量统计为 [0, 0, 2, 3, 4, 4, 4]
 */
func countKeysLess(equalArr []int, arrRange int) []int {
	tempArr := make([]int, arrRange)

	for index := 1; index < arrRange; index++ {
		tempArr[index] = tempArr[index - 1] + equalArr[index - 1]
	}

	return tempArr
}

/**
 * 根据小于当前元素的统计数量来重新生成新的数组
 * 数组排序后小于 X 的对应数量数据为 [0, 0, 2, 3, 4, 4, 4]
 * 获取原数组 [1, 1, 2, 3, 6]
 * 遍历原数组
 * 将元素 1 对应在 less 数据中的值取出 得到实际下标 0
 * 将元素 1 赋值到下标 0 的地方
 * 将元素 1 的起始下标位置 +1，用于后续相同元素寻找下标
 * 将元素 1 对应在 less 数据中的值取出 得到实际下标 1
 * 将元素 1 赋值到下标 1 的地方
 * 将元素 1 的起始下标位置 +1，用于后续相同元素寻找下标
 * 将元素 2 对应在 less 数据中的值取出 得到实际下标 2
 * 将元素 2 赋值到下标 2 的地方
 * 将元素 2 的起始下标位置 +1，用于后续相同元素寻找下标
 * 将元素 3 对应在 less 数据中的值取出 得到实际下标 3
 * 将元素 3 赋值在下标 3 的地方
 * 将元素 3 的起始下标位置 +1，用于后续相同元素寻找下标
 * 将元素 6 对应在 less 数据中的值取出 得到实际下标 4
 * 将元素 6 赋值在下标 4 的地方
 * 将元素 6 的起始下标位置 +1，用于后续相同元素寻找下标
 */
func rearrange(arr []Item, lessArr []int, arrSize int, arrRange int) []Item {
	tempArr := make([]Item, arrSize)

	for index := 0; index < arrSize; index++ {
		tempItem := arr[index]

		tempArr[lessArr[tempItem.score]] = tempItem
		lessArr[tempItem.score]++
		// fmt.Printf("less arr: %v\n", lessArr)
		// fmt.Printf("temp arr: %v\n", tempArr)
	}

	return tempArr
}

func countingSort(arr []Item, arrSize int, arrRange int) []Item {
	equalArr := countKeysEqual(arr, arrSize, arrRange)
	fmt.Printf("equal arr: %v\n", equalArr)
	lessArr := countKeysLess(equalArr, arrRange)
	fmt.Printf("less arr: %v\n", lessArr)
	tempArr := rearrange(arr, lessArr, arrSize, arrRange)

	return tempArr
}
