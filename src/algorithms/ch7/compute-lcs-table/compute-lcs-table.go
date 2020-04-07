package main

import "fmt"

func main() {
	fmt.Println(computeLcsTable("CATCGA", "GTACCGTCA"))
}

func computeLcsTable(strLeft string, strRight string) [][]byte {
	strLeftLen := len(strLeft)
	strRightLen := len(strRight)

	tempArr := make([][]byte, strLeftLen + 1)

	for index := range tempArr {
		tempArr[index] = make([]byte, strRightLen + 1)
	}

	for leftIndex := 1; leftIndex <= strLeftLen; leftIndex++ {
		for rightIndex := 1; rightIndex <= strRightLen; rightIndex++ {
			if strLeft[leftIndex - 1] == strRight[rightIndex - 1] {
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