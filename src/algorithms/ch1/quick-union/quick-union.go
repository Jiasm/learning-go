package main

import "fmt"

type UF struct {
	count int
	points []int
}

func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) find(point int) int {
	target := point

	for target != uf.points[point] {
		target = uf.points[target]
	}

	return target
}

func (uf *UF) connected(q, p int) bool {
	return uf.find(q) == uf.find(p)
}

func (uf *UF) union(q, p int) {
	pRoot, qRoot := uf.find(p), uf.find(q)

	if pRoot == qRoot {
		return
	}

	uf.points[pRoot] = qRoot

	uf.count--
}

func main() {
	count := 10
	points := make([]int, count)

	for index := range points {
		points[index] = index
	}

	uf := UF{ count, points }

	uf.union(4, 3)
	uf.union(3, 8)
	uf.union(6, 5)
	uf.union(9, 4)
	uf.union(2, 1)
	uf.union(5, 0)
	uf.union(7, 1)
	uf.union(6, 1)

	fmt.Printf("%d components\n", uf.Count())
}