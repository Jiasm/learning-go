package main

import "fmt"

type Point struct{ X, Y int }

func (p Point) Distance() int {
	return p.X * p.Y
}

type Path []Point

func (path Path) Distance() int {
	sum := 0

	for i := range path {
		if i > 0 {
			sum += path[i].Distance()
		}
	}

	return sum
}


func main() {
	perim := Path{
		{ 1, 1},
		{ 2, 2},
		{ 3, 3},
		{ 4, 4},
	}
	fmt.Println(perim.Distance())
}