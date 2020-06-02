package main

import "fmt"

func main() {
	fmt.Println(logicAnd(1, 0))
	fmt.Println(logicOr(1, 0))
}

func logicAnd(a, b int) int {
	if a == 1 && b == 1 {
		return 1
	}

	return 0
}

func logicOr(a, b int) int {
	if a == 1 || b == 1 {
		return 1
	}

	return 0
}
