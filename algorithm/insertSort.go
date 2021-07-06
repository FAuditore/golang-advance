package main

import "fmt"

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	insertSort(arr)
	fmt.Println(arr)
}

//平均O(n2) 稳定
func insertSort(arr []int) {
	for i := range arr {
		val := arr[i]
		index := i - 1

		//把arr[i]提出来，插到第一个比他小的元素后面
		for index >= 0 && arr[index] > val {
			arr[index+1] = arr[index]
			index--
		}

		arr[index+1] = val
	}
}
