package main

import "fmt"

func main () {
	arrA := [...][2]int{[...]int{1, 0}, [...]int{0, 1}, [...]int{1, 0}}
	arrB := [...][3]int{[...]int{1, 1, 0}, [...]int{0, 1, 1}}

	fmt.Println(boolResult(arrA, arrB))
}

func boolResult(matrixA [3][2]int, matrixB [2][3]int) [][]int {
	matrixResult := make([][]int, 3)

	for aIndex := 0; aIndex < 3; aIndex++ {
		matrixResult[aIndex] = make([]int, 3)
		for bIndex := 0; bIndex < 3; bIndex++ {
			matrixResult[aIndex][bIndex] = checkOr(checkAnd(matrixA[aIndex][0], matrixB[0][aIndex]), checkAnd(matrixA[aIndex][1], matrixB[1][aIndex]))
		}
	}

	return matrixResult
}

func checkAnd(num1, num2 int) int {
	if num1 == 1 && num2 == 1 {
		return 1
	}

	return 0
}

func checkOr(num1, num2 int) int {
	if num1 == 1 || num2 == 1 {
		return 1
	}

	return 0
}