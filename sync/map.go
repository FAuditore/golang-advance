package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("a", 1)
	fmt.Println(m.Load("a"))           //1 true
	fmt.Println(m.LoadOrStore("a", 2)) //1 true

	m.Delete("a")
	fmt.Println(m.LoadOrStore("a", 3)) //3 false
	fmt.Println(m.LoadOrStore("b", 4)) //4 false

	//func返回false时结束遍历
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
