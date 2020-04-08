package main

import "fmt"

func main() {
	leftStr := "CATCGA"
	rightStr := "GTACCGTCA"
	leftStrLen := len(leftStr)
	rightStrLen := len(rightStr)

	tempArr := computeLcsTable(leftStr, rightStr)

	fmt.Println(assembleLcs(leftStr, rightStr, tempArr, leftStrLen, rightStrLen))
}

func computeLcsTable(leftStr string, rightStr string) [][]byte {
	leftStrLen := len(leftStr)
	rightStrLen := len(rightStr)

	tempArr := make([][]byte, leftStrLen + 1)

	for index := range tempArr {
		tempArr[index] = make([]byte, rightStrLen + 1)
	}

	for leftIndex := 1; leftIndex <= leftStrLen; leftIndex++ {
		for rightIndex := 1; rightIndex <= rightStrLen; rightIndex++ {
			if leftStr[leftIndex - 1] == rightStr[rightIndex - 1] {
				tempArr[leftIndex][rightIndex] = tempArr[leftIndex - 1][rightIndex - 1] + 1
			} else {
				leftScore := tempArr[leftIndex][rightIndex - 1]
				rightScore := tempArr[leftIndex - 1][rightIndex]
				if leftScore > rightScore {
					tempArr[leftIndex][rightIndex] = leftScore
				} else {
					tempArr[leftIndex][rightIndex] = rightScore
				}
			}
		}
	}

	return tempArr
}

func assembleLcs(leftStr, rightStr string, tempArr [][]byte, leftIndex, rightIndex int) string {
	if tempArr[leftIndex][rightIndex] == 0 {
		return ""
	} else {
		if leftStr[leftIndex - 1] == rightStr[rightIndex - 1] {
			return assembleLcs(leftStr, rightStr, tempArr, leftIndex - 1, rightIndex - 1) + string(leftStr[leftIndex - 1])
		} else if tempArr[leftIndex][rightIndex - 1] > tempArr[leftIndex - 1][rightIndex] {
			return assembleLcs(leftStr, rightStr, tempArr, leftIndex, rightIndex - 1)
		} else {
			return assembleLcs(leftStr, rightStr, tempArr, leftIndex - 1, rightIndex)
		}
	}
}