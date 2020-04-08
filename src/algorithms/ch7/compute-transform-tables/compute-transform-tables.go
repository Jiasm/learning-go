package main

import "fmt"

const (
	COPY 		int = -1
	REPLACE	int = 1
	DELETE	int = 2
	INSERT	int = 2
)

type Op struct {
	str 		string
	opType	string
	replaceStr string
}

func main() {
	letStr := "ACAAGC"
	rightStr := "CCGT"
	_, op := computeTransformTables(letStr, rightStr, COPY, REPLACE, DELETE, INSERT)

	fmt.Println(assembleTransformation(op, len(letStr), len(rightStr)))
}

func computeTransformTables(leftStr, rightStr string, copyScore, replaceScore, deleteScore, insertScore int) (cost [][]int, op [][]Op) {
	leftLen := len(leftStr)
	rightLen := len(rightStr)

	cost = make([][]int, leftLen + 1)
	op = make([][]Op, leftLen + 1)

	for index := range cost {
		cost[index] = make([]int, rightLen + 1)
		op[index] = make([]Op, rightLen + 1)
	}

	for index := 1; index <= leftLen; index++ {
		cost[index][0] = index * deleteScore
		op[index][0] = Op{ string(leftStr[index - 1]), "delete", "" }
	}

	for index := 1; index <= rightLen; index++ {
		cost[0][index] = index * insertScore
		op[0][index] = Op{ string(rightStr[index - 1]), "insert", "" }
	}

	for outerIndex := 1; outerIndex <= leftLen; outerIndex++ {
		for innerIndex := 1; innerIndex <= rightLen; innerIndex++ {
			if leftStr[outerIndex - 1] == rightStr[innerIndex - 1] {
				cost[outerIndex][innerIndex] = cost[outerIndex - 1][innerIndex - 1] + copyScore
				op[outerIndex][innerIndex] = Op{ string(leftStr[outerIndex - 1]), "copy", "" }
			} else {
				cost[outerIndex][innerIndex] = cost[outerIndex - 1][innerIndex - 1] + replaceScore
				op[outerIndex][innerIndex] = Op{ string(leftStr[outerIndex - 1]), "replace", string(rightStr[innerIndex - 1]) }
			}

			if cost[outerIndex - 1][innerIndex] + deleteScore < cost[outerIndex][innerIndex] {
				cost[outerIndex][innerIndex] = cost[outerIndex - 1][innerIndex] + deleteScore
				op[outerIndex][innerIndex] = Op{ string(leftStr[outerIndex - 1]), "delete", "" }
			}

			if cost[outerIndex][innerIndex - 1] + insertScore < cost[outerIndex][innerIndex] {
				cost[outerIndex][innerIndex] = cost[outerIndex][innerIndex - 1] + insertScore
				op[outerIndex][innerIndex] = Op{ string(rightStr[innerIndex - 1]), "insert", "" }
			}
		}
	}

	return
}

func assembleTransformation(op [][]Op, outerIndex, innerIndex int) string {
	if outerIndex == 0 && innerIndex == 0 {
		return ""
	}

	opType := op[outerIndex][innerIndex].opType
	str := op[outerIndex][innerIndex].str
	replaceStr := op[outerIndex][innerIndex].replaceStr
	if opType == "copy" || opType == "replace" {
		return assembleTransformation(op, outerIndex - 1, innerIndex - 1) + "|" + str + opType + replaceStr
	} else if opType == "delete" {
		return assembleTransformation(op, outerIndex - 1, innerIndex) + "|" + str + opType + replaceStr
	} else {
		return assembleTransformation(op, outerIndex, innerIndex - 1) + "|" + str + opType + replaceStr
	}
}