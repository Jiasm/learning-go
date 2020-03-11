package main

import (
	"fmt"
	"image/color"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	return p.X + q.X
}

func (p *Point) ScaleBy(ratio float64) {
	p.X *= ratio
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint

	cp.X = 1

	fmt.Println(cp.Point.X)

	cp.Point.Y = 2

	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	p.Point.ScaleBy(2)
	q.ScaleBy(2)

	fmt.Println(p.Distance(q.Point))
}