package main

import "fmt"

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	selectSort(arr)
	fmt.Println(arr)
}

//平均O(n2) 不稳定
func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i

		//找到从i+1开始最小的元素的下标赋值给min
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
