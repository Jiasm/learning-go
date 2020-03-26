package main

import "fmt"

func main() {
	things := map[int]string{
		1: "undershorts",
		2: "socks",
		3: "compression shorts",
		4: "hose",
		5: "cup",
		6: "pants",
		7: "skates",
		8: "leg pads",
		9: "T shirt",
		10: "chest pad",
		11: "sweater",
		12: "mask",
		13: "catch glove",
		14: "blocker",
	}

	max := len(things)

	arr := make([][]int, max + 1)

	arr[1] = []int{3}
	arr[2] = []int{4}
	arr[3] = []int{4, 5}
	arr[4] = []int{6}
	arr[5] = []int{6}
	arr[6] = []int{7, 11}
	arr[7] = []int{8}
	arr[8] = []int{13}
	arr[9] = []int{10}
	arr[10] = []int{11}
	arr[11] = []int{12}
	arr[12] = []int{13}
	arr[13] = []int{14}
	arr[14] = []int{}

	fmt.Println("sort graph is:")
	result := topologicalSort(arr)
	for _, item := range result {
		fmt.Printf("%s,\n", things[item])
	}
}

func topologicalSort(graph [][]int) []int {
	max := len(graph)

	var result []int
	tempArr := make([]int, max + 1)
	var next []int

	for index := 1; index < max; index++ {
		graphItem := graph[index]

		for _, item := range graphItem {
			tempArr[item]++
		}
	}

	for index := 1; index < max; index++ {
		if tempArr[index] == 0 {
			next = append(next, index)
		}
	}

	var nextItem int
	for ;; {
		if len(next) == 0 {
			break
		}

		nextItem, next = next[len(next) - 1], next[:len(next) - 1]

		result = append(result, nextItem)

		graphItem := graph[nextItem]

		for _, item := range graphItem {
			tempArr[item]--

			if tempArr[item] == 0 {
				next = append(next, item)
			}
		}
	}


	return result
}