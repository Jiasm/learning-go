package main

import "fmt"

func main() {
	fmt.Println(power(10, 2))
	fmt.Println(power(2, 3))
}

func power(num, count int) int {
	if count == 0 {
		return 1
	}

	return num * power(num, count - 1)
}