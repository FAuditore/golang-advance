package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	heapSort(arr, len(arr))
	fmt.Println(arr)
}

func heapSort(tree []int, n int) {
	buildHeap(tree, n)
	for i := n - 1; i >= 0; i-- {
		tree[0], tree[i] = tree[i], tree[0]
		heapify(tree, i, 0)
	}
}

func buildHeap(tree []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(tree, n, i)
	}
}

func heapify(tree []int, n int, i int) {
	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}
	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}
	if max != i {
		tree[i], tree[max] = tree[max], tree[i]
		heapify(tree, n, max)
	}
}
