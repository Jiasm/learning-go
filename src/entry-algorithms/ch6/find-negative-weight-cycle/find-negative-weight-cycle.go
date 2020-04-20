package main

import "fmt"

type Point byte

type GraphData struct {
	start		Point
	end			Point
	weight	int
}

type Graph struct {
	paths 		map[Point]map[Point]int
	inDegree	map[Point]int
	shortest	map[Point]int
	pred			map[Point]Point
}

const max = 999

func main() {
	g := Graph{
		paths: map[Point]map[Point]int{},
		shortest: map[Point]int{},
		pred: map[Point]Point{},
		inDegree: map[Point]int{},
	}

	startNode := Point('s')

	g.append(GraphData{ startNode, 't', -6 })
	g.append(GraphData{ startNode, 'y', 7 })
	g.append(GraphData{ 't', 'x', 5 })
	g.append(GraphData{ 't', 'z', -4 })
	g.append(GraphData{ 't', 'y', 8 })
	g.append(GraphData{ 'y', 'x', -3 })
	g.append(GraphData{ 'y', 'z', 9 })
	g.append(GraphData{ 'x', 't', -2 })
	g.append(GraphData{ 'z', 's', 2 })
	g.append(GraphData{ 'z', 'x', 7 })

	g.bellmanFord(startNode)
	// g.print(startNode)
	for _, item := range g.findNegativeWeightCycle() {
		fmt.Printf("%c ", item)
	}
	fmt.Println()
}

func (g Graph) append(item GraphData) {
	if g.paths[item.start] == nil {
		g.paths[item.start] = map[Point]int{}
	}

	g.paths[item.start][item.end] = item.weight
	
	// make sure has top level element
	if g.inDegree[item.start] == 0 {
		g.inDegree[item.start] = 0
	}

	g.inDegree[item.end]++
}

func (g Graph) weight(start Point, end Point) int {
	for point, weight := range g.paths[start] {
		if point == end {
			return weight
		}
	}

	return 0
}

func (g Graph) relax(start Point, end Point) {
	newWeight := g.weight(start, end) + g.shortest[start]
	// fmt.Printf("join relax start: %c end: %c newWeight: %d cursor shortest: %d\n", start, end, newWeight, g.shortest[end])
	if newWeight < g.shortest[end] {
		g.shortest[end] = newWeight
		g.pred[end] = start
	}
}

func (g Graph) bellmanFord(source Point) {
	var wholePoint []Point
	for point := range g.paths {
		g.shortest[point] = max

		wholePoint = append(wholePoint, point)
	}

	g.shortest[source] = 0

	for range wholePoint[:len(wholePoint) - 1] {
		for _, point := range wholePoint {
			for endPoint := range g.paths[point] {
				g.relax(point, endPoint)
			}
		}
	}
}

func (g Graph) findNegativeWeightCycle() []Point {
	for startPoint := range g.paths {
		for endPoint := range g.paths[startPoint] {
			newWeight := g.weight(startPoint, endPoint) + g.shortest[startPoint]
			if newWeight < g.shortest[endPoint] {
				visited := map[Point]bool{}
				
				cursor := startPoint

				cycle := []Point{}
				for ;; {
					if visited[cursor] == false {
						visited[cursor] = true
						cycle = append([]Point{ cursor }, cycle...)
						cursor = g.pred[cursor]
					} else {
						break
					}
				}

				// endPoint = g.pred[cursor]

				// cycle = append(cycle, cursor)

				// for ;; {
				// 	if endPoint != cursor {
				// 		cycle = append([]Point{ endPoint }, cycle...)
				// 		endPoint = g.pred[endPoint]
				// 	} else {
				// 		break
				// 	}
				// }

				return cycle
			}
		}
	}

	return []Point{}
}

func (g Graph) print(source Point) {
	for path, item := range g.shortest {
		fmt.Printf("到达路径 %c 最小的权重为: %d\n", path, item)
	}
}
