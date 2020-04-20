package main

import "fmt"

func main() {
	arrSize := 8
	arrRange := 36

	arr := make([]string, arrSize)

	arr[0] = "F6"
	arr[1] = "E5"
	arr[2] = "R6"
	arr[3] = "X6"
	arr[4] = "X2"
	arr[5] = "T5"
	arr[6] = "F2"
	arr[7] = "T3"

	originArr := make([]string, arrSize)

	for index, item := range arr {
		originArr[index] = item
	}

	var tempArr = arr

	for index := len(arr[0]) - 1; index >= 0; index-- {
		tempArr = countingSort(tempArr, arrSize, arrRange, index)
	}

	fmt.Printf("origin arr: %v, sorted arr: %v", originArr, tempArr)
}

/**
 * 遍历数组，统计元素重复次数，如数组 [1, 1, 2, 3, 6] 统计结果为： [0, 2, 1, 1, 0, 0, 1]
 */
func countKeysEqual(arr []string, arrSize int, arrRange int, charIndex int) []int {
	tempArr := make([]int, arrRange)

	for _, item := range arr {
		tempArr[value(item, charIndex)]++
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
func rearrange(arr []string, lessArr []int, arrSize int, arrRange int, charIndex int) []string {
	tempArr := make([]string, arrSize)

	for index := 0; index < arrSize; index++ {
		tempItem := arr[index]

		tempArr[lessArr[value(tempItem, charIndex)]] = tempItem
		lessArr[value(tempItem, charIndex)]++
		// fmt.Printf("less arr: %v\n", lessArr)
		// fmt.Printf("temp arr: %v\n", tempArr)
	}

	return tempArr
}

func value(str string, index int) int {
	valueMap := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
		'G': 16,
		'H': 17,
		'I': 18,
		'J': 19,
		'K': 20,
		'L': 21,
		'M': 22,
		'N': 23,
		'O': 24,
		'P': 25,
		'Q': 26,
		'R': 27,
		'S': 28,
		'T': 29,
		'U': 30,
		'V': 31,
		'W': 32,
		'X': 33,
		'Y': 34,
		'Z': 35,
	}

	return valueMap[str[index]]
}

func countingSort(arr []string, arrSize int, arrRange int, charIndex int) []string {
	equalArr := countKeysEqual(arr, arrSize, arrRange, charIndex)
	fmt.Printf("equal arr: %v\n", equalArr)
	lessArr := countKeysLess(equalArr, arrRange)
	fmt.Printf("less arr: %v\n", lessArr)
	tempArr := rearrange(arr, lessArr, arrSize, arrRange, charIndex)

	return tempArr
}
