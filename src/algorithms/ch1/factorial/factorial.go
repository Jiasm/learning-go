package main

import "fmt"

func main() {
	fmt.Printf("value is %d", factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n - 1)
	}
}