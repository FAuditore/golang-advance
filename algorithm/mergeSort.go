package main

import "fmt"

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	arr = mergeSort(arr)
	fmt.Println(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left := mergeSort(arr[0 : len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0)
	m, n := 0, 0
	l, r := len(left), len(right)

	for m < l && n < r {
		if left[m] < right[n] {
			result = append(result, left[m])
			m++
		} else {
			result = append(result, right[n])
			n++
		}
	}
	result = append(result, left[m:]...)
	result = append(result, right[n:]...)
	return result
}
