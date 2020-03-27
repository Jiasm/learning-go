package main

/**
 * 有向无环图关键路径查询
 * 1. 所有顶点前置的顶点数量
 * 2. 将所有顶点除了起始元素以外的 最小权重 设置为 MAX（关键的一步，不然遍历引用顶点调用 relax 将没有效果，因为所有的最小权重默认都是 0，无法正常执行）
 * 3. 将所有无引用的顶点放入数组中，并存储他所引用的各个顶点
 * 4. 遍历所有他所引用的顶点，调用 relax 获取最短的路径权重
 *    1. 如果当前顶点 到 a 顶点的权重 + 当前顶点最小权重 小于 顶点 a 最小权重，则更新顶点 a 最小权重为 当前顶点 到 a 顶点的权重 + 当前顶点最小权重
 *    2. 更新顶点 a 最小权重对应引用路径顶点为 a
 * 5. 减少当前顶点所引用的各个顶点对应的引用数量，如果引用数量为 0 则认为该节点已经没有依赖，将该节点推入无引用顶点数组
 * 6. 直到 无引用顶点 数组中已经没有数据，停止算法
 */

import (
	"fmt"
)

const max = 9999

func main() {
	gp := GraphPath{
		paths: map[byte][]PathItem{},
		shortest: map[byte]int{},
		pred: map[byte]byte{},
		inDegree: map[byte]int{},
	}

	topNode := byte('s')

	// gp.append('r', 's', 5)
	// gp.append('r', 't', 3)
	gp.append(topNode, 't', 2)
	gp.append(topNode, 'x', 6)
	gp.append('t', 'x', 7)
	gp.append('t', 'y', 4)
	gp.append('t', 'z', 2)
	gp.append('x', 'y', -1)
	gp.append('x', 'z', 1)
	gp.append('y', 'z', -2)

	gp.dagShortestPaths(topNode)

	gp.print(topNode)
}

type GraphItem struct {
	start byte
	end byte
	weight int
}

type GraphPath struct {
	paths map[byte][]PathItem
	shortest map[byte]int
	pred map[byte]byte
	inDegree map[byte]int
}

type PathItem struct {
	path byte
	weight int
}

func (gp GraphPath) append(start byte, end byte, weight int) {
	if gp.inDegree[start] == 0 {
		gp.inDegree[start] = 0
	}
	gp.inDegree[end]++
	gp.paths[start] = append(gp.paths[start], PathItem{ end, weight })
}

func (gp GraphPath) weight(start byte, end byte) int {
	paths := gp.paths[start]
	for _, item := range paths {
		if item.path == end {
			return item.weight
		}
	}
	return 0
}

func (gp GraphPath) relax(start byte, end byte) {
	if gp.shortest[start] + gp.weight(start, end) < gp.shortest[end] {
		gp.shortest[end] = gp.shortest[start] + gp.weight(start, end)
		gp.pred[end] = start
	}
}

func (gp GraphPath) dagShortestPaths(source byte) {
	var next []byte
	for path, item := range gp.inDegree {
		// ignore virtual node
		// if path != 'r' {
		gp.shortest[path] = max
		// }
		if item == 0 {
			next = append(next, path)
		}
	}
	gp.shortest[source] = 0
	var result []byte
	var nextItem byte
	for ;; {
		if len(next) == 0 {
			break
		}
		nextItem, next = next[len(next) - 1], next[:len(next) - 1]
		result = append(result, nextItem)
		for _, pathInfo := range gp.paths[nextItem] {
			gp.inDegree[pathInfo.path]--
			if gp.inDegree[pathInfo.path] == 0 {
				next = append(next, pathInfo.path)
			}
			gp.relax(nextItem, pathInfo.path)
		}
	}
}

func (gp GraphPath) print(source byte) {
	for path, score := range gp.shortest {
		if path != 's' {
			fmt.Printf("从顶点 %c 到 %c 的最小权重为 %d 路径为：", source, path, score)

			tempPath := path
			var tempPathList []byte
			for ;; {
				tempPathList = append(tempPathList, tempPath)

				tempPath = gp.pred[tempPath]

				if (tempPath == source) {
					break
				}
			}

			for i, j := 0, len(tempPathList)-1; i < j; i, j = i + 1, j - 1 {
				tempPathList[i], tempPathList[j] = tempPathList[j], tempPathList[i]
			}

			fmt.Printf("%c", source)

			for _, path := range tempPathList {
				fmt.Printf(" -> %c", path)
			}

			fmt.Printf("\n")
		}
	}
}
