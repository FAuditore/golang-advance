package main

import "fmt"

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(data []int, a int, b int) {
	if a >= b {
		return
	}
	pivot := doPivot(data, a, b)
	quickSort(data, a, pivot-1)
	quickSort(data, pivot+1, b)
}

func doPivot(data []int, a int, b int) int {
	key := data[a]
	for a < b {
		for a < b && data[b] >= key {
			b--
		}
		if a < b {
			data[a] = data[b]
		}
		for a < b && data[a] <= key {
			a++
		}
		if a < b {
			data[b] = data[a]
		}
	}
	data[a] = key
	return a
}
