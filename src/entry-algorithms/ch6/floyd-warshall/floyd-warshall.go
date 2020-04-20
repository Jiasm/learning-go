package main

import (
	"fmt"
	"math"
)

func main() {
	g := Graph{
		paths: map[Point]map[Point]Weight{},
		shortest: make([][][]Weight, 5),
		pred: make([][][]Point, 5),
	}

	g.append(Path{ 1, 2, 3 })
	g.append(Path{ 1, 3, 8 })
	g.append(Path{ 2, 4, 1 })
	g.append(Path{ 3, 2, 4 })
	g.append(Path{ 4, 1, 2 })
	g.append(Path{ 4, 3, -5 })

	g.floydWarshall()
	g.print()
}

type Point int

type Weight float64

type Path struct {
	start 	Point
	end 		Point
	weight 	Weight
}

type Graph struct {
	shortest [][][]Weight
	pred		 [][][]Point
	paths 	 map[Point]map[Point]Weight
}

func (g Graph) append(path Path) {
	if _, ok := g.paths[path.start]; !ok {
		g.paths[path.start] = map[Point]Weight{}
	}

	g.paths[path.start][path.end] = path.weight
}

func (g Graph) floydWarshall() {
	pointLen := len(g.paths)

	for start := 1; start <= pointLen; start++ {
		g.shortest[start] = make([][]Weight, pointLen + 1)
		g.pred[start] = make([][]Point, pointLen + 1)
		for end := 1; end <= pointLen; end++ {
			g.shortest[start][end] = make([]Weight, pointLen + 1)
			g.pred[start][end] = make([]Point, pointLen + 1)
			if weight, ok := g.paths[Point(start)][Point(end)]; ok {
				g.shortest[start][end][0] = weight
				g.pred[start][end][0] = Point(start)
			} else {
				g.shortest[start][end][0] = Weight(math.Inf(1))
				g.pred[start][end][0] = 0
			}
		}

		g.shortest[start][start][0] = 0
	}

	for index := 1; index <= pointLen; index++ {
		for start := 1; start <= pointLen; start++ {
			for end := 1; end <= pointLen; end++ {
				newWeight := g.shortest[start][index][index - 1] + g.shortest[index][end][index - 1]

				if g.shortest[start][end][index - 1] > newWeight {
					g.shortest[start][end][index] = newWeight
					g.pred[start][end][index] = g.pred[index][end][index - 1]
				} else {
					g.shortest[start][end][index] = g.shortest[start][end][index - 1]
					g.pred[start][end][index] = g.pred[start][end][index - 1]
				}
			}
		}

		// fmt.Println(g.shortest)
	}
}

func (g Graph) print() {
	fmt.Println(1, 3, g.shortest[1][3][4])
	fmt.Println(1, 4, g.shortest[1][4][4])
	fmt.Println(1, 2, g.shortest[1][2][4])
	fmt.Println(2, 3, g.shortest[2][3][4])
	fmt.Println(2, 1, g.shortest[2][1][4])
	fmt.Println(2, 4, g.shortest[2][4][4])
	fmt.Println(3, 1, g.shortest[3][1][4])
	fmt.Println(3, 2, g.shortest[3][2][4])
	fmt.Println(3, 4, g.shortest[3][4][4])
	fmt.Println(4, 1, g.shortest[4][1][4])
	fmt.Println(4, 2, g.shortest[4][2][4])
	fmt.Println(4, 3, g.shortest[4][3][4])
}