package main

import "fmt"

const max = 9999

type Point byte

type GraphItem struct {
	start 	Point
	end 		Point
	weight 	int
}

type GraphPath struct {
	paths 		map[Point]map[Point]GraphInfo
	inDegree 	map[Point]int
	shortest	map[Point]int
	pred			map[Point]Point
}

type GraphInfo struct {
	point 	Point
	weight 	int
}

func main() {
	gp := GraphPath{
		paths: map[Point]map[Point]GraphInfo{},
		shortest: map[Point]int{},
		pred: map[Point]Point{},
		inDegree: map[Point]int{},
	}

	sourceTag := Point('s')

	gp.append(sourceTag, 't', 6)
	gp.append(sourceTag, 'y', 4)
	gp.append('y', 't', 1)
	gp.append('t', 'y', 2)
	gp.append('t', 'x', 3)
	gp.append('y', 'z', 3)
	gp.append('z', 'x', 5)
	gp.append('x', 'z', 4)
	gp.append('z', sourceTag, 7)

	gp.djikstra(sourceTag)

	gp.print(sourceTag)
}

func (gp GraphPath) append(start Point, end Point, weight int) {
	if gp.paths[start] == nil {
		gp.paths[start] = map[Point]GraphInfo{}
	}
	gp.paths[start][end] = GraphInfo{ end, weight }

	if gp.inDegree[start] == 0 {
		gp.inDegree[start] = 0
	}

	gp.inDegree[end]++
}

func (gp GraphPath) weight(start Point, end Point) int {
	weight := gp.paths[start][end].weight

	return weight
}

func (gp GraphPath) relax(start Point, end Point) {
	currentWeight := gp.shortest[start] + gp.weight(start, end)
	if currentWeight < gp.shortest[end] {
		fmt.Println("call change shortest")
		gp.shortest[end] = currentWeight
		gp.pred[end] = start
	}
}

func (gp GraphPath) djikstra(source Point) {
	var wholeNode []Point

	for path := range gp.inDegree {
		gp.shortest[path] = max

		fmt.Println("call insert")
		wholeNode = append(wholeNode, path)
	}

	gp.shortest[source] = 0

	for ;; {
		pendingLen := len(wholeNode)
		if pendingLen <= 0 {
			break
		}

		min := wholeNode[0]
		minIndex := 0

		fmt.Println("call extract min")
		for index := 1; index < len(wholeNode); index++ {
			cursor := wholeNode[index]
			if gp.shortest[min] > gp.shortest[cursor] {
				min = cursor
				minIndex = index
			}
		}

		for point := range gp.paths[min] {
			gp.relax(min, point)
		}

		wholeNode = append(wholeNode[0:minIndex], wholeNode[minIndex + 1:len(wholeNode)]...)
	}
}

func (gp GraphPath) print(source Point) {
	for path, item := range gp.shortest {
		fmt.Printf("到达路径 %c 最小的权重为: %d\n", path, item)
	}
}
