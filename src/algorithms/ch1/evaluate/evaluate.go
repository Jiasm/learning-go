package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evaluate("(1+2)"))
	fmt.Println(evaluate("((1+2)+4)"))
}

func evaluate(text string) float64 {
	var ops []byte
	var vals []float64

	for _, char := range text {
		switch char {
			case '(':
			case '+', '-', '*', '/': 
				ops = append(ops, byte(char))
			case ')':
				op := ops[0]
				ops = ops[1:]
				v, v2 := vals[0], vals[1]
				vals = vals[2:]

				switch op {
					case '+': v = v + v2
					case '-': v = v - v2
					case '*': v = v * v2
					case '/': v = v / v2
				}

				vals = append(vals, v)
			default: {
				val, err := strconv.ParseFloat(string(char), 64)

				if err == nil {
					vals = append(vals, val)
				}
			}
		}
	}

	return vals[0]
}