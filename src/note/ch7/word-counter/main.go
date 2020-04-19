package main

import (
	"fmt"
	"bufio"
	"strings"
)

type WordCounter int

func (w *WordCounter) Write(words []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(words)))

	scanner.Split(bufio.ScanWords)

	count := 0

	for scanner.Scan() {
		count++
	}

	*w += WordCounter(count)

	return count, nil
}

func main() {
	var w1 WordCounter

	w1.Write([]byte("hello world"))

	fmt.Println(w1)

	w1 = 0

	fmt.Fprintf(&w1, "nice to meet your\nnmy love.")

	fmt.Println(w1)
}