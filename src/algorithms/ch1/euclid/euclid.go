package main

import "fmt"

func main() {
	fmt.Println(euclid(55, 100))
	fmt.Println("====")
	fmt.Println(euclid(100, 55))
}

func euclid(num1, num2 int) int {
	fmt.Println(num1, num2)
	if num2 == 0 {
		return num1
	}
	gcd := num1 % num2

	return euclid(num2, gcd)
}