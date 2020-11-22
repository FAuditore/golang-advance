package main

import (
	"fmt"
	"sort"
)

func main() {
	var s students = []stu{
		stu{1},
		stu{3},
		stu{7},
		stu{4}}
	sort.Sort(s)
	fmt.Println(s)

	//algorithm.Ints(text)
	var a = []int{1, 34, 5215, 39, 4, 315, 2, 15, 40, 23}
	quicksort(a, 0, len(a)-1)
	fmt.Println(a)

	var b = []int{1, 34, 5215, 39, 4, 315, 2, 15, 40, 23}
	heapSort(b,len(b)-1)
	fmt.Println(b)
}

func quicksort(data []int, a int, b int) {
	if a < b {
		pivot := doPivot(data, a, b)
		quicksort(data, a, pivot-1)
		quicksort(data, pivot+1, b)
	}
}

func doPivot(data []int, i int, j int) int {
	key := data[i]
	for i < j {
		for i < j && data[j] >= key {
			j--
		}
		if i < j {
			data[i] = data[j]
		}
		for i < j && data[i] <= key {
			i++
		}
		if i < j {
			data[j] = data[i]
		}
	}
	fmt.Println(i, j)
	data[i] = key
	return i
}

func heapSort(tree []int, n int) {
	buildHeap(tree, n)
	for i := n; i >= 0; i-- {
		tree[0], tree[i] = tree[i], tree[0]
		heapify(tree, i, 0)
	}
}

func buildHeap(tree []int, n int) {
	lastParent := (n - 1) / 2
	for i := lastParent; i >= 0; i-- {
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
		tree[max], tree[i] = tree[i], tree[max]
		heapify(tree, n, max)
	}
}

type stu struct {
	id int
}

type students []stu

func (s students) Len() int {
	return len(s)
}

func (s students) Less(i, j int) bool {
	return s[i].id < s[j].id
}

func (s students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
