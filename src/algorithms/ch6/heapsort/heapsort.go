package main

import (
	"fmt"
	"math"
)

const max = 999

func main() {
	fmt.Println([]int{ 18, 15, 2, 4, 8, 10, 11, 14, 16, 17 })
	fmt.Println(heapSort([]int{ 18, 15, 2, 4, 8, 10, 11, 14, 16, 17 }))
}

func heapSort(data []int) (result []int) {
	tree := twoForkTree(data)

	var min int

	for range data {
		min, tree = extractMin(tree)

		result = append(result, min)
	}

	return
}

func extractMin(tree []int) (int, []int) {
	root := tree[0]

	treeLen := len(tree)

	if treeLen == 1 {
		return root, []int{}
	}

	lastItem := tree[len(tree) - 1]
	tree[0] = lastItem
	tree = check(tree[:len(tree) - 1], 0)
	
	return root, tree
}

func getIndex(i int) int {
	return int(math.Floor(float64((i - 1) / 2)))
}

func check(tree []int, root int) []int {
	left := root * 2 + 1
	right := root * 2 + 2

	treeLen := len(tree)

	if right >= treeLen {
		return tree
	}

	if tree[root] > tree[left] && (left == treeLen - 1 || tree[left] < tree[right]) {
		tree[left], tree[root] = tree[root], tree[left]
		tree = check(tree, left)
	} else if tree[root] > tree[right] && tree[right] < tree[left] {
		tree[right], tree[root] = tree[root], tree[right]
		tree = check(tree, right)
	}

	return tree
}

func twoForkTree(data []int) (tree []int) {
	for _, item := range data {
		tree  = appendTree(tree, item)
	}

	return tree
}

func appendTree(tree []int, item int) []int {
	index := len(tree)
	tree = append(tree, item)

	if index == 0 {
		return tree
	}

	var newIndex int

	for ;; {
		newIndex = getIndex(index)

		if item <= tree[newIndex] {
			tree[index], tree[newIndex] = tree[newIndex], tree[index]
			if newIndex == 0 {
				break
			}

			index = newIndex
		} else {
			break
		}
	}

	return tree
}