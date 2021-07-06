package main

import "fmt"

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}

	bubbleSort(arr)
	fmt.Println(arr)
}

//平均O(n2) 稳定
func bubbleSort(arr []int) {
	for i := range arr {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
